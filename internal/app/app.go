package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/api/handlers"
	"github.com/andranikuz/shortener/internal/api/middlewares"
	"github.com/andranikuz/shortener/internal/config"
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
	config.Init()
	if err := storage.Init(); err != nil {
		return err
	}

	return nil
}

func (app *Application) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.RequestCompressor)
	r.Post("/", handlers.GenerateShortURLHandler)
	r.Get("/{id}", handlers.GetFullURLHandler)
	r.Post("/api/shorten", handlers.GenerateShortURLJSONHandler)
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}
