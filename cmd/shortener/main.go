package main

import (
	"github.com/andranikuz/shortener/internal/app"
)

func main() {
	a := app.Application{}
	if err := a.Run(); err != nil {
		panic(err)
	}
}
