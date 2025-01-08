package helper

import (
	"context"

	desc "github.com/pashest/object-storage-service/pkg/helper"
	"google.golang.org/grpc"
)

// Client define client for work with PromoCodeClient
type Client struct {
	client desc.HelperServiceClient
}

// New return new instance of Client
func New(conn *grpc.ClientConn) *Client {
	return &Client{client: desc.NewHelperServiceClient(conn)}
}

// Heartbeat method for checking health of storage server
func (c *Client) Heartbeat(ctx context.Context) (*desc.HeartbeatResponse, error) {
	res, err := c.client.Heartbeat(ctx, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
