package client

import (
	"sync"

	"github.com/pashest/object-storage-service/internal/client/helper"
	"github.com/pashest/object-storage-service/internal/client/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientPool struct {
	mu          sync.RWMutex
	connections map[string]*grpc.ClientConn
	helperPool  map[string]*helper.Client
	storagePool map[string]*storage.Client
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		connections: make(map[string]*grpc.ClientConn),
		helperPool:  make(map[string]*helper.Client),
		storagePool: make(map[string]*storage.Client),
	}
}

// AddConnection add new connection
func (p *ClientPool) AddConnection(address string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.connections[address]; exists {
		return nil
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	p.connections[address] = conn
	p.helperPool[address] = helper.New(conn)
	p.storagePool[address] = storage.New(conn)
	return nil
}

// RemoveConnection remove connection
func (p *ClientPool) RemoveConnection(address string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if conn, exists := p.connections[address]; exists {
		conn.Close()
		delete(p.connections, address)
		delete(p.helperPool, address)
		delete(p.storagePool, address)
	}
}

// GetHelperClient get available helper client
func (p *ClientPool) GetHelperClient(address string) (*helper.Client, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	client, exists := p.helperPool[address]
	return client, exists
}

// GetStorageClient get available storage client
func (p *ClientPool) GetStorageClient(address string) (*storage.Client, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	client, exists := p.storagePool[address]
	return client, exists
}
