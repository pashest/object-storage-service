package storage

import (
	"fmt"
	"io"
	"os"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
)

// UploadChunks method for upload chunks
func (i *Implementation) UploadChunks(stream desc.StorageService_UploadChunksServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&desc.UploadChunksResponse{
				Success: true,
				Message: "All chunks uploaded successfully",
			})
		}
		if err != nil {
			log.Error().Msg(fmt.Sprintf("Failed to receive chunk: %v", err))
			return err
		}

		err = saveChunk(req.GetChunkId(), req.GetData())
		if err != nil {
			msg := fmt.Sprintf("Failed to save chunk %s: %v", req.GetChunkId(), err)
			log.Error().Msg(msg)
			return stream.SendAndClose(&desc.UploadChunksResponse{
				Success: false,
				Message: msg,
			})
		}
	}
	return nil
}

func saveChunk(chunkID string, data []byte) error {
	fileName := "storage_dir/" + chunkID
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
