package storage_monitoring

import (
	"container/heap"
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// Service ...
type Service struct {
	connectionPool connectionPool
	mu             sync.RWMutex
	storageHeap    StorageHeap
}

// New return new instance of Service
func New(ctx context.Context, connectionPool connectionPool) *Service {
	s := &Service{
		connectionPool: connectionPool,
	}
	heap.Init(&s.storageHeap)

	go s.startHeartbeat(ctx)

	return s
}

// AddServer adds new storage server
func (s *Service) AddServer(address string) error {
	err := s.connectionPool.AddConnection(address)
	if err != nil {
		return err
	}
	s.UpdateStorageHeap(address, 0)

	return nil
}

// heartbeatHandler checks storage servers
func (s *Service) heartbeatHandler(ctx context.Context) error {
	storageServers := make([]StorageServer, 0)
	s.mu.RLock()
	_ = copy(storageServers, s.storageHeap)
	s.mu.RUnlock()
	for _, server := range storageServers {
		helperClnt, exist := s.connectionPool.GetHelperClient(server.Address)
		if !exist {
			log.Error().Msgf("Not found connection for server: %s", server.Address)
			s.UpdateStorageHeap(server.Address, 0)
			continue
		}

		var freeSpace int64
		heartbeat, err := helperClnt.Heartbeat(ctx)
		if err != nil {
			log.Error().Msgf("Failed to get heartbeat from server: %s", server.Address)
		}

		if !heartbeat.Alive {
			log.Error().Msgf("Heartbeat failed. Server: %s Msg: %s", server.Address, heartbeat.Message)
		} else {
			freeSpace = heartbeat.FreeSpace
		}

		s.UpdateStorageHeap(server.Address, freeSpace)
	}
	return nil
}

func (s *Service) startHeartbeat(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.heartbeatHandler(ctx)
	}
}
