package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"
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
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := runGRPC(ctx); err != nil {
		log.Fatal().Msgf("Failed to serve gRPC server: %v", err)
	}
}

func runGRPC(ctx context.Context) error {
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

	srvErr := make(chan error, 1)
	go func() {
		log.Info().Msg("Server is running on port 50051...")
		if err = grpcServer.Serve(listener); err != nil {
			srvErr <- err
		}
	}()

	select {
	case err = <-srvErr:
		log.Error().Err(err).Msg("GRPC: Serve")
		return err
	case <-ctx.Done():
		log.Info().Msg("Shutdown signal received, stopping GRPC server...")
		grpcServer.GracefulStop()
		log.Info().Msg("GRPC Server was closed")
	}

	return nil
}
