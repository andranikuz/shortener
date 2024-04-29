package container

import (
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/services/shortener"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
	"github.com/andranikuz/shortener/internal/storage/postgres"
)

type Container struct {
	storage   storage.Storage
	shortener *shortener.Shortener
}

func NewContainer() (*Container, error) {
	var err error
	var cnt Container
	cnt.storage, err = cnt.Storage()
	if err != nil {
		return nil, err
	}
	cnt.shortener, err = cnt.Shortener()
	if err != nil {
		return nil, err
	}

	return &cnt, nil
}

func (c Container) Storage() (storage.Storage, error) {
	if c.storage == nil {
		if config.Config.DatabaseDSN != "" {
			storage, err := postgres.NewPostgresStorage(config.Config.DatabaseDSN)
			if err != nil {
				return nil, err
			}
			storage.Migrate()
			c.storage = storage
		} else if config.Config.FileStoragePath != "" {
			storage, err := file.NewFileStorage(config.Config.FileStoragePath)
			if err != nil {
				return nil, err
			}
			c.storage = storage
		} else {
			storage, err := memory.NewMemoryStorage()
			if err != nil {
				return nil, err
			}
			c.storage = storage
		}
	}

	return c.storage, nil
}

func (c Container) Shortener() (*shortener.Shortener, error) {
	if c.shortener == nil {
		var err error
		storage, err := c.Storage()
		if err != nil {
			return nil, err
		}
		c.shortener = shortener.NewShortener(storage)
	}

	return c.shortener, nil
}
