package app

import (
	"github.com/andranikuz/shortener/internal/api"
	"github.com/andranikuz/shortener/internal/storage"
	"net/http"
)

type Application struct {
}

func (app *Application) Init() {
	storage.Init()
}

func (app *Application) Handle(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		api.GenerateShortUrlHandler(res, req)
	} else if req.Method == http.MethodGet {
		api.GetFullUrlHandler(res, req)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
