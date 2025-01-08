package server

type FileMeta struct {
	FileName string   `json:"file_name"`
	ChunkIDs []string `json:"chunk_ids"`
}
