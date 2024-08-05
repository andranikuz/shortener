package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andranikuz/shortener/internal/api/rest/middlewares"
	"github.com/andranikuz/shortener/internal/container"
	"github.com/andranikuz/shortener/internal/services"
	"github.com/andranikuz/shortener/internal/storage"
)

// HTTPHandler хендлер для обработки http запросов.
type HTTPHandler struct {
	shortener services.Shortener
	storage   storage.Storage
}

// NewHTTPHandler функция для инициализации NewHTTPHandler.
func NewHTTPHandler(cnt *container.Container) HTTPHandler {
	h := HTTPHandler{}
	h.shortener, _ = cnt.Shortener()
	h.storage, _ = cnt.Storage()

	return h
}

// Router метод получения роутера. Используется библиотека chi.
func (h HTTPHandler) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.RequestLogger)
	r.Use(middlewares.RequestCompressor)
	r.Post("/", h.GenerateShortURLHandler)
	r.Get("/{id}", h.GetFullURLHandler)
	r.Post("/api/shorten", h.GenerateShortURLJSONHandler)
	r.Post("/api/shorten/batch", h.GenerateShortURLBatchHandler)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		h.PingHandler(w)
	})
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/api/user/urls", h.GetUserURLsHandler)
	r.Delete("/api/user/urls", h.DeleteURLsHandler)
	r.Get("/api/internal/stats", h.GetInternalStats)
	r.Mount("/debug", middleware.Profiler())

	return r
}
