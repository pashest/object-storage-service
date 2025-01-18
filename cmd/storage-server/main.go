package main

import (
	"net"
	"time"

	"github.com/pashest/object-storage-service/internal/app/helper"
	"github.com/pashest/object-storage-service/internal/app/storage"
	helperpb "github.com/pashest/object-storage-service/pkg/helper"
	storagepb "github.com/pashest/object-storage-service/pkg/storage"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const maxMessageSize = 2 * 1024 * 1024 * 1024 // 2 GB

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Msgf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMessageSize),
		grpc.ConnectionTimeout(5*time.Minute),
	)

	helperpb.RegisterHelperServiceServer(grpcServer, helper.NewHelperService())
	storagepb.RegisterStorageServiceServer(grpcServer, storage.NewStorageService())

	log.Info().Msg("Server is running on port 50051...")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal().Msgf("Failed to serve gRPC server: %v", err)
	}
}
