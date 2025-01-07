package helper

import "github.com/pashest/object-storage-service/pkg/helper"

type Implementation struct {
	helper.UnimplementedHelperServiceServer
}

// NewHelperService return new instance of Implementation.
func NewHelperService() *Implementation {
	return &Implementation{}
}
