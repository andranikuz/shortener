package usecases

import (
	"github.com/andranikuz/shortener/internal/storage"
)

func GetFullURL(id string) string {
	url, err := storage.Get(id)
	if err != nil {
		return ""
	}

	return url.FullURL
}
