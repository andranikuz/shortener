package usecases

import (
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/models"
	"github.com/andranikuz/shortener/internal/utils/generator"
)

func GenerateShortURL(a app.Application, fullURL string) string {
	id := generator.GenerateUniqueID()
	url := models.URL{ID: id, FullURL: fullURL}
	if err := a.DB.Save(a.CTX, url); err != nil {
		log.Error().Msg(err.Error())
		return ""
	}

	return url.GetShorter()
}
