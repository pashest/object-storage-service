package storage_servers

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pashest/object-storage-service/internal/db"
)

// Repository define repo for work with model.StorageServer
type Repository struct {
	*pgxpool.Pool
}

const tableName = "storage_servers"

// New return new instance of Repository
func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

func (r *Repository) AddressesList(ctx context.Context) ([]string, error) {
	sql, _, err := db.PgQb().
		Select("address").
		From(tableName).
		ToSql()
	if err != nil {
		return nil, err
	}

	var addresses []string

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses, err = pgx.CollectRows(rows, pgx.RowTo[string])
	return addresses, err
}

func (r *Repository) AddServer(ctx context.Context, address string) error {
	sql, _, err := db.PgQb().
		Insert(tableName).
		Columns("address").
		Values(address).
		ToSql()
	if err != nil {
		return err
	}

	var uid string
	err = r.Pool.QueryRow(ctx, sql).Scan(&uid)
	if err != nil {
		return err
	}

	return nil
}
