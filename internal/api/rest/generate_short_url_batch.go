package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/services/shortener"
	"github.com/andranikuz/shortener/internal/utils/authorize"
)

// GenerateShortURLBatchHandlerRequest структура запроса.
type GenerateShortURLBatchHandlerRequest []shortener.OriginalItem

// GenerateShortURLBatchHandlerResponse структура ответа.
type GenerateShortURLBatchHandlerResponse []shortener.ShortenItem

// GenerateShortURLBatchHandler json хендлер создания массива сокращенных URLs.
func (h HTTPHandler) GenerateShortURLBatchHandler(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res.Header().Set("Content-Type", "application/json")
	if err := req.ParseForm(); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var request GenerateShortURLBatchHandlerRequest
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &request); err != nil {
		log.Info().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if len(request) == 0 {
		return
	}
	userID := authorize.GetOrGenerateUserID(res, req)
	response, err := h.shortener.GenerateShortURLBatch(ctx, request, userID)
	if err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
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
