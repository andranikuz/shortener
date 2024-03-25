package handlers

import (
	"encoding/json"
	"github.com/andranikuz/shortener/internal/app/usecases"
	"io"
	"net/http"
)

type GetShortenHandlerRequest struct {
	URL string `json:"url"`
}

type GetShortenHandlerResponse struct {
	Result string `json:"result"`
}

func GetShortenByFullURLJSONHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var request GetShortenHandlerRequest
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
	resp, err := json.Marshal(GetShortenHandlerResponse{Result: shortURL})
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
}
