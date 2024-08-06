package app

import (
	"net/http"
	"os"

	grpcserver "github.com/andranikuz/shortener/internal/api/grpc"
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

	go func() {
		if err := grpcserver.StartGRPCServer(`:3200`, a.cnt); err != nil {
			panic(err)
		}
	}()

	if config.Config.EnableHTTPS {
		pwd, _ := os.Getwd()
		path := pwd + `/internal/config/crt/`

		return a.server.ListenAndServeTLS(path+"server.crt", path+"server.key")
	} else {
		return a.server.ListenAndServe()
	}
}
