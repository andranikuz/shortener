package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/usecases"
)

func GetFullURLHandler(res http.ResponseWriter, req *http.Request, a app.Application) {
	id := chi.URLParam(req, "id")
	if id == "" {
		log.Info().Msg("empty id")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fullURL := usecases.GetFullURL(a, id)
	if fullURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", fullURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
