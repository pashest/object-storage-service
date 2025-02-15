package storage

import (
	"context"

	"github.com/pashest/object-storage-service/internal/client/storage"
	"github.com/pashest/object-storage-service/internal/model"
)

type (
	connectionPool interface {
		GetStorageClient(address string) (*storage.Client, bool)
	}

	storageMonitoringService interface {
		GetBestStorageServerAddress() (address string, err error)
	}

	filesRepo interface {
		AddFileInfo(ctx context.Context, file model.FileInfo) (string, error)
		GetFileInfoByFileNameAndUser(ctx context.Context, fileName, user string) (*model.FileInfo, error)
		UpdateFileInfo(ctx context.Context, file model.FileInfo) error
		AddChunkInfo(ctx context.Context, chunk model.ChunkInfo) error
		GetChunksInfoByFileNameAndUser(ctx context.Context, fileName, user string) ([]model.ChunkInfo, error)
	}
)
