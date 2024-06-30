package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/utils/authorize"
)

// DeleteURLsHandlerRequest структура запроса.
type DeleteURLsHandlerRequest []string

// DeleteURLsHandler json хендлер удаления массива URLs.
func (h HTTPHandler) DeleteURLsHandler(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res.WriteHeader(http.StatusAccepted)
	userID, err := authorize.GetUserID(req)
	if err != nil || userID == "" {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	var request DeleteURLsHandlerRequest
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &request); err != nil {
		log.Info().Msg(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	h.shortener.DeleteURLs(ctx, request, userID)
}
