package main

import (
	"github.com/andranikuz/shortener/internal/app"
	"log"
	"net/http"
)

func main() {
	a := app.Application{}
	a.Init()
	http.HandleFunc("/", a.Handle)
	log.Println("Starting server on :8080...")
	http.ListenAndServe("localhost:8080", nil)
}
