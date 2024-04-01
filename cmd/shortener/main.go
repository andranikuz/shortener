package main

import (
	"github.com/andranikuz/shortener/internal/api"
	"github.com/andranikuz/shortener/internal/app"
)

func main() {
	a := app.Application{}
	if err := a.Run(api.Router()); err != nil {
		panic(err)
	}
}
