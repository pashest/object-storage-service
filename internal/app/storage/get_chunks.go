package storage

import (
	"io"
	"os"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
)

const (
	blockSize  = 4 * 1024 * 1024
	storageDir = "storage_dir/"
)

// GetChunk method for getting chunks
func (i *Implementation) GetChunk(req *desc.GetChunkRequest, stream desc.StorageService_GetChunkServer) error {
	chunkID := req.GetChunkId()
	file, err := os.Open(storageDir + chunkID)
	if err != nil {
		log.Printf("Failed to open chunk %s: %v", chunkID, err)
		return err
	}
	defer file.Close()

	buffer := make([]byte, blockSize)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			err = stream.Send(&desc.GetChunkResponse{
				ChunkId: chunkID,
				Data:    buffer[:n],
			})
			if err != nil {
				log.Printf("Failed to send chunk %s: %v", chunkID, err)
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Failed to read chunk %s: %v", chunkID, err)
			return err
		}
	}

	log.Printf("Sent chunk %s: %v", chunkID, err)

	return nil
}
