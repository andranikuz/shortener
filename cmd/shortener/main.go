package main

import (
	"os"
	"os/signal"
	"syscall"

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

	// Канал для сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Запуск сервера в отдельной горутине
	go func() {
		if err := a.Run(); err != nil {
			log.Fatal().Msgf("Server failed:  %s", err.Error())
		}
	}()

	// Ожидание сигнала завершения
	sig := <-sigChan
	log.Info().Msgf("Server failed:  %s", sig)

	// Завершение сервера
	a.Stop()

	// Завершение процесса с кодом 0
	os.Exit(0)
}
