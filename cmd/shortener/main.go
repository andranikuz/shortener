package main

import (
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
)

func main() {
	a, err := app.NewApplication()
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	if err := a.Run(); err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
}
