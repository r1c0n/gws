package main

import (
	"crypto/tls"
	"log"
)

func getCertificates(certFile, keyFile string) []tls.Certificate {
	if certFile == "" || keyFile == "" {
		log.Fatal("Certificate file or key file not specified")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	return []tls.Certificate{cert}
}
