package meta

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pashest/object-storage-service/config"
)

func NewConnectionPool(cfg *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.MetaService.DB.Address)
	if err != nil {
		return nil, err
	}
	return pgxpool.NewWithConfig(context.Background(), config)
}
