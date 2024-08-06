package shortener

import (
	"context"

	"github.com/rs/zerolog/log"
)

// GetInternalStats метод получения статистики сервиса.
func (s Shortener) GetInternalStats(ctx context.Context) (int64, int64, error) {
	urls, err := s.storage.GetURLsCount(ctx)
	if err != nil {
		log.Info().Msg(err.Error())
		return 0, 0, err
	}

	users, err := s.storage.GetUsersCount(ctx)
	if err != nil {
		log.Info().Msg(err.Error())
		return 0, 0, err
	}

	return urls, users, nil
}
