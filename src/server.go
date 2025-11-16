package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/r1c0n/gws/middleware"
)

// notFoundInterceptor wraps ResponseWriter to intercept 404 status codes
// and serve custom error pages before http.FileServer writes its default response
type notFoundInterceptor struct {
	http.ResponseWriter
	config        Config
	headerWritten bool
	customServed  bool
}

func (nfi *notFoundInterceptor) WriteHeader(code int) {
	if nfi.headerWritten {
		return
	}
	nfi.headerWritten = true

	if code == 404 {
		// Serve custom 404 page
		if pageName, exists := nfi.config.ErrorPages.Pages["404"]; exists {
			errorPagePath := filepath.Join(nfi.config.ErrorPages.ErrorPagesDir, pageName)
			if content, err := os.ReadFile(errorPagePath); err == nil {
				nfi.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
				nfi.ResponseWriter.WriteHeader(http.StatusNotFound)
				nfi.ResponseWriter.Write(content)
				nfi.customServed = true
				return
			}
		}
	}
	nfi.ResponseWriter.WriteHeader(code)
}

func (nfi *notFoundInterceptor) Write(b []byte) (int, error) {
	// If we already served a custom page, ignore any further writes
	if nfi.customServed {
		return len(b), nil
	}

	if !nfi.headerWritten {
		nfi.WriteHeader(http.StatusOK)
	}
	return nfi.ResponseWriter.Write(b)
}

// handle404 wraps a handler to intercept and customize 404 responses
func handle404(handler http.Handler, config Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interceptor := &notFoundInterceptor{
			ResponseWriter: w,
			config:         config,
		}
		handler.ServeHTTP(interceptor, r)
	})
}

// startServer initializes and starts the HTTP/HTTPS server with configured middleware
func startServer(config Config) {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(config.StaticDir))

	// Start with the file server
	var handler http.Handler = http.StripPrefix("/", fs)

	// Apply middleware in reverse order (innermost to outermost)
	// Order: Logging → Rate Limit → CORS → Gzip → 404 Handler

	// Apply logging middleware (innermost - executes last)
	if config.Middleware.LoggingMiddlewareEnabled {
		if err := middleware.InitLogFiles(); err != nil {
			log.Fatalf("Could not initialize log files: %v", err)
		}
		handler = middleware.LoggingMiddleware(handler)
	}

	// Apply rate limiting middleware
	if config.RateLimit.Enabled {
		errorPagePath := ""
		if config.ErrorPages.Enabled {
			if pageName, exists := config.ErrorPages.Pages["429"]; exists {
				errorPagePath = filepath.Join(config.ErrorPages.ErrorPagesDir, pageName)
			}
		}
		rateLimitConfig := middleware.RateLimitConfig{
			Enabled:           config.RateLimit.Enabled,
			RequestsPerMinute: config.RateLimit.RequestsPerMinute,
			Burst:             config.RateLimit.Burst,
			Whitelist:         config.RateLimit.Whitelist,
			ExemptPaths:       config.RateLimit.ExemptPaths,
			ErrorPagePath:     errorPagePath,
		}
		handler = middleware.RateLimitMiddleware(rateLimitConfig)(handler)
	}

	// Apply CORS middleware
	if config.CORS.Enabled {
		corsConfig := middleware.CORSConfig{
			Enabled:          config.CORS.Enabled,
			AllowedOrigins:   config.CORS.AllowedOrigins,
			AllowedMethods:   config.CORS.AllowedMethods,
			AllowedHeaders:   config.CORS.AllowedHeaders,
			AllowCredentials: config.CORS.AllowCredentials,
			MaxAge:           config.CORS.MaxAge,
		}
		handler = middleware.CORSMiddleware(corsConfig)(handler)
	}

	// Apply gzip middleware
	if config.Middleware.GzipMiddlewareEnabled {
		handler = middleware.GzipMiddleware(handler)
	}

	// Wrap with 404 interceptor
	if config.ErrorPages.Enabled {
		handler = handle404(handler, config)
	}

	r.PathPrefix("/").Handler(handler)

	if config.TLSConfig.Enabled {
		server := &http.Server{
			Addr:    config.Port,
			Handler: r,
			TLSConfig: &tls.Config{
				Certificates: getCertificates(config.TLSConfig.CertFile, config.TLSConfig.KeyFile),
			},
		}

		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		log.Fatal(http.ListenAndServe(config.Port, r))
	}
}
