package usecases

import (
	"github.com/andranikuz/shortener/internal/app"
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

func GenerateShortURLBatch(a app.Application, items []OriginalItem) ([]ShortenItem, error) {
	var urls []models.URL
	var result []ShortenItem
	var url models.URL
	for _, item := range items {
		url = models.URL{item.CorrelationID, item.OriginalURL}
		urls = append(urls, url)
		result = append(result, ShortenItem{item.CorrelationID, url.GetShorter()})
	}

	if err := a.DB.SaveBatch(a.CTX, urls); err != nil {
		return result, err
	}

	return result, nil
}
