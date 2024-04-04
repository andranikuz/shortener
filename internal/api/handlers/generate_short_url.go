package handlers

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/usecases"
)

func GenerateShortURLHandler(res http.ResponseWriter, req *http.Request, a app.Application) {
	if err := req.ParseForm(); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(req.Body)

	fullURL := string(body)
	if fullURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := usecases.GenerateShortURL(a, fullURL)

	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "text/plain")
	if _, err := io.WriteString(res, shortURL); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
