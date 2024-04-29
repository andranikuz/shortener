package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/api/rest/middlewares"
	"github.com/andranikuz/shortener/internal/container"
	"github.com/andranikuz/shortener/internal/services/shortener"
)

type HTTPHandler struct {
	shortener *shortener.Shortener
}

func NewHTTPHandler(cnt *container.Container) HTTPHandler {
	h := HTTPHandler{}
	h.shortener, _ = cnt.Shortener()

	return h
}

func (h HTTPHandler) Router(ctx context.Context) chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.RequestCompressor)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		h.GenerateShortURLHandler(ctx, w, r)
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		h.GetFullURLHandler(ctx, w, r)
	})
	r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		h.GenerateShortURLJSONHandler(ctx, w, r)
	})
	r.Post("/api/shorten/batch", func(w http.ResponseWriter, r *http.Request) {
		h.GenerateShortURLBatchHandler(ctx, w, r)
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		h.PingHandler(w)
	})
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}
