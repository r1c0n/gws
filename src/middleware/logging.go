package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseLogger struct {
	http.ResponseWriter
	status int
}

func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status
	rl.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// record the current time
		start := time.Now()

		// create a response logger
		rl := &responseLogger{w, http.StatusOK}

		// perform operations before handling the request
		log.Printf("[%s] Received request: %s %s from %s\n", start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)

		// log headers
		log.Println("Headers:")
		for key, value := range r.Header {
			log.Printf("%s: %s", key, value)
		}

		// log query parameters
		log.Println("Query Parameters:")
		r.ParseForm()
		for key, value := range r.Form {
			log.Printf("%s: %s", key, value)
		}

		// call the next handler
		next.ServeHTTP(rl, r)

		// log the response status code and duration
		log.Printf("[%s] Responded with status: %d in %s", time.Now().Format("2006-01-02 15:04:05"), rl.status, time.Since(start))
	})
}
