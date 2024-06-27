package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andranikuz/shortener/internal/api/rest/middlewares"
	"github.com/andranikuz/shortener/internal/container"
	"github.com/andranikuz/shortener/internal/services/shortener"
)

// HTTPHandler хендлер для обработки http запросов.
type HTTPHandler struct {
	shortener *shortener.Shortener
}

// NewHTTPHandler функция для инициализации NewHTTPHandler.
func NewHTTPHandler(cnt *container.Container) HTTPHandler {
	h := HTTPHandler{}
	h.shortener, _ = cnt.Shortener()

	return h
}

// Router метод получения роутера. Используется библиотека chi.
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
	r.Get("/api/user/urls", func(w http.ResponseWriter, r *http.Request) {
		h.GetUserURLsHandler(ctx, w, r)
	})
	r.Delete("/api/user/urls", func(w http.ResponseWriter, r *http.Request) {
		h.DeleteURLsHandler(ctx, w, r)
	})
	r.Mount("/debug", middleware.Profiler())

	return r
}
