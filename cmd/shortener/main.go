package main

import (
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	a, err := app.NewApplication()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Build version: " + buildVersion)
	log.Info().Msg("Build date: " + buildDate)
	log.Info().Msg("Build commit: " + buildCommit)
	if err := a.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
