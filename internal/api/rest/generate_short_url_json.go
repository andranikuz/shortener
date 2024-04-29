package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/utils/authorize"
)

type GenerateShortURLJSONHandlerRequest struct {
	URL string `json:"url"`
}

type GenerateShortURLJSONHandlerResponse struct {
	Result string `json:"result"`
}

func (h HTTPHandler) GenerateShortURLJSONHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if err := req.ParseForm(); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var request GenerateShortURLJSONHandlerRequest
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &request); err != nil {
		log.Info().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	code := http.StatusCreated
	userID := authorize.GetOrGenerateUserId(res, req)
	shortURL, err := h.shortener.GenerateShortURL(ctx, request.URL, userID)
	if err != nil {
		if errors.Is(err, models.ErrURLAlreadyExists) {
			code = http.StatusConflict
		} else {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if shortURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(GenerateShortURLJSONHandlerResponse{Result: shortURL})
	if err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(code)
	if _, err := res.Write(resp); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
	}
}
