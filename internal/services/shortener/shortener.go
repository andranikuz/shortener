package shortener

import (
	"github.com/andranikuz/shortener/internal/storage"
)

// Shortener сервис. Хранит в себе реализацию бизнес логики приложения.
type Shortener struct {
	storage storage.Storage
}

// NewShortener функция для инициализации Shortener.
func NewShortener(storage storage.Storage) *Shortener {
	return &Shortener{storage: storage}
}
