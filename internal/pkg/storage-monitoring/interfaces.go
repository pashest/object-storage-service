package storage_monitoring

import (
	"context"

	"github.com/pashest/object-storage-service/internal/model"
)

type (
	connectionPool interface {
		AddConnection(address string) error
		RemoveConnection(address string)
		GetHelperClient(address string) (helperClient, bool)
	}
	helperClient interface {
		Heartbeat(ctx context.Context) (*model.Heartbeat, error)
	}
)
