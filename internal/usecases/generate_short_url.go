package usecases

import (
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/utils/generator"
)

func GenerateShortURL(a app.Application, fullURL string) (string, error) {
	id := generator.GenerateUniqueID()
	url := models.URL{ID: id, FullURL: fullURL}
	if err := a.DB.Save(a.CTX, url); err != nil {
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			var oldURL *models.URL
			oldURL, err = a.DB.GetByFullURL(a.CTX, fullURL)
			if err != nil {
				return oldURL.GetShorter(), err
			}

			return oldURL.GetShorter(), models.URLAlreadyExistsError
		} else {
			log.Error().Msg(err.Error())
			return "", err
		}
	}

	return url.GetShorter(), nil
}
