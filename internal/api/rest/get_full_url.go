package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h HTTPHandler) GetFullURLHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		log.Info().Msg("empty id")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fullURL := h.shortener.GetFullURL(ctx, id)
	if fullURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", fullURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
