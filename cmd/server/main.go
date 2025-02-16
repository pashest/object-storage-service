package main

import (
	"context"

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

// TODO: graceful shutdown
func main() {
	ctx := context.Background()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get config")
	}

	connectionPool := client.NewConnectionPool()
	dbPool, err := meta.NewConnectionPool(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get db pool")
	}

	// Repositories
	storageServRepo := storageservers.New(dbPool)
	fileRepo := files.New(dbPool)

	// Services
	// TODO: host and port to env config
	storageMonitoringService := storagemonitoring.New(ctx, connectionPool, storageServRepo, "storage-server", 50051)
	storageService := storage.New(storageMonitoringService, fileRepo, connectionPool)

	server := server.NewServer(storageService)

	log.Printf("Starting server on %s", ":8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}
