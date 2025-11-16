package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// GzipMiddleware compresses HTTP responses using gzip if the client supports it
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()

			gzWriter := gzipResponseWriter{ResponseWriter: w, Writer: gz}
			next.ServeHTTP(gzWriter, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// gzipResponseWriter wraps http.ResponseWriter to write gzip-compressed data
type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
