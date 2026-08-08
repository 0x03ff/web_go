package main

import (
	_ "embed"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// ==============================================================================
// CONFIGURATION (Please modify these values directly)
// ==============================================================================
const (
	// Choose the running address here:
	// "127.0.0.1:8080"                  -> Local testing
	// "0.0.0.0:8080"                    -> Public on port 8080
	// "0.0.0.0:80"                      -> Public on standard HTTP port (Requires sudo)
	HTTP_ADDR = "0.0.0.0:8080"

	// The precise validation filename provided by Certificate Authority (CA) like ZeroSSL / Let's Encrypt
	TEXT_NAME = "challenge.txt"
)

/*
	Please place your text file inside the same directory as main.go
	and ensure the filename matches the directive below.
*/

//go:embed challenge.txt
var validationFileContents string

func main() {
	if TEXT_NAME == "" || HTTP_ADDR == "" {
		log.Fatal("Configuration Error: Configuration constants cannot be empty.")
	}

	mux := http.NewServeMux()

	// Handler serves embedded string data for any valid verification request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check if requested URL matches expected validation filename
		filename := filepath.Base(r.URL.Path)
		if filename != TEXT_NAME {
			http.Error(w, "404 Challenge File Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(strings.TrimSpace(validationFileContents)))
	})

	// Match both ACME (Let's Encrypt) and PKI (ZeroSSL / Sectigo) validation routes
	mux.Handle("/.well-known/acme-challenge/", loggerMiddleware(handler))
	mux.Handle("/.well-known/pki-validation/", loggerMiddleware(handler))

	log.Printf("[INFO] web_go running on %s. Serving embedded challenge for %s", HTTP_ADDR, TEXT_NAME)

	srv := &http.Server{
		Addr:         HTTP_ADDR,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}

// Lightweight logging middleware to track incoming verification hits
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[HTTP] %s %s from %s took %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}