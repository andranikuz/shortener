package usecases

import (
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage"
	"github.com/andranikuz/shortener/internal/utils/generator"
)

func GenerateShortURL(fullURL string) string {
	id := generator.GenerateUniqueID()
	url := models.URL{ID: id, FullURL: fullURL}
	if err := storage.Save(url); err != nil {
		return ""
	}

	return url.GetShorter()
}
