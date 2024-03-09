package app

import (
	"github.com/andranikuz/shortener/internal/storage"
)

type Application struct {
}

func (app *Application) Init() {
	storage.Init()
}
