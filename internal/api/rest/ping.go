package rest

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/storage/postgres"
)

func (h HTTPHandler) PingHandler(res http.ResponseWriter) {
	if config.Config.DatabaseDSN == "" {
		log.Info().Msg("database dsn is not init")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	db, err := postgres.NewPostgresStorage(config.Config.DatabaseDSN)
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
