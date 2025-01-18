package helper

import (
	"context"

	"github.com/pashest/object-storage-service/internal/model"
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
func (c *Client) Heartbeat(ctx context.Context) (*model.Heartbeat, error) {
	res, err := c.client.Heartbeat(ctx, nil)
	if err != nil {
		return nil, err
	}

	return protoHeartbeatToHeartbeat(res), nil
}

func protoHeartbeatToHeartbeat(hb *desc.HeartbeatResponse) *model.Heartbeat {
	if hb == nil {
		return nil
	}
	return &model.Heartbeat{
		Alive:     hb.Alive,
		Message:   hb.Message,
		FreeSpace: hb.FreeSpace,
	}
}
