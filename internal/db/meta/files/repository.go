package files

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pashest/object-storage-service/internal/db"
	"github.com/pashest/object-storage-service/internal/model"
)

// Repository define repo for work with model.StorageServer
type Repository struct {
	*pgxpool.Pool
}

const (
	fileTableName  = "files_info"
	chunkTableName = "chunks_info"
)

// New return new instance of Repository
func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

func (r *Repository) AddFileInfo(ctx context.Context, file model.FileInfo) (string, error) {
	sql, args, err := db.PgQb().
		Insert(fileTableName).
		Columns(
			"file_name",
			"user_id",
			"file_size",
			"status",
			"created_at",
			"updated_at",
		).
		Values(
			file.FileName,
			file.User,
			file.FileSize,
			model.FileCreated,
			time.Now(),
			time.Now(),
		).
		Suffix("RETURNING file_uid").
		ToSql()
	if err != nil {
		return "", err
	}

	var fileUID string
	err = r.Pool.QueryRow(ctx, sql, args).Scan(&fileUID)
	if err != nil {
		return "", err
	}

	return fileUID, nil
}

func (r *Repository) UpdateFileInfo(ctx context.Context, file model.FileInfo) error {
	sql, args, err := db.PgQb().
		Update(fileTableName).
		SetMap(map[string]interface{}{
			"status":     file.Status,
			"updated_at": time.Now(),
		}).
		Where(squirrel.Eq{"f.file_name": file.FileName}).
		Where(squirrel.Eq{"f.user_id": file.User}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetFileInfoByFileNameAndUser(ctx context.Context, fileName, user string) (*model.FileInfo, error) {
	sql, args, err := db.PgQb().
		Select(
			"file_uid",
			"file_size",
			"status",
		).
		From(fileTableName).
		Where(squirrel.Eq{"file_name": fileName}).
		Where(squirrel.Eq{"user_id": user}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var info model.FileInfo
	err = r.Pool.QueryRow(ctx, sql, args).Scan(&info.FileUID, &info.FileSize, &info.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &info, err
}

func (r *Repository) AddChunkInfo(ctx context.Context, chunk model.ChunkInfo) error {
	sql, args, err := db.PgQb().
		Insert(chunkTableName).
		Columns(
			"chunk_name",
			"file_uid",
			"chunk_number",
			"chunk_size",
			"server_address",
			"created_at",
		).
		Values(
			chunk.ChunkName,
			chunk.FileUID,
			chunk.ChunkNumber,
			chunk.ChunkSize,
			chunk.ServerAddress,
			time.Now(),
		).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetChunksInfoByFileNameAndUser(ctx context.Context, fileName, user string) ([]model.ChunkInfo, error) {
	sql, args, err := db.PgQb().
		Select(
			"ch.chunk_name",
			"ch.file_uid",
			"ch.chunk_number",
			"ch.chunk_size",
			"ch.server_address",
			"ch.created_at",
		).
		From(fileTableName + " f").
		InnerJoin(chunkTableName + " ch on ch.file_uid = f.file_uid").
		Where(squirrel.Eq{"f.file_name": fileName}).
		Where(squirrel.Eq{"f.user_id": user}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var chunks []model.ChunkInfo

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chunks, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.ChunkInfo])
	return chunks, err
}
