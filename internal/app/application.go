package app

import (
	"context"
	"net/http"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/container"
)

// Application структура используется для запуска приложения. В ней инициализуется контейнер.
type Application struct {
	cnt *container.Container
	ctx context.Context
}

// NewApplication создает новое приложение.
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

// Run запускет http сервер.
func (a *Application) Run() error {
	httpHandler := rest.NewHTTPHandler(a.cnt)
	return http.ListenAndServe(config.Config.ServerAddress, httpHandler.Router(a.ctx))
}
