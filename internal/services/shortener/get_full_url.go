package shortener

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
)

// GetFullURL метод получения полного URL.
func (s Shortener) GetFullURL(ctx context.Context, id string) (string, error) {
	url, err := s.storage.Get(ctx, id)
	if err != nil {
		log.Info().Msg(err.Error())
		return "", err
	}

	if url.DeletedFlag {
		return "", models.ErrURLDeleted
	}

	return url.FullURL, nil
}
