package helper

import (
	"context"
	"fmt"

	"github.com/pashest/object-storage-service/internal/utils"
	desc "github.com/pashest/object-storage-service/pkg/helper"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Heartbeat method for checking health of storage server
func (i *Implementation) Heartbeat(_ context.Context, _ *emptypb.Empty) (*desc.HeartbeatResponse, error) {
	freeSpace, err := utils.GetFreeDiskSpace("./storage_dir")
	if err != nil {
		msg := fmt.Sprintf("Heartbeat: failed to get disk space: %v", err)
		log.Error().Msg(msg)
		return &desc.HeartbeatResponse{
			Alive:   false,
			Message: msg,
		}, nil
	}

	return &desc.HeartbeatResponse{
		Alive:     true,
		FreeSpace: freeSpace,
	}, nil
}
