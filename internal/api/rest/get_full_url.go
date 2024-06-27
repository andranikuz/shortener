package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
)

// GetFullURLHandler хендлер редиректа с сокращенного на полный URL.
func (h HTTPHandler) GetFullURLHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		log.Info().Msg("empty id")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fullURL, err := h.shortener.GetFullURL(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrURLDeleted) {
			log.Info().Msg(err.Error())
			res.WriteHeader(http.StatusGone)
			return
		} else {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	res.Header().Set("Location", fullURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
