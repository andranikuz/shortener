package storage

import (
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
)

type DBInterface interface {
	Get(id string) (*models.URL, error)
	Save(url models.URL) error
}

var db DBInterface

func Init() error {
	filePath := config.Config.FileStoragePath
	var err error
	if filePath != "" {
		db, err = file.NewFileDB(filePath)
		if err != nil {
			return err
		}
	} else {
		db, err = memory.NewMemoryDB()
		if err != nil {
			return err
		}
	}

	return nil
}

func Get(id string) (*models.URL, error) {
	return db.Get(id)
}

func Save(url models.URL) error {
	return db.Save(url)
}
