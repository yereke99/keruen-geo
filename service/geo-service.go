package service

import (
	"keruen-geo/helper"
	"keruen-geo/models"
	"keruen-geo/store"
	"log"
)

type GeoService interface {
	Create(id int64, data models.Location) error
	Get() ([]*store.DataStore, error)
	GetAll() ([]models.Location, error)
	GetNearby(myGeo helper.Location) ([]models.Location, error)
}

type geoService struct {
	Store store.Store
}

func NewStoreService(st store.Store) *geoService {
	store := &geoService{
		Store: st,
	}

	return store
}

func (s *geoService) Create(id int64, data models.Location) error {
	if err := s.Store.Insert(id, data); err != nil {
		log.Println("error in geo save store map")
		return err
	}

	return nil
}

func (s *geoService) Get() ([]*store.DataStore, error) {
	results, err := s.Store.Take()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (s *geoService) GetAll() ([]models.Location, error) {
	results, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *geoService) GetNearby(myGeo helper.Location) ([]models.Location, error) {

	var nearby []models.Location

	results, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}

	for i, res := range results {
		if helper.CheckDistance(myGeo, res.Latitude, res.Longtitude) {
			nearby = append(nearby, results[i])
		}
	}

	return nearby, nil
}
