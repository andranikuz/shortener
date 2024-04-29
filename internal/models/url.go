package models

import (
	"errors"

	"github.com/andranikuz/shortener/internal/config"
)

type URL struct {
	ID      string `json:"id"`
	FullURL string `json:"full-url"`
}

func (url *URL) GetShorter() string {
	return config.Config.BaseURL + "/" + url.ID
}

var ErrURLAlreadyExists = errors.New(`url already exists`)
