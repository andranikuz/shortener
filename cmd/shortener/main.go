package main

import (
	"github.com/andranikuz/shortener/internal/app"
	"net/http"
)

func main() {
	a := app.Application{}
	a.Init()
	http.ListenAndServe("localhost:8080", a.Router())
}
