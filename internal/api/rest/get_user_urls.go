package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/utils/authorize"
)

// GetUserURLsHandlerItem структура одного URL в ответе.
type GetUserURLsHandlerItem struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// GetUserURLsHandler хендлер получения списка пользователей.
func (h HTTPHandler) GetUserURLsHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	userID, err := authorize.GetUserID(req)
	if err != nil || userID == "" {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	urls, err := h.shortener.GetUserURLs(ctx, userID)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var response []GetUserURLsHandlerItem
	for _, url := range urls {
		response = append(response, GetUserURLsHandlerItem{ShortURL: url.GetShorter(), OriginalURL: url.FullURL})
	}

	if len(response) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}
	res.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(response)
	if err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
	}
}
