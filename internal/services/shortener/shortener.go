package shortener

import (
	"github.com/andranikuz/shortener/internal/storage"
)

type Shortener struct {
	storage storage.Storage
}

func NewShortener(storage storage.Storage) *Shortener {
	return &Shortener{storage: storage}
}
