package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/pashest/object-storage-service/internal/client"
	"github.com/pashest/object-storage-service/internal/utils"
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

func main() {
	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/download", downloadFileHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal().Err(err)
	}
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	clientPool := client.NewClientPool()

	for _, address := range storageServers {
		err := clientPool.AddConnection(address)
		if err != nil {
			log.Printf("Failed to add server %s: %v", address, err)
		}
	}

	fileName := r.Header.Get("File-Name")
	_ = fileName
	fileSize := r.ContentLength
	if fileSize <= 0 {
		http.Error(w, "Invalid file size", http.StatusBadRequest)
		return
	}

	chunks, err := utils.SplitIntoChunks(r.Body, fileSize)
	if err != nil {
		http.Error(w, "Failed to split file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < 6; i++ {
		if clnt, ok := clientPool.GetStorageClient(storageServers[i]); ok {
			err = clnt.UploadChunks(ctx, chunkIDs[i], chunks[i])
			if err != nil {
				http.Error(w, "Failed to upload file: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	if err != nil {
		http.Error(w, "Failed to upload chunks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully")
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	clientPool := client.NewClientPool()

	for _, address := range storageServers {
		err := clientPool.AddConnection(address)
		if err != nil {
			log.Printf("Failed to add server %s: %v", address, err)
		}
	}

	fileName := r.URL.Query().Get("file_name")
	if fileName == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}

	fileChunks := make([][]byte, len(chunkIDs))

	for i := 0; i < 6; i++ {
		if clnt, ok := clientPool.GetStorageClient(storageServers[i]); ok {
			chunksMap, err := clnt.GetChunks(ctx, []string{chunkIDs[i]})
			if err != nil {
				http.Error(w, "Failed to download chunks: "+err.Error(), http.StatusInternalServerError)
				return
			}

			for chunkID, data := range chunksMap {
				fileChunks[getIndex(chunkIDs, chunkID)] = data
			}
		}
	}

	fileData := utils.JoinChunks(fileChunks)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileData)
}

func getIndex(chunkIDs []string, chunkID string) int {
	for i, id := range chunkIDs {
		if id == chunkID {
			return i
		}
	}
	return -1
}
