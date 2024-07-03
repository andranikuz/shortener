package container

import (
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/services"
	"github.com/andranikuz/shortener/internal/services/shortener"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
	"github.com/andranikuz/shortener/internal/storage/postgres"
)

// Container структура отвечает за контейнеризацию сервисов и репозиториев приложения.
type Container struct {
	storage   storage.Storage
	shortener services.Shortener
}

// NewContainer создает новый контейнер.
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

// Storage возвращает storage.Storage.
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

// Shortener возвращает сервис Shortener.
func (c Container) Shortener() (services.Shortener, error) {
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
