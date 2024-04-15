package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
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

	code := http.StatusCreated
	shortURL, err := usecases.GenerateShortURL(a, fullURL)
	if err != nil {
		if errors.Is(err, models.ErrURLAlreadyExists) {
			code = http.StatusConflict
		} else {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	res.WriteHeader(code)
	res.Header().Set("Content-Type", "text/plain")
	if _, err := io.WriteString(res, shortURL); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
