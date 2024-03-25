package handlers

import (
	"encoding/json"
	"github.com/andranikuz/shortener/internal/app/usecases"
	"io"
	"net/http"
)

type GetShortenHandlerRequest struct {
	Url string `json:"url"`
}

type GetShortenHandlerResponse struct {
	Result string `json:"result"`
}

func GetShortenByFullUrlJSONHandler(res http.ResponseWriter, req *http.Request) {
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
	url, err := usecases.GetURLByFullURL(request.Url)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(GetShortenHandlerResponse{Result: url.GetShorter()})
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write(resp); err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
}
