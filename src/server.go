package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

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

func handle404(handler http.Handler, config Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		interceptor := &notFoundInterceptor{
			ResponseWriter: w,
			config:         config,
		}
		handler.ServeHTTP(interceptor, r)
	})
}

func startServer(config Config) {
	r := mux.NewRouter()

	ApplyMiddleware(r, config)

	fs := http.FileServer(http.Dir(config.StaticDir))

	// Wrap file server with 404 interceptor if error pages are enabled
	if config.ErrorPages.Enabled {
		r.PathPrefix("/").Handler(http.StripPrefix("/", handle404(fs, config)))
	} else {
		r.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	}

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
