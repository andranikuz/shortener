package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/usecases"
)

type GenerateShortURLJSONHandlerRequest struct {
	URL string `json:"url"`
}

type GenerateShortURLJSONHandlerResponse struct {
	Result string `json:"result"`
}

func GenerateShortURLJSONHandler(res http.ResponseWriter, req *http.Request, a app.Application) {
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
	shortURL := usecases.GenerateShortURL(a, request.URL)
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
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
	}
}
