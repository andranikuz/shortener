package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/api/handlers"
	"github.com/andranikuz/shortener/internal/api/middlewares"
)

func Router() chi.Router {
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
