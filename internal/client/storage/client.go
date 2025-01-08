package storage

import (
	"context"
	"fmt"
	"io"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"google.golang.org/grpc"
)

// Client define client for work with PromoCodeClient
type Client struct {
	client desc.StorageServiceClient
}

// New return new instance of Client
func New(conn *grpc.ClientConn) *Client {
	return &Client{client: desc.NewStorageServiceClient(conn)}
}

// UploadChunks method for upload chunks
func (c *Client) UploadChunks(ctx context.Context, chunkID string, data []byte) error {
	stream, err := c.client.UploadChunks(ctx)
	if err != nil {
		return err
	}
	err = stream.Send(&desc.UploadChunksRequest{
		ChunkId: chunkID,
		Data:    data,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetChunks method for getting chunks
func (c *Client) GetChunks(ctx context.Context, chunkIDs []string) (map[string][]byte, error) {
	stream, err := c.client.GetChunks(ctx, &desc.GetChunksRequest{
		ChunkIds: chunkIDs,
	})
	if err != nil {
		return nil, err
	}

	chunks := make(map[string][]byte)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to receive chunk response: %w", err)
		}

		chunks[resp.GetChunkId()] = resp.Data
	}

	return chunks, nil
}
