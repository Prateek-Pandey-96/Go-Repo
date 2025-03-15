package storage

import "github.com/prateek96/pendancySystem/entities"

type InMemoryStorage struct {
	entities map[int][]string
}

func GetInMemStorage() IStorage {
	return &InMemoryStorage{
		entities: make(map[int][]string),
	}
}

func (ims *InMemoryStorage) Add(entity entities.Entity) error {
	ims.entities[entity.Id] = entity.Tags
	return nil
}

func (ims *InMemoryStorage) Remove(id int) error {
	delete(ims.entities, id)
	return nil
}

func (ims *InMemoryStorage) GetCount(tags []string) (int, error) {
	count := 0
	for _, value := range ims.entities {
		if satisfy(tags, value) {
			count += 1
		}
	}
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
