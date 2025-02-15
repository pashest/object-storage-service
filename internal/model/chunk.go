package model

type ChunkInfo struct {
	ChunkName     string `db:"chunk_name"`
	FileUID       string `db:"file_uid"`
	ChunkNumber   int16  `db:"chunk_number"`
	ChunkSize     int64  `db:"chunk_size"`
	ServerAddress string `db:"server_address"`
	CreatedAt     string `db:"created_at"`
	// FileSize
	// FileStatus
}
