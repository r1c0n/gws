package middleware

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	sessionLogFile  *os.File
	realtimeLogFile *os.File
	logFilePath     string
	realtimeLogPath = "logs/realtimelogs.log"
)

// InitLogFiles creates and initializes session and realtime log files
func InitLogFiles() error {
	// create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return err
	}

	// create the session log file with a timestamped name
	logFileName := time.Now().Format("2006-01-02-15-04-05") + ".log"
	logFilePath = filepath.Join(logDir, logFileName)
	var err error
	sessionLogFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// clear the realtimelogs.log file
	realtimeLogFile, err = os.OpenFile(realtimeLogPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	// set log output to the session log file and also to realtime log file
	log.SetOutput(sessionLogFile)
	multiWriter := io.MultiWriter(sessionLogFile, realtimeLogFile)
	log.SetOutput(multiWriter)

	return nil
}

func CloseLogFiles() {
	if sessionLogFile != nil {
		sessionLogFile.Close()
	}
	if realtimeLogFile != nil {
		realtimeLogFile.Close()
	}
}

type responseLogger struct {
	http.ResponseWriter
	status int
}

func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status
	rl.ResponseWriter.WriteHeader(status)
}

// LoggingMiddleware logs all HTTP requests with method, path, and timestamp
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
