package app

import (
	"github.com/andranikuz/shortener/internal/handlers"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
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
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}
