package storage_monitoring

import (
	"context"

	"github.com/pashest/object-storage-service/internal/client/helper"
)

type (
	connectionPool interface {
		AddConnection(address string) error
		RemoveConnection(address string)
		GetHelperClient(address string) (*helper.Client, bool)
	}

	storageServersRepo interface {
		AddServer(ctx context.Context, address string) error
	}
)
