package store

import (
	"keruen-geo/models"
	"sync"
)

type DataStore map[int64]models.Location

type Store struct {
	Storage map[int64]models.Location
	Mu      *sync.RWMutex
}

func NewStore(storage DataStore) Store {
	store := Store{
		Storage: storage,
		Mu:      &sync.RWMutex{},
	}

	return store
}

func (s *Store) Insert(id int64, data models.Location) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Storage[id] = data

	return nil
}

func (s *Store) Take() ([]*DataStore, error) {
	var locations []*DataStore

	s.Mu.Lock()
	defer s.Mu.Unlock()

	for key, value := range s.Storage {
		copy := value
		locations = append(locations, &DataStore{key: copy})
	}

	return locations, nil
}

func (s *Store) GetAll() ([]models.Location, error) {
	var location []models.Location

	s.Mu.Lock()
	defer s.Mu.Unlock()

	for key, _ := range s.Storage {
		location = append(location, s.Storage[key])
	}

	return location, nil
}
