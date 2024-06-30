package storage

import (
	"context"

	"github.com/andranikuz/shortener/internal/models"
)

// Storage интерфейс репозитория.
type Storage interface {
	Save(ctx context.Context, url models.URL) error
	Get(ctx context.Context, fullURL string) (*models.URL, error)
	GetByUserID(ctx context.Context, userID string) ([]models.URL, error)
	GetByFullURL(ctx context.Context, id string) (*models.URL, error)
	SaveBatch(ctx context.Context, urls []models.URL) error
	DeleteURLs(ctx context.Context, ids []string, userID string) error
}
