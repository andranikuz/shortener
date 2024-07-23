package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Канал для завершения сервера
	shutdownChan := make(chan struct{})

	// Запуск сервера в отдельной горутине
	go func() {
		if err := a.Run(); err != nil {
			log.Fatal().Msgf("Server failed:  %s", err.Error())
		}
		close(shutdownChan)
	}()

	go func() {
		sig := <-sigChan
		log.Info().Msgf("Received signal: %s", sig)

		// Контекст с таймаутом для завершения активных запросов
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Завершение сервера
		if err := a.ShutdownServer(ctx); err != nil {
			log.Info().Msgf("Server forced to shutdown: %v", err)
		}

		close(shutdownChan)
	}()

	<-shutdownChan
	log.Info().Msg("Server gracefully stopped")
	os.Exit(0)
}
