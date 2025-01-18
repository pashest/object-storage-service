package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/pashest/object-storage-service/internal/client"
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

const chunkCount = 6

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadFileHandler)
	mux.HandleFunc("/upload-chunk", chunkUploadHandler)
	mux.HandleFunc("/download", downloadFileHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", ":8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}

func chunkUploadHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	clientPool := client.NewClientPool()

	for _, address := range storageServers {
		err := clientPool.AddConnection(address)
		if err != nil {
			log.Printf("Failed to add server %s: %v", address, err)
		}
	}

	// Limit the chunk size ~2.5GB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<28)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, "Error retrieving chunk", http.StatusBadRequest)
		return
	}
	defer file.Close()

	chunkInd := r.FormValue("chunkIndex")
	chunkSize := r.FormValue("chunkSize")
	chSize, _ := strconv.ParseInt(chunkSize, 10, 64)
	chunkIndex, _ := strconv.Atoi(chunkInd)

	if clnt, ok := clientPool.GetStorageClient(storageServers[chunkIndex]); ok {
		log.Printf("Chunk %s size: %s bytes", chunkIDs[chunkIndex], chunkSize)

		err = clnt.UploadChunks(ctx, chunkIDs[chunkIndex], file, chSize)
		if err != nil {
			log.Printf("Failed to upload chunk %s, err: %v", chunkIDs[chunkIndex], err)
			http.Error(w, "Failed to upload chunk", http.StatusBadRequest)
			return
		}
		log.Printf("Chunk %s uploaded successfully", chunkIDs[chunkIndex])
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chunk uploaded successfully")
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	clientPool := client.NewClientPool()

	for _, address := range storageServers {
		err := clientPool.AddConnection(address)
		if err != nil {
			log.Printf("Failed to add server %s: %v", address, err)
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<30)
	fileSize := r.ContentLength

	log.Print("Starting process file")
	err := processFile(r.Body, fileSize, clientPool)
	if err != nil {
		http.Error(w, "Failed to process file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully")
}

func processFile(file io.Reader, fileSize int64, pool *client.ClientPool) error {
	ctx := context.Background()
	chunkSize := fileSize / chunkCount
	extraBytes := fileSize % chunkCount

	for i := 0; i < chunkCount; i++ {
		if i == chunkCount-1 {
			chunkSize += extraBytes
		}

		if clnt, ok := pool.GetStorageClient(storageServers[i]); ok {
			log.Printf("Chunk %s size: %d bytes", chunkIDs[i], chunkSize)

			log.Print("Starting UploadChunk")
			err := clnt.UploadChunks(ctx, chunkIDs[i], file, chunkSize)
			if err != nil {
				log.Printf("Failed to upload chunk %s, err: %v", chunkIDs[i], err)
				return err
			}
			log.Printf("Chunk %s uploaded successfully", chunkIDs[i])
		}
	}

	return nil
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

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)

	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		for i := 0; i < 6; i++ {
			if clnt, ok := clientPool.GetStorageClient(storageServers[i]); ok {
				err := clnt.GetChunk(ctx, chunkIDs[i], writer)
				if err != nil {
					log.Printf("Failed to download chunks %s, %v", chunkIDs[i], err)
					http.Error(w, "Failed to download chunks: "+err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}()

	_, err := io.Copy(w, reader)
	if err != nil {
		log.Printf("Failed to send file %s: %v", fileName, err)
	}
}

func getIndex(chunkIDs []string, chunkID string) int {
	for i, id := range chunkIDs {
		if id == chunkID {
			return i
		}
	}
	return -1
}
