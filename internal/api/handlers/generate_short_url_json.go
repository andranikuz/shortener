package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
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
	code := http.StatusCreated
	shortURL, err := usecases.GenerateShortURL(a, request.URL)
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
