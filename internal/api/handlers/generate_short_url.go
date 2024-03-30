package handlers

import (
	"io"
	"net/http"

	"github.com/andranikuz/shortener/internal/app/usecases"
)

func GenerateShortURLHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(req.Body)

	fullURL := string(body)
	if fullURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := usecases.GenerateShortURL(fullURL)

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "text/plain")
	if _, err := io.WriteString(res, shortURL); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
