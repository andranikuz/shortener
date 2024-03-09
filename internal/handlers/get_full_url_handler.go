package handlers

import (
	"github.com/andranikuz/shortener/internal/app/usecases"
	"net/http"
	"strings"
)

func GetFullURLHandler(res http.ResponseWriter, req *http.Request) {
	id := strings.ReplaceAll(req.RequestURI, "/", "")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fullURL := usecases.GetFullURL(id)
	if fullURL == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", fullURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
