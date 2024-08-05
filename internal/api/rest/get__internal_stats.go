package rest

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/models"
)

// GetInternalStatsResponse структура ответа.
type GetInternalStatsResponse struct {
	URLs  int `json:"urls"`
	Users int `json:"users"`
}

// GetInternalStats хендлер получения статистики сервиса.
func (h HTTPHandler) GetInternalStats(res http.ResponseWriter, req *http.Request) {
	if !checkIP(req.Header.Get("X-Real-IP")) {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	urls, users, err := h.shortener.GetInternalStats(req.Context())
	if err != nil {
		if errors.Is(err, models.ErrURLDeleted) {
			log.Info().Msg(err.Error())
			res.WriteHeader(http.StatusGone)
			return
		} else {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	response := GetInternalStatsResponse{
		URLs:  urls,
		Users: users,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := res.Write(resp); err != nil {
		log.Info().Msg(err.Error())
		res.WriteHeader(http.StatusBadRequest)
	}
}

func checkIP(ip string) bool {
	if config.Config.TrustedSubnet == "" {
		return false
	}

	_, subnet, err := net.ParseCIDR(config.Config.TrustedSubnet)
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}

	clientIP := net.ParseIP(ip)
	if clientIP == nil || !subnet.Contains(clientIP) {
		return false
	}

	return true
}
