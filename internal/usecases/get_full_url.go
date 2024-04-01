package usecases

import (
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
)

func GetFullURL(a app.Application, id string) string {
	url, err := a.DB.Get(id)
	if err != nil {
		log.Info().Msg(err.Error())
		return ""
	}

	return url.FullURL
}
