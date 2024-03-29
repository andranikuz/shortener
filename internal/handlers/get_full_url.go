package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/andranikuz/shortener/internal/app/usecases"
)

func GetFullURLHandler(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
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
