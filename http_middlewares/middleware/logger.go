package middleware

import (
	"log"
	"net/http"
	"time"
)

func NewLoggerMiddleware(nanosecond bool) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("Serving request")
			next.ServeHTTP(w, r)
			end := time.Now()
			if nanosecond {
				log.Printf("Done serving. Elapsed time is %dns\n", end.Sub(start).Nanoseconds())
			} else {
				log.Printf("Done serving. Elapsed time is %dms\n", end.Sub(start).Milliseconds())
			}
		})
	}
}
