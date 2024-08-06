package services

import (
	"context"

	"github.com/andranikuz/shortener/internal/models"
)

// Storage интерфейс репозитория.
type Shortener interface {
	DeleteURLs(ids []string, userID string)
	GenerateShortURL(ctx context.Context, fullURL string, userID string) (string, error)
	GenerateShortURLBatch(ctx context.Context, items []OriginalItem, userID string) ([]ShortenItem, error)
	GetFullURL(ctx context.Context, id string) (string, error)
	GetUserURLs(ctx context.Context, userID string) ([]models.URL, error)
	GetInternalStats(ctx context.Context) (int64, int64, error)
}

// OriginalItem структура оригинального URL.
type OriginalItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// ShortenItem структура сокращенного URL.
type ShortenItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
