package storage

import (
	"io"
	"os"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
)

// GetChunks method for getting chunks
func (i *Implementation) GetChunks(req *desc.GetChunksRequest, stream desc.StorageService_GetChunksServer) error {
	for _, chunkID := range req.GetChunkIds() {
		data, err := readChunk(chunkID)

		if err != nil {
			log.Printf("Failed to read chunk %s: %v", chunkID, err)
			return err
		}

		err = stream.Send(&desc.GetChunksResponse{
			ChunkId: chunkID,
			Data:    data,
		})
		if err != nil {
			log.Printf("Failed to send chunk %s: %v", chunkID, err)
			return err
		}
	}

	return nil
}

func readChunk(chunkID string) ([]byte, error) {
	fileName := "storage_dir/" + chunkID
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
