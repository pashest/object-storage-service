package client

import (
	"sync"

	"github.com/pashest/object-storage-service/internal/client/helper"
	"github.com/pashest/object-storage-service/internal/client/storage"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const maxMessageSize = 2 * 1024 * 1024 * 1024 // 2 GB

type ConnectionPool struct {
	mu          sync.RWMutex
	connections map[string]*grpc.ClientConn
	helperPool  map[string]*helper.Client
	storagePool map[string]*storage.Client
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[string]*grpc.ClientConn),
		helperPool:  make(map[string]*helper.Client),
		storagePool: make(map[string]*storage.Client),
	}
}

// AddConnection add new connection
func (p *ConnectionPool) AddConnection(address string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.connections[address]; exists {
		return nil
	}

	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMessageSize)),
	)
	if err != nil {
		return err
	}

	p.connections[address] = conn
	p.helperPool[address] = helper.New(conn)
	p.storagePool[address] = storage.New(conn)
	return nil
}

// RemoveConnection remove connection
func (p *ConnectionPool) RemoveConnection(address string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if conn, exists := p.connections[address]; exists {
		err := conn.Close()
		if err != nil {
			log.Err(err).Msgf("Failed to close connection: %s", address)
		}
		delete(p.connections, address)
		delete(p.helperPool, address)
		delete(p.storagePool, address)
	}
}

func (p *ConnectionPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for addr, conn := range p.connections {
		if conn != nil {
			conn.Close()
			delete(p.connections, addr)
		}
	}

	p.helperPool = nil
	p.storagePool = nil
}

// GetHelperClient get available helper client
func (p *ConnectionPool) GetHelperClient(address string) (*helper.Client, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	client, exists := p.helperPool[address]
	return client, exists
}

// GetStorageClient get available storage client
func (p *ConnectionPool) GetStorageClient(address string) (*storage.Client, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	client, exists := p.storagePool[address]
	return client, exists
}
