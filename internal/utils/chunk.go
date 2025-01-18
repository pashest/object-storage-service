package utils

import (
	"bytes"
)

const chunkCount = 6

// // SplitIntoChunks split file into chunks
// func SplitIntoChunks(file io.Reader, fileSize int64, chunkChan chan<- model.Chunk) error {
// 	chunks := make([][]byte, chunkCount)
//
// 	chunkSize := fileSize / chunkCount
// 	extraBytes := fileSize % chunkCount
//
// 	for i := 0; i < chunkCount; i++ {
// 		size := chunkSize
// 		if i == chunkCount-1 {
// 			size += extraBytes
// 		}
//
// 		chunk := make([]byte, size)
// 		n, err := file.Read(chunk)
// 		if err != nil && err != io.EOF {
// 			return fmt.Errorf("failed to read chunk %d: %w", i, err)
// 		}
// 		if n == 0 {
// 			break
// 		}
//
// 		chunks[i] = chunk[:n]
// 	}
//
// 	// Add zero bytes if the file is less than 6 bytes
// 	for i := len(chunks) - 1; i >= 0; i-- {
// 		if len(chunks[i]) == 0 {
// 			chunks[i] = []byte{0}
// 		}
// 	}
//
// 	return chunks, nil
// }

func JoinChunks(chunks [][]byte) []byte {
	var buffer bytes.Buffer
	for _, chunk := range chunks {
		buffer.Write(chunk)
	}
	return buffer.Bytes()
}
