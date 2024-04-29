package shortener

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (s *Shortener) GetFullURL(ctx context.Context, id string) string {
	url, err := s.storage.Get(ctx, id)
	if err != nil {
		log.Info().Msg(err.Error())
		return ""
	}

	return url.FullURL
}
