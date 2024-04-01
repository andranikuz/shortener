package app

import (
	"fmt"
	"net/http"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
)

type Application struct {
	isInit bool
	DB     DBInterface
}

type DBInterface interface {
	Get(id string) (*models.URL, error)
	Save(url models.URL) error
}

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
	app.isInit = true

	return nil
}

func (app *Application) Run(handler http.Handler) error {
	if !app.isInit {
		return fmt.Errorf("running not inited application")
	}

	if err := http.ListenAndServe(config.Config.ServerAddress, handler); err != nil {
		return err
	}

	return nil
}
