package api

import (
	"github.com/andranikuz/shortener/internal/app/usecases"
	"io"
	"net/http"
)

func GenerateShortUrlHandler(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI != "/" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ := io.ReadAll(req.Body)

	fullUrl := string(body)
	if fullUrl == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl := usecases.GenerateShortUrl(fullUrl)
	res.WriteHeader(http.StatusCreated)
	if _, err := res.Write([]byte(shortUrl)); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/plain")

}
