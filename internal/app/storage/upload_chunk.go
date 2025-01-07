package storage

import (
	"fmt"
	"io"

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

		// err = SaveChunk(req.Id, req.Data)
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
