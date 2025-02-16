package meta

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pashest/object-storage-service/config"
	"github.com/pressly/goose/v3"
)

const (
	maxRetries = 3
)

func NewConnectionPool(cfg *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.MetaService.DB.Address)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	bckOff := backoff.WithMaxRetries(backoff.NewConstantBackOff(2*time.Second), maxRetries)
	err = backoff.Retry(func() error {
		return runMigrations(pool, cfg)
	}, bckOff)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

// TODO: to migrations script with checking db
func runMigrations(pool *pgxpool.Pool, cfg *config.Config) error {
	migrationDir := cfg.MetaService.DB.Migrations
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("unable to acquire connection from pool: %v", err)
	}
	defer conn.Release()

	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	if err := goose.Up(db, migrationDir); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}
