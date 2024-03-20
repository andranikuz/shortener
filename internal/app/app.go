package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/config"
	"github.com/andranikuz/shortener/internal/handlers"
	"github.com/andranikuz/shortener/internal/storage"
)

type Application struct {
}

func (app *Application) Run() {
	app.Init()
	http.ListenAndServe(config.Config.ServerAddress, app.Router())
}

func (app *Application) Init() {
	storage.Init()
	config.Init()
}

func (app *Application) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(app.Logger)
	r.Post("/", handlers.GenerateShortURLHandler)
	r.Get("/{id}", handlers.GetFullURLHandler)
	r.Post("/{url}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return r
}

func (app *Application) Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		t1 := time.Now()
		defer func() {
			requestLog := fmt.Sprintf("uri=%s method=%s in %s", r.RequestURI, r.Method, time.Since(t1))
			responseLog := fmt.Sprintf("code=%v size=%v", ww.Status(), ww.BytesWritten())
			log.Info().
				Str("request", requestLog).
				Str("response", responseLog).
				Msg("")
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
