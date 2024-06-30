package rest

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/storage/postgres"
)

// PingHandler хендлер ping метода.
func (h HTTPHandler) PingHandler(res http.ResponseWriter) {
	switch v := h.storage.(type) {
	case *postgres.PostgresStorage:
		err := v.DB.Ping()

		if err != nil {
			log.Info().Msg(fmt.Sprintf("postgres ping error: %s", err.Error()))
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		log.Info().Msg("storage is not PgSQL")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
