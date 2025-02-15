package model

import "time"

type FileStatus string

const (
	FileCreated            FileStatus = "CREATED"
	FilePartlyUploaded     FileStatus = "PARTLY_UPLOADED"
	FileCompletelyUploaded FileStatus = "COMPLETELY_UPLOADED"
)

type FileInfo struct {
	FileUID   string     `db:"file_uid"`
	FileName  string     `db:"file_name"`
	User      string     `db:"user_id"`
	FileSize  int64      `db:"file_size"`
	Status    FileStatus `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
