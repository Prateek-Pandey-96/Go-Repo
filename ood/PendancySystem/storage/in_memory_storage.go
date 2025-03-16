package storage

import (
	"sync"

	"github.com/prateek96/pendancySystem/entities"
)

type InMemoryStorage struct {
	entities map[int][]string
	mu       sync.Mutex
}

func GetInMemStorage() IStorage {
	return &InMemoryStorage{
		entities: make(map[int][]string),
		mu:       sync.Mutex{},
	}
}

func (ims *InMemoryStorage) Add(entity entities.Entity) error {
	ims.mu.Lock()
	ims.entities[entity.Id] = entity.Tags
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) Remove(id int) error {
	ims.mu.Lock()
	delete(ims.entities, id)
	defer ims.mu.Unlock()
	return nil
}

func (ims *InMemoryStorage) GetCount(tags []string) (int, error) {
	ims.mu.Lock()
	count := 0
	for _, value := range ims.entities {
		if satisfy(tags, value) {
			count += 1
		}
	}
	defer ims.mu.Unlock()
	return count, nil
}

func satisfy(tags []string, entityTags []string) bool {
	i, j := 0, 0
	m, n := len(tags), len(entityTags)
	for i < m && j < n {
		if tags[i] == entityTags[j] {
			i += 1
			j += 1
		} else {
			break
		}
	}
	return i == m
}
