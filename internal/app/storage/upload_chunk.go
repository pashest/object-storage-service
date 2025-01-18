package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
)

// UploadChunks method for upload chunks
func (i *Implementation) UploadChunks(stream desc.StorageService_UploadChunksServer) error {
	var chunkID string
	tmpFile, err := os.CreateTemp("storage_dir", "temp-chunk-*")
	if err != nil {
		log.Printf("Failed to create temp file: %v", err)
		return err
	}
	defer tmpFile.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if errors.Is(stream.Context().Err(), context.Canceled) {
				log.Printf("Context canceled for chunk %s: %v", chunkID, stream.Context().Err())
			}
			log.Error().Msg(fmt.Sprintf("Failed to receive chunk: %v", err))
			return err
		}

		if chunkID == "" {
			chunkID = req.GetChunkId()
			log.Printf("Started receiving chunk %s", chunkID)
		}

		_, err = tmpFile.Write(req.GetData())
		if err != nil {
			log.Printf("Failed to write chunk %s: %v", chunkID, err)
			return err
		}
	}

	err = saveChunkToStorage(chunkID, tmpFile.Name())
	if err != nil {
		msg := fmt.Sprintf("Failed to save chunk %s: %v", chunkID, err)
		log.Error().Msg(msg)
		return stream.SendAndClose(&desc.UploadChunksResponse{
			Success: false,
			Message: msg,
		})
	}

	return stream.SendAndClose(&desc.UploadChunksResponse{
		Success: true,
		Message: "All chunks uploaded successfully",
	})
}

func saveChunkToStorage(chunkID string, tempFilePath string) error {
	finalPath := "storage_dir/" + chunkID

	err := os.Rename(tempFilePath, finalPath)
	if err != nil {
		return fmt.Errorf("failed to move chunk %s to storage: %w", chunkID, err)
	}

	return nil
}
