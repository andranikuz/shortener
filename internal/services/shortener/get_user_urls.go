package shortener

import (
	"context"

	"github.com/andranikuz/shortener/internal/models"

	"github.com/rs/zerolog/log"
)

// GetUserURLs метод получения URLs пользователя.
func (s *Shortener) GetUserURLs(ctx context.Context, userID string) ([]models.URL, error) {
	urls, err := s.storage.GetByUserID(ctx, userID)
	if err != nil {
		log.Info().Msg(err.Error())
		return nil, err
	}

	return urls, nil
}
