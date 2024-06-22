package service

import (
	"errors"
	"sync"

	pb "github.com/prateek69/go-grpc/pb/proto"
)

var ErrorLaptopAlreadyExists = errors.New("laptop already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrorLaptopAlreadyExists
	}

	store.data[laptop.Id] = laptop
	return nil
}

// Todo :- Implement to further store the document in mongo db
type DBStore struct {
}
