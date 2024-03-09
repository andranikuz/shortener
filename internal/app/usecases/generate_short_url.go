package usecases

import (
	"github.com/andranikuz/shortener/internal/app/services/generator"
	"github.com/andranikuz/shortener/internal/storage"
)

func GenerateShortUrl(fullUrl string) string {
	id := generator.GenerateUniqueId()
	url := storage.Url{Id: id, Url: fullUrl}
	if err := storage.Save(url); err != nil {
		return ""
	}

	return "http://localhost:8080/" + id
}
