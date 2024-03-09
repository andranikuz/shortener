package handlers

import (
	"github.com/andranikuz/shortener/internal/app/usecases"
	"io"
	"net/http"
)

func GenerateShortURLHandler(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI != "/" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
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
	if _, err := res.Write([]byte(shortURL)); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
