package usecases

import (
	"github.com/andranikuz/shortener/internal/storage"
)

func GetFullUrl(id string) string {
	url, err := storage.Get(id)
	if err != nil {
		return ""
	}

	return url.Url
}
