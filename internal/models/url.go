package models

import "github.com/andranikuz/shortener/internal/config"

type URL struct {
	ID      string
	FullURL string
}

func (url *URL) GetShorter() string {
	return config.Config.BaseURL + "/" + url.ID
}
