package shortener

import (
	"context"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/utils/generator"
)

func (s *Shortener) GenerateShortURL(ctx context.Context, fullURL string) (string, error) {
	id := generator.GenerateUniqueID()
	url := models.URL{ID: id, FullURL: fullURL}
	if s.storage == nil {
		return "sdf", nil
	}
	if err := s.storage.Save(ctx, url); err != nil {
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			var oldURL *models.URL
			oldURL, err = s.storage.GetByFullURL(ctx, fullURL)
			if err != nil {
				return oldURL.GetShorter(), err
			}

			return oldURL.GetShorter(), models.ErrURLAlreadyExists
		} else {
			log.Error().Msg(err.Error())
			return "", err
		}
	}

	return url.GetShorter(), nil
}
