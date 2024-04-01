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

var App *Application

func (app *Application) Init() error {
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
	err := app.Init()
	if err != nil {
		return err
	}
	App = app

	if err := http.ListenAndServe(config.Config.ServerAddress, handler); err != nil {
		return err
	}

	return nil
}
