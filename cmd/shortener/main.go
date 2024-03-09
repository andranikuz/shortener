package main

import (
	"github.com/andranikuz/shortener/internal/app"
	"github.com/andranikuz/shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	a := app.Application{}
	a.Init()
	r := chi.NewRouter()
	r.Post("/", handlers.GenerateShortURLHandler)
	r.Get("/{id}", handlers.GetFullURLHandler)
	http.ListenAndServe("localhost:8080", r)
}
