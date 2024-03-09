package app

import (
	"github.com/andranikuz/shortener/internal/handlers"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Application struct {
}

func (app *Application) Init() {
	storage.Init()
}

func (app *Application) Router() chi.Router {
	r := chi.NewRouter()
	r.Post("/", handlers.GenerateShortURLHandler)
	r.Get("/{id}", handlers.GetFullURLHandler)

	return r
}
