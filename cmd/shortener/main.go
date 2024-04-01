package main

import (
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/api"
	"github.com/andranikuz/shortener/internal/app"
)

func main() {
	a, err := app.NewApplication()
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	if err := a.Run(api.Router(*a)); err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
}
