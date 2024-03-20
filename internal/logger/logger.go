package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/andranikuz/shortener/internal/config"
)

func RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		defer func() {
			requestLog := fmt.Sprintf("uri=%s method=%s in %s", r.RequestURI, r.Method, time.Since(start))
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

func Init() error {
	if config.Config.LogLevel != "" {
		level, err := zerolog.ParseLevel(config.Config.LogLevel)
		if err != nil {
			return fmt.Errorf("wrong LOG_LEVEL param in ENV: %s", config.Config.LogLevel)
		}

		zerolog.SetGlobalLevel(level)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return nil
}
