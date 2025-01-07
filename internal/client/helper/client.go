package helper

import (
	"context"

	desc "github.com/pashest/object-storage-service/pkg/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client define client for work with PromoCodeClient
type Client struct {
	client desc.HelperServiceClient
}

// New return new instance of Client
func New(serviceName string) (*Client, error) {
	conn, err := grpc.NewClient(serviceName, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{client: desc.NewHelperServiceClient(conn)}, nil
}

// Heartbeat method for checking health of storage server
func (c *Client) Heartbeat(ctx context.Context) (*desc.HeartbeatResponse, error) {
	res, err := c.client.Heartbeat(ctx, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
