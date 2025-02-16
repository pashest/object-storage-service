package storage_monitoring

import (
	"container/heap"
	"fmt"
	"time"
)

type StorageServer struct {
	Address       string
	FreeSpace     int64
	LastHeartbeat int64
}

type StorageHeap []StorageServer

func (h StorageHeap) Len() int           { return len(h) }
func (h StorageHeap) Less(i, j int) bool { return h[i].FreeSpace > h[j].FreeSpace } // Max-Heap
func (h StorageHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StorageHeap) Push(x any) {
	*h = append(*h, x.(StorageServer))
}

func (h *StorageHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (s *Service) UpdateStorageHeap(address string, freeSpace int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	for i := 0; i < len(s.storageHeap); i++ {
		if s.storageHeap[i].Address == address {
			s.storageHeap[i].FreeSpace = freeSpace
			s.storageHeap[i].LastHeartbeat = time.Now().Unix()
			heap.Fix(&s.storageHeap, i)
			found = true
			break
		}
	}

	if !found {
		heap.Push(&s.storageHeap, StorageServer{
			Address:       address,
			FreeSpace:     freeSpace,
			LastHeartbeat: time.Now().Unix(),
		})
	}
}

func (s *Service) GetBestStorageServerAddress() (address string, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.storageHeap) == 0 {
		return "", fmt.Errorf("there aren't any storage servers")
	}

	serv := s.storageHeap[0]
	if serv.FreeSpace <= 0 {
		return "", fmt.Errorf("there is no suitable storage server")
	}

	return serv.Address, nil
}

func (s *Service) GetStorageServers() []StorageServer {
	s.mu.RLock()
	storageServers := make([]StorageServer, len(s.storageHeap))
	_ = copy(storageServers, s.storageHeap)
	s.mu.RUnlock()
	return storageServers
}
