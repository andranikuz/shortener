package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/api/handlers"
	"github.com/andranikuz/shortener/internal/api/middlewares"
	"github.com/andranikuz/shortener/internal/app"
)

func Router(a app.Application) chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.RequestCompressor)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateShortURLHandler(w, r, a)
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetFullURLHandler(w, r, a)
	})
	r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateShortURLJSONHandler(w, r, a)
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		handlers.PingHandler(w, a)
	})
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}
