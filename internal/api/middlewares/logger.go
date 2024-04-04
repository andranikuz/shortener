package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
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
