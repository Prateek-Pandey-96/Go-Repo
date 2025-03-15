package storage

import "github.com/prateek96/pendancySystem/entities"

type IStorage interface {
	Add(entity entities.Entity) error
	Remove(id int) error
	GetCount(tags []string) (int, error)
}
