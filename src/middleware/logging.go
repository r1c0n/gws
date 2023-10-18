package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// perform operations before handling the request
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		// call the next handler
		next.ServeHTTP(w, r)
	})
}
