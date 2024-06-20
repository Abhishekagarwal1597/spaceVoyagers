package store

import (
	"errors"
	"spaceVoyagers/models"
	"sync"

	"github.com/sirupsen/logrus"
)

type Exoplanet struct {
	sync.RWMutex
	exoplanets map[string]models.ExoplanetModel
}

func NewExoplanet() *Exoplanet {
	return &Exoplanet{
		exoplanets: make(map[string]models.ExoplanetModel),
	}
}

var logger = logrus.New()

func (store *Exoplanet) AddExoplanet(addExoplanet models.ExoplanetModel) {
	store.Lock()
	defer store.Unlock()
	store.exoplanets[addExoplanet.ID] = addExoplanet
}

func (store *Exoplanet) UpdateExoplanet(key string, UpdateExoplanet models.ExoplanetModel) error {
	store.Lock()
	defer store.Unlock()
	if _, exist := store.exoplanets[key]; !exist {
		logger.Errorf("Exoplanet with key %s not found", key)
		return errors.New("exoplanet not found")
	}
	store.exoplanets[key] = UpdateExoplanet
	return nil
}

func (store *Exoplanet) DeleteExoplanet(key string) error {
	store.Lock()
	defer store.Unlock()
	if _, exist := store.exoplanets[key]; !exist {
		logger.Errorf("Exoplanet with key %s not found", key)
		return errors.New("exoplanet not found")
	}
	delete(store.exoplanets, key)
	return nil
}

func (store *Exoplanet) ListExoplanet() interface{} {
	store.RLock()
	defer store.RUnlock()
	exoplanets := make([]models.ExoplanetModel, 0, len(store.exoplanets))

	for _, exoplanet := range store.exoplanets {
		exoplanets = append(exoplanets, exoplanet)
	}
	return exoplanets
}

func (store *Exoplanet) GetExoplanetById(key string) (models.ExoplanetModel, error) {
	store.Lock()
	defer store.Unlock()
	if _, exist := store.exoplanets[key]; !exist {
		logger.Errorf("Exoplanet with key %s not found", key)
		return models.ExoplanetModel{}, errors.New("exoplanet not found")
	}
	return store.exoplanets[key], nil
}
