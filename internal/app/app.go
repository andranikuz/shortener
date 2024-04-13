package app

import (
	"context"
	"net/http"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage/file"
	"github.com/andranikuz/shortener/internal/storage/memory"
	"github.com/andranikuz/shortener/internal/storage/postgres"
)

type Application struct {
	DB  DBInterface
	CTX context.Context
}

type DBInterface interface {
	Get(ctx context.Context, id string) (*models.URL, error)
	Save(ctx context.Context, url models.URL) error
	SaveBatch(ctx context.Context, urls []models.URL) error
	Migrate(ctx context.Context) error
}

func NewApplication() (*Application, error) {
	a := Application{
		CTX: context.Background(),
	}
	config.Init()
	var err error
	if config.Config.DatabaseDSN != "" {
		a.DB, err = postgres.NewPostgresDB(config.Config.DatabaseDSN)
		if err != nil {
			return nil, err
		}
	} else if config.Config.FileStoragePath != "" {
		a.DB, err = file.NewFileDB(config.Config.FileStoragePath)
		if err != nil {
			return nil, err
		}
	} else {
		a.DB, err = memory.NewMemoryDB()
		if err != nil {
			return nil, err
		}
	}
	if err = a.DB.Migrate(a.CTX); err != nil {
		return nil, err
	}

	return &a, nil
}

func (app *Application) Run(handler http.Handler) error {
	if err := http.ListenAndServe(config.Config.ServerAddress, handler); err != nil {
		return err
	}

	return nil
}
