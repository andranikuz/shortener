package shortener

import (
	"context"

	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/services"
)

// GenerateShortURLBatch метод создания массива сокращенных ссылок.
func (s Shortener) GenerateShortURLBatch(ctx context.Context, items []services.OriginalItem, userID string) ([]services.ShortenItem, error) {
	var urls []models.URL
	var result []services.ShortenItem
	var url models.URL
	for _, item := range items {
		url = models.URL{ID: item.CorrelationID, FullURL: item.OriginalURL, UserID: userID}
		urls = append(urls, url)
		result = append(result, services.ShortenItem{item.CorrelationID, url.GetShorter()})
	}

	if err := s.storage.SaveBatch(ctx, urls); err != nil {
		return result, err
	}

	return result, nil
}
