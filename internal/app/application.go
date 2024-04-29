package app

import (
	"context"
	"net/http"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/container"
)

type Application struct {
	cnt *container.Container
	ctx context.Context
}

func NewApplication() (*Application, error) {
	config.Init()
	cnt, err := container.NewContainer()
	if err != nil {
		return nil, err
	}
	a := Application{
		ctx: context.Background(),
		cnt: cnt,
	}

	return &a, nil
}

func (a *Application) Run() error {
	httpHandler := rest.NewHTTPHandler(a.cnt)
	return http.ListenAndServe(config.Config.ServerAddress, httpHandler.Router(a.ctx))
}
