package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pashest/object-storage-service/internal/model"
	"github.com/rs/zerolog/log"
)

var storageServers = []string{
	"storage-server-1:50051",
	"storage-server-2:50051",
	"storage-server-3:50051",
	"storage-server-4:50051",
	"storage-server-5:50051",
	"storage-server-6:50051",
}

var chunkIDs = []string{
	"chunk1",
	"chunk2",
	"chunk3",
	"chunk4",
	"chunk5",
	"chunk6",
}

var chunkServerMap = map[string]string{
	"chunk1": "storage-server-1:50051",
	"chunk2": "storage-server-2:50051",
	"chunk3": "storage-server-3:50051",
	"chunk4": "storage-server-4:50051",
	"chunk5": "storage-server-5:50051",
	"chunk6": "storage-server-6:50051",
}

type Server struct {
	*http.Server

	storageService storageService
}

func NewServer(storageService storageService) Server {
	s := Server{
		storageService: storageService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/upload", s.uploadFileHandler)
	mux.HandleFunc("/download", s.downloadFileHandler)

	s.Server = &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		IdleTimeout:  time.Minute,
	}
	return s
}

func (s Server) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	r.Body = http.MaxBytesReader(w, r.Body, 10<<30)
	fileSize := r.ContentLength
	fileName := r.Header.Get("X-Filename")
	user := r.Header.Get("User")

	log.Print("Starting process file")
	err := s.storageService.UploadFile(ctx, r.Body,
		model.FileInfo{
			FileName: fileName,
			User:     user,
			FileSize: fileSize,
		})
	if err != nil {
		http.Error(w, "Failed to upload file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully")
}

func (s Server) downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	user := r.Header.Get("User")
	fileName := r.URL.Query().Get("file_name")
	if fileName == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	reader, err := s.storageService.DownloadFile(ctx,
		model.FileInfo{
			FileName: fileName,
			User:     user,
		})
	if err != nil {
		http.Error(w, "Failed to start file download: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, reader)
	if err != nil {
		log.Printf("Failed to send file %s: %v", fileName, err)
	}
}
