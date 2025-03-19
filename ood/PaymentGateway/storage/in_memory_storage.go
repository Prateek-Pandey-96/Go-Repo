package storage

import (
	"sync"

	"github.com/prateek96/paymentGateway/entities"
)

type InMemoryStorage struct {
	clients  map[string]*entities.Client
	paymodes map[entities.Mode]struct{}
	mu       sync.Mutex
}

func GetInMemStorage() IStorage {
	return &InMemoryStorage{
		clients:  make(map[string]*entities.Client),
		paymodes: make(map[entities.Mode]struct{}),
		mu:       sync.Mutex{},
	}
}

func (ims *InMemoryStorage) AddClient(client *entities.Client) error {
	ims.mu.Lock()
	ims.clients[client.Id] = client
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) RemoveClient(client *entities.Client) error {
	ims.mu.Lock()
	delete(ims.clients, client.Id)
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) HasClient(client *entities.Client) (bool, error) {
	ims.mu.Lock()
	_, ok := ims.clients[client.Id]
	if !ok {
		return false, nil
	}
	defer ims.mu.Unlock()
	return true, nil
}

func (ims *InMemoryStorage) AddPaymode(paymode entities.Mode) error {
	ims.mu.Lock()
	ims.paymodes[paymode] = struct{}{}
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) RemovePaymode(paymode entities.Mode) error {
	ims.mu.Lock()
	delete(ims.paymodes, paymode)
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) ListPaymodes() ([]entities.Mode, error) {
	ims.mu.Lock()
	paymodes := make([]entities.Mode, len(ims.paymodes))
	for paymode := range ims.paymodes {
		paymodes = append(paymodes, paymode)
	}
	defer ims.mu.Unlock()
	return paymodes, nil
}
