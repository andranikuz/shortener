package app

import (
	"context"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/container"
)

// Application структура используется для запуска приложения. В ней инициализуется контейнер.
type Application struct {
	cnt    *container.Container
	server *http.Server
}

// NewApplication создает новое приложение.
func NewApplication() (*Application, error) {
	if err := config.Init(); err != nil {
		return nil, err
	}
	cnt, err := container.NewContainer()
	if err != nil {
		return nil, err
	}
	a := Application{
		cnt: cnt,
	}

	return &a, nil
}

// Run запускет http сервер.
func (a *Application) Run() error {
	httpHandler := rest.NewHTTPHandler(a.cnt)
	a.server = &http.Server{
		Addr:    config.Config.ServerAddress,
		Handler: httpHandler.Router(),
	}

	if config.Config.EnableHTTPS {
		pwd, _ := os.Getwd()
		path := pwd + `/internal/config/crt/`

		return a.server.ListenAndServeTLS(path+"server.crt", path+"server.key")
	} else {
		return a.server.ListenAndServe()
	}
}

// ShutdownServer завершает работу HTTP сервера
func (a *Application) ShutdownServer(ctx context.Context) error {
	log.Info().Msg("Shutting down server...")
	return a.server.Shutdown(ctx)
}
