package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/pashest/object-storage-service/config"
	"github.com/pashest/object-storage-service/internal/client"
	"github.com/pashest/object-storage-service/internal/db/meta"
	"github.com/pashest/object-storage-service/internal/db/meta/files"
	storageservers "github.com/pashest/object-storage-service/internal/db/meta/storage_servers"
	"github.com/pashest/object-storage-service/internal/pkg/storage"
	storagemonitoring "github.com/pashest/object-storage-service/internal/pkg/storage-monitoring"
	"github.com/pashest/object-storage-service/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get config")
	}

	connectionPool := client.NewConnectionPool()
	defer connectionPool.Close()

	dbPool, err := meta.NewConnectionPool(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get db pool")
	}
	defer dbPool.Close()

	// Repositories
	storageServRepo := storageservers.New(dbPool)
	fileRepo := files.New(dbPool)

	// Services
	// TODO: host and port to env config
	storageMonitoringService := storagemonitoring.New(ctx, connectionPool, storageServRepo, "storage-server", 50051)
	storageService := storage.New(storageMonitoringService, fileRepo, connectionPool)

	server := server.NewServer(storageService)

	if err = runServer(ctx, server); err != nil {
		log.Fatal().Msgf("Failed to serve HTTP server: %v", err)
	}
}

func runServer(ctx context.Context, server *server.Server) error {
	srvErr := make(chan error, 1)
	go func() {
		log.Info().Msg("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("HTTP Server")
		return err
	case <-ctx.Done():
		log.Info().Msg("Shutdown signal received, stopping server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("Failed to gracefully shutdown server")
		}

		log.Info().Msg("HTTP Server was closed")
	}

	return nil
}
