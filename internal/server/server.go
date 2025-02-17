package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pashest/object-storage-service/internal/model"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*http.Server

	storageService storageService
}

func NewServer(storageService storageService) *Server {
	s := &Server{
		storageService: storageService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/upload", s.uploadFileHandler)
	mux.HandleFunc("/download", s.downloadFileHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"X-Filename", "Content-Type", "Username"},
		AllowCredentials: true,
	}).Handler(mux)

	s.Server = &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler,
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
	user := r.Header.Get("Username")
	if user == "" {
		user = "user"
	}

	fileName := r.Header.Get("X-Filename")
	if fileName == "" {
		http.Error(w, "Missing file name in X-Filename header", http.StatusBadRequest)
		return
	}

	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		if _, err := io.Copy(writer, r.Body); err != nil {
			log.Error().Msgf("Error writing to pipe, err: %v", err)
		}
	}()

	log.Print("Starting UploadFile file")
	err := s.storageService.UploadFile(ctx, reader,
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

	user := r.Header.Get("Username")
	if user == "" {
		user = "user"
	}

	fileName := r.URL.Query().Get("file_name")
	if fileName == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}

	reader, err := s.storageService.DownloadFile(ctx,
		model.FileInfo{
			FileName: fileName,
			User:     user,
		})
	if err != nil {
		http.Error(w, "Failed to start file download: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, reader)
	if err != nil {
		log.Printf("Failed to send file %s: %v", fileName, err)
	}
}
