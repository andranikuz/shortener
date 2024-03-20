package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/handlers"
	"github.com/andranikuz/shortener/internal/logger"
	"github.com/andranikuz/shortener/internal/storage"
)

type Application struct {
}

func (app *Application) Run() error {
	if err := app.Init(); err != nil {
		return err
	}
	if err := http.ListenAndServe(config.Config.ServerAddress, app.Router()); err != nil {
		return err
	}

	return nil
}

func (app *Application) Init() error {
	if err := storage.Init(); err != nil {
		return err
	}
	config.Init()
	if err := logger.Init(); err != nil {
		return err
	}

	return nil
}

func (app *Application) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(logger.RequestLogger)
	r.Post("/", handlers.GenerateShortURLHandler)
	r.Get("/{id}", handlers.GetFullURLHandler)
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}
