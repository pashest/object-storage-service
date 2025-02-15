package server

import (
	"context"
	"io"

	"github.com/pashest/object-storage-service/internal/model"
)

type (
	storageService interface {
		UploadFile(ctx context.Context, file io.Reader, fileInfo model.FileInfo) error
		DownloadFile(ctx context.Context, fileInfo model.FileInfo) (io.Reader, error)
	}
)
