package usecases

import (
	"fmt"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/storage"
)

func GetURLByFullURL(fullURL string) (*models.URL, error) {
	url, err := storage.GetByFullURL(fullURL)
	if err != nil {
		return nil, fmt.Errorf("get shorten by fullUrl=%s error=%s", fullURL, err)
	}

	return url, nil
}
