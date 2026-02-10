package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// For dependencies
type application struct {
	config config
}

// Configuration
type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/.well-known/pki-validation/*", func(w http.ResponseWriter, req *http.Request) {
			// Use os.DirFS to create a file system rooted at the current directory
			// and wrap it with http.FS.
			// In a real application, you might use go:embed for a more robust solution.
			const fileName = "7F222F8032FFA57AEC8DB93BEAEC3895.txt"

			// Open the file
			f, err := os.Open(fileName)
			if err != nil {
				http.Error(w, "File not found", http.StatusNotFound)
				return
			}
			defer f.Close()

			// Serve the file directly
			http.ServeContent(w, req, fileName, time.Now(), f)
		})

	})

	return r
}

// HTTPS server only (TLS)
func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
