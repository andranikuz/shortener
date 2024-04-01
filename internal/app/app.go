package app

import (
	"net/http"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
)

type Application struct {
	DB DBInterface
}

type DBInterface interface {
	Get(id string) (*models.URL, error)
	Save(url models.URL) error
}

func NewApplication() (*Application, error) {
	a := Application{}
	if err := a.init(); err != nil {
		return nil, err
	}

	return &a, nil
}

func (app *Application) init() error {
	config.Init()
	filePath := config.Config.FileStoragePath
	var err error
	if app.DB == nil {
		if filePath != "" {
			app.DB, err = file.NewFileDB(filePath)
			if err != nil {
				return err
			}
		} else {
			app.DB, err = memory.NewMemoryDB()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *Application) Run(handler http.Handler) error {
	if err := http.ListenAndServe(config.Config.ServerAddress, handler); err != nil {
		return err
	}

	return nil
}
