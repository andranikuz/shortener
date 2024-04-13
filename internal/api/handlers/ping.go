package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/storage/postgres"
)

func PingHandler(res http.ResponseWriter, a app.Application) {
	if config.Config.DatabaseDSN == "" {
		log.Info().Msg(fmt.Sprintf("database dsn is not init"))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	db, err := postgres.NewPostgresDB(config.Config.DatabaseDSN)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("postgres connect error: %s", err.Error()))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.DB.Close()

	err = db.DB.Ping()
	if err != nil {
		log.Info().Msg(fmt.Sprintf("postgres ping error: %s", err.Error()))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
