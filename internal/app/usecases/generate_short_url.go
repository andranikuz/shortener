package usecases

import (
	"github.com/andranikuz/shortener/internal/app/services/generator"
	"github.com/andranikuz/shortener/internal/storage"
)

func GenerateShortURL(fullURL string) string {
	id := generator.GenerateUniqueID()
	url := storage.URL{ID: id, FullURL: fullURL}
	if err := storage.Save(url); err != nil {
		return ""
	}

	return "http://localhost:8080/" + id
}
