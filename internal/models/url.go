package models

import (
	"errors"

	"github.com/andranikuz/shortener/internal/config"
)

// URL структура. Основная модель данных приложения.
type URL struct {
	ID          string `json:"id"`
	FullURL     string `json:"full-url"`
	UserID      string `json:"user-id"`
	DeletedFlag bool   `json:"deleted-flag"`
}

// GetShorter метод для получения сокращенного URL.
func (url *URL) GetShorter() string {
	return config.Config.BaseURL + "/" + url.ID
}

// ErrURLAlreadyExists ошибка, получаемая при попытке сохранить в БД уже сущетсвующий URL.
var ErrURLAlreadyExists = errors.New(`url already exists`)

// ErrURLDeleted ошибка, получаемая при попытке получить удаленный URL.
var ErrURLDeleted = errors.New(`url deleted`)
