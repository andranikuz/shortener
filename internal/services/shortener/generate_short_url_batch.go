package shortener

import (
	"context"

	"github.com/andranikuz/shortener/internal/models"
)

type OriginalItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ShortenItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (s *Shortener) GenerateShortURLBatch(ctx context.Context, items []OriginalItem) ([]ShortenItem, error) {
	var urls []models.URL
	var result []ShortenItem
	var url models.URL
	for _, item := range items {
		url = models.URL{ID: item.CorrelationID, FullURL: item.OriginalURL}
		urls = append(urls, url)
		result = append(result, ShortenItem{item.CorrelationID, url.GetShorter()})
	}

	if err := s.storage.SaveBatch(ctx, urls); err != nil {
		return result, err
	}

	return result, nil
}
