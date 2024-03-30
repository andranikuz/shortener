package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andranikuz/shortener/internal/app/usecases"
)

type GenerateShortURLJSONHandlerRequest struct {
	URL string `json:"url"`
}

type GenerateShortURLJSONHandlerResponse struct {
	Result string `json:"result"`
}

func GenerateShortURLJSONHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var request GenerateShortURLJSONHandlerRequest
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	shortURL := usecases.GenerateShortURL(request.URL)
	if shortURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(GenerateShortURLJSONHandlerResponse{Result: shortURL})
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
}
