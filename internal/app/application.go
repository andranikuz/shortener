package app

import (
	"net/http"
	"os"

	"github.com/andranikuz/shortener/internal/api/rest"
	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/container"
)

// Application структура используется для запуска приложения. В ней инициализуется контейнер.
type Application struct {
	cnt *container.Container
}

// NewApplication создает новое приложение.
func NewApplication() (*Application, error) {
	config.Init()
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
	if config.Config.EnableHTTPS {
		pwd, _ := os.Getwd()
		path := pwd + `/internal/config/crt/`

		return http.ListenAndServeTLS(config.Config.ServerAddress, path+"server.crt", path+"server.key", httpHandler.Router())
	} else {
		return http.ListenAndServe(config.Config.ServerAddress, httpHandler.Router())
	}
}
