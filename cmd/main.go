package main

import (
	"log"
	"net/http"
	"os"
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
	HTTP_ADDR = "127.0.0.1:8080"

	// please enter the cert location depend the file structure
	DATA_DIR  = "/tmp/acme-challenges"


)
func main() {
	// Ensure challenge directory exists
	if err := os.MkdirAll(DATA_DIR, 0755); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}

	mux := http.NewServeMux()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract filename from URL (e.g., /.well-known/pki-validation/<filename>)
		filename := filepath.Base(r.URL.Path)
		filePath := filepath.Join(DATA_DIR, filename)

		// Check if file exists
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "404 Challenge File Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(strings.TrimSpace(string(data))))
	})

	// Match both ACME and PKI validation routes
	mux.Handle("/.well-known/acme-challenge/", handler)
	mux.Handle("/.well-known/pki-validation/", handler)

	log.Printf("[INFO] web_go running on %s. Serving challenges from %s", HTTP_ADDR, DATA_DIR)
	srv := &http.Server{Addr: HTTP_ADDR, Handler: mux, ReadTimeout: 5 * time.Second}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}