package storage

import (
	"context"

	desc "github.com/pashest/object-storage-service/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client define client for work with PromoCodeClient
type Client struct {
	client desc.StorageServiceClient
}

// New return new instance of Client
func New(serviceName string) (*Client, error) {
	conn, err := grpc.NewClient(serviceName, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{client: desc.NewStorageServiceClient(conn)}, nil
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
func (c *Client) GetChunks(ctx context.Context, chunkIDs []string) error {
	stream, err := c.client.GetChunks(ctx, &desc.GetChunksRequest{
		ChunkIds: chunkIDs,
	})
	if err != nil {
		return err
	}

	defer func() {
		for {
			if _, err = stream.Recv(); err != nil {
				break
			}
		}
	}()

	return nil
}
