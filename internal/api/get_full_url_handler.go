package api

import (
	"github.com/andranikuz/shortener/internal/app/usecases"
	"net/http"
	"strings"
)

func GetFullUrlHandler(res http.ResponseWriter, req *http.Request) {
	id := strings.ReplaceAll(req.RequestURI, "/", "")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fullUrl := usecases.GetFullUrl(id)
	if fullUrl == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", fullUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
