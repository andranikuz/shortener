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

type GenerateShortURLBatchHandlerRequest []shortener.OriginalItem

type GenerateShortURLBatchHandlerResponse []shortener.ShortenItem

func (h HTTPHandler) GenerateShortURLBatchHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
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
	userID := authorize.GetOrGenerateUserId(res, req)
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
