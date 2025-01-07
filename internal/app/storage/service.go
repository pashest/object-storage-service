package storage

import (
	"github.com/pashest/object-storage-service/pkg/storage"
)

type Implementation struct {
	storage.UnimplementedStorageServiceServer
}

// NewStorageService return new instance of Implementation.
func NewStorageService() *Implementation {
	return &Implementation{}
}
