package main

import (
	"fmt"

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
		log.Error().Msg(err.Error())
		panic(err)
	}
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	if err := a.Run(); err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
}
