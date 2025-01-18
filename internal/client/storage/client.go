package storage

import (
	"context"
	"fmt"
	"io"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const blockSize = 4 * 1024 * 1024

// Client define client for work with PromoCodeClient
type Client struct {
	client desc.StorageServiceClient
}

// New return new instance of Client
func New(conn *grpc.ClientConn) *Client {
	return &Client{client: desc.NewStorageServiceClient(conn)}
}

// UploadChunks method for upload chunks
func (c *Client) UploadChunks(ctx context.Context, chunkID string, reader io.Reader, chunkSize int64) error {
	log.Printf("Chunk size: %d", chunkSize)
	stream, err := c.client.UploadChunks(ctx)
	if err != nil {
		return err
	}

	buffer := make([]byte, blockSize)
	for chunkSize > 0 {
		if chunkSize < blockSize {
			buffer = buffer[:chunkSize]
		}
		n, err := reader.Read(buffer)
		if n > 0 {
			err := stream.Send(&desc.UploadChunksRequest{
				ChunkId: chunkID,
				Data:    buffer[:n],
			})
			if err != nil {
				log.Printf("Failed to upload chunk %s: %v", chunkID, err)
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
		chunkSize -= int64(n)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to close stream: %w", err)
	}

	if !resp.Success {
		return fmt.Errorf("server error: %s", resp.Message)
	}

	return nil
}

// GetChunk method for getting chunks
func (c *Client) GetChunk(ctx context.Context, chunkID string, writer io.Writer) error {
	stream, err := c.client.GetChunk(ctx, &desc.GetChunkRequest{
		ChunkId: chunkID,
	})
	if err != nil {
		return fmt.Errorf("failed to get chunk stream: %w", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive chunk response: %w", err)
		}

		_, err = writer.Write(resp.GetData())
		if err != nil {
			log.Printf("Failed to write chunk %s: %v", resp.GetChunkId(), err)
			return err
		}
	}

	return nil
}
