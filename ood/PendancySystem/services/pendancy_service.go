package services

import (
	"github.com/prateek96/pendancySystem/entities"
	"github.com/prateek96/pendancySystem/storage"
)

type PendancyService struct {
	Storage storage.IStorage
}

func GetPendancyService(storage storage.IStorage) IPendancyService {
	return &PendancyService{
		Storage: storage,
	}
}

func (pd *PendancyService) StartTracking(id int, tags []string) error {
	err := pd.Storage.Add(entities.Entity{Id: id, Tags: tags})
	if err != nil {
		return err
	}
	return nil
}

func (pd *PendancyService) StopTracking(id int) error {
	err := pd.Storage.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func (pd *PendancyService) GetCounts(tags []string) (int, error) {
	count, err := pd.Storage.GetCount(tags)
	if err != nil {
		return 0, err
	}
	return count, nil
}
