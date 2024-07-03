package rest

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// PingHandler хендлер ping метода.
func (h HTTPHandler) PingHandler(res http.ResponseWriter) {
	err := h.storage.Ping()
	if err != nil {
		log.Info().Msg(fmt.Sprintf("postgres ping error: %s", err.Error()))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
