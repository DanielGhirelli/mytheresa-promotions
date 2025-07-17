package http

import (
	"log"
	"net/http"
	"time"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Continue handling request before logging
		next.ServeHTTP(w, r)

		// After response is handled log
		duration := time.Since(start)
		log.Printf("[%s] %s %s?%s - IP: %s - Duration: %.2fms",
			r.Method,
			r.Host,
			r.URL.Path,
			r.URL.RawQuery,
			r.RemoteAddr,
			float64(duration.Microseconds())/1000,
		)
	})
}

func APIKeyMiddleware(validApiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("x-api-key")

			// Validate API key
			if apiKey != validApiKey {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error": "Invalid or missing API key"}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
