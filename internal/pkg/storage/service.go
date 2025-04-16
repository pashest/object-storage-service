package storage

import (
	"cmp"
	"context"
	"fmt"
	"io"
	"slices"

	"github.com/pashest/object-storage-service/internal/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Service ...
type Service struct {
	storageMonitoringService storageMonitoringService
	filesRepo                filesRepo
	connectionPool           connectionPool
}

const chunkCount = 6

// New return new instance of Service
func New(
	storageMonitoringService storageMonitoringService,
	filesRepo filesRepo,
	connectionPool connectionPool,
) *Service {
	return &Service{
		storageMonitoringService: storageMonitoringService,
		filesRepo:                filesRepo,
		connectionPool:           connectionPool,
	}
}

func (s Service) UploadFile(ctx context.Context, file io.Reader, fileInfo model.FileInfo) error {
	exFileInfo, err := s.filesRepo.GetFileInfoByFileNameAndUser(ctx, fileInfo.FileName, fileInfo.User)
	if err != nil {
		return errors.Wrap(err, "UploadFile: GetFileInfoByFileNameAndUser")
	}

	existChunksMap := make(map[int16]int64, chunkCount)
	if exFileInfo != nil {
		fileInfo.FileUID = exFileInfo.FileUID
		if exFileInfo.Status == model.FileCompletelyUploaded || exFileInfo.FileSize != fileInfo.FileSize {
			// TODO: create new file with serial number (versions)
			return fmt.Errorf("UploadFile: there is already a file with that name: %s", fileInfo.FileName)
		}

		chunksInfo, err := s.filesRepo.GetChunksInfoByFileNameAndUser(ctx, fileInfo.FileName, fileInfo.User)
		if err != nil {
			return errors.Wrap(err, "UploadFile: GetChunksInfoByFileNameAndUser")
		}

		for _, ch := range chunksInfo {
			existChunksMap[ch.ChunkNumber] = ch.ChunkSize
		}
	} else {
		fileInfo.FileUID, err = s.filesRepo.AddFileInfo(ctx, fileInfo)
		if err != nil {
			return errors.Wrap(err, "UploadFile: AddFileInfo")
		}
	}
	err = s.processFile(ctx, file, fileInfo, existChunksMap)
	if err != nil {
		return errors.Wrap(err, "UploadFile: processFile")
	}

	fileInfo.Status = model.FileCompletelyUploaded
	err = s.filesRepo.UpdateFileInfo(ctx, fileInfo)
	if err != nil {
		return errors.Wrap(err, "UploadFile: UpdateFileInfo")
	}

	return nil
}

func (s Service) processFile(
	ctx context.Context,
	file io.Reader,
	fileInfo model.FileInfo,
	existChunksMap map[int16]int64,
) error {
	chunkSize := fileInfo.FileSize / chunkCount
	extraBytes := fileInfo.FileSize % chunkCount

	for i := int16(0); i < chunkCount; i++ {
		if i == chunkCount-1 {
			chunkSize += extraBytes
		}

		if size, ok := existChunksMap[i]; ok {
			if size != chunkSize {
				return fmt.Errorf("chunk size is different with the saved one")
			}

			_, err := io.CopyN(io.Discard, file, size)
			if err != nil {
				return fmt.Errorf("failed to skip bytes, err: %v", err)
			}
			continue
		}

		err := s.uploadChunk(ctx, file, fileInfo, chunkSize, i)
		if err != nil {
			return errors.Wrapf(err, "processFile: uploadChunk %d", i)
		}
	}

	return nil
}

func (s Service) uploadChunk(
	ctx context.Context,
	file io.Reader,
	fileInfo model.FileInfo,
	chunkSize int64,
	chunkNumber int16,
) error {
	chunkName := s.getChunkName(fileInfo.FileUID, chunkNumber)
	storageServer, err := s.storageMonitoringService.GetBestStorageServerAddress()
	if err != nil {
		return fmt.Errorf("failed to get storage server, err: %v", err)
	}

	clnt, ok := s.connectionPool.GetStorageClient(storageServer)
	if !ok {
		return fmt.Errorf("storage client not found for server: %s", storageServer)
	}

	log.Printf("Starting UploadChunk %s bytes: %d", chunkName, chunkSize)
	tx, err := s.filesRepo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction, err: %v", err)
	}
	defer func() {
		if err != nil {
			_ = s.filesRepo.RollbackTx(ctx, tx)
		}
	}()

	err = s.filesRepo.AddChunkInfoInTx(ctx, tx,
		model.ChunkInfo{
			ChunkName:     chunkName,
			FileUID:       fileInfo.FileUID,
			ChunkNumber:   chunkNumber,
			ChunkSize:     chunkSize,
			ServerAddress: storageServer,
		},
	)
	if err != nil {
		return fmt.Errorf("AddChunkInfoInTx chunk %s, err: %v", chunkName, err)
	}

	err = clnt.UploadChunk(ctx, chunkName, file, chunkSize)
	if err != nil {
		return fmt.Errorf("failed to upload chunk %s, err: %v", chunkName, err)
	}

	err = s.filesRepo.CommitTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction, err: %v", err)
	}

	log.Printf("Chunk %s uploaded successfully", chunkName)

	return nil
}

func (s Service) DownloadFile(ctx context.Context, fileInfo model.FileInfo) (io.Reader, error) {
	chunksInfo, err := s.filesRepo.GetChunksInfoByFileNameAndUser(ctx, fileInfo.FileName, fileInfo.User)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadFile: GetChunksInfoByFileNameAndUser")
	}
	if len(chunksInfo) == 0 {
		return nil, fmt.Errorf("DownloadFile: there isn't file with name: %s", fileInfo.FileName)
	}
	// TODO: check status instead of chunkCount
	if len(chunksInfo) < chunkCount {
		return nil, fmt.Errorf("DownloadFile: file isn't fully uploaded: %s", fileInfo.FileName)
	}

	slices.SortFunc(chunksInfo, func(a, b model.ChunkInfo) int {
		return cmp.Compare(a.ChunkNumber, b.ChunkNumber)
	})

	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()

		for _, chunk := range chunksInfo {
			if clnt, ok := s.connectionPool.GetStorageClient(chunk.ServerAddress); ok {
				err = clnt.GetChunk(ctx, chunk.ChunkName, writer)
				if err != nil {
					log.Error().Msgf("DownloadFile: chunk: %s, server: %s, err: %v", chunk.ChunkName, chunk.ServerAddress, err)
					return
				}
			}
		}
	}()

	return reader, nil
}

func (s Service) getChunkName(fileUID string, chunkNumber int16) string {
	return fmt.Sprintf("%s-chunk%d", fileUID, chunkNumber)
}
