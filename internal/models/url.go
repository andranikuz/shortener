package models

import (
	"errors"

	"github.com/andranikuz/shortener/internal/config"
)

type URL struct {
	ID          string `json:"id"`
	FullURL     string `json:"full-url"`
	UserID      string `json:"user-id"`
	DeletedFlag bool   `json:"deleted-flag"`
}

func (url *URL) GetShorter() string {
	return config.Config.BaseURL + "/" + url.ID
}

var ErrURLAlreadyExists = errors.New(`url already exists`)

var ErrURLDeleted = errors.New(`url deleted`)
