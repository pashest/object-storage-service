package storage_monitoring

import (
	"container/heap"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// Service ...
type Service struct {
	storageHeap        StorageHeap
	storageServersRepo storageServersRepo
	host               string
	port               int32

	connectionPool connectionPool
	mu             sync.RWMutex
}

// New return new instance of Service
func New(
	ctx context.Context,
	connectionPool connectionPool,
	storageServersRepo storageServersRepo,
	host string,
	port int32,
) *Service {
	s := &Service{
		connectionPool:     connectionPool,
		storageServersRepo: storageServersRepo,
		host:               host,
		port:               port,
	}
	heap.Init(&s.storageHeap)

	go s.monitorServers(ctx)
	go s.startHeartbeat(ctx)

	return s
}

// AddServer adds new storage server
func (s *Service) AddServer(ctx context.Context, address string) error {
	err := s.connectionPool.AddConnection(address)
	if err != nil {
		return err
	}

	err = s.storageServersRepo.AddServer(ctx, address)
	if err != nil {
		return err
	}

	s.UpdateStorageHeap(address, 0)

	return nil
}

// heartbeatHandler checks storage servers
func (s *Service) heartbeatHandler(ctx context.Context) error {
	storageServers := s.GetStorageServers()
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

		if heartbeat != nil {
			if !heartbeat.Alive {
				log.Error().Msgf("Heartbeat failed. Server: %s Msg: %s", server.Address, heartbeat.Message)
			} else {
				freeSpace = heartbeat.FreeSpace
			}
		}

		s.UpdateStorageHeap(server.Address, freeSpace)
	}
	return nil
}

// monitorHandler monitor storage servers
func (s *Service) monitorHandler(ctx context.Context) error {
	addrs, err := net.LookupHost(s.host)
	if err != nil {
		log.Error().Msgf("Failed to lookup host: %s, err: %v", s.host, err)
	}
	storageServers := s.GetStorageServers()

	serversMap := make(map[string]struct{}, len(storageServers))
	for _, serv := range storageServers {
		serversMap[serv.Address] = struct{}{}
	}

	for _, addr := range addrs {
		addr = fmt.Sprintf("%s:%d", addr, s.port)
		if _, ok := serversMap[addr]; !ok {
			err = s.AddServer(ctx, addr)
			if err != nil {
				log.Error().Msgf("Failed to add server: %s err: %v", addr, err)
			}
		}
	}
	return nil
}

func (s *Service) startHeartbeat(ctx context.Context) {
	log.Print("Starting Heartbeat")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.heartbeatHandler(ctx)
	}
}

func (s *Service) monitorServers(ctx context.Context) {
	log.Print("Starting monitorServers")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.monitorHandler(ctx)
	}
}
