package main

import (
	"net/http"
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

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.checkServerHandler)
		r.Get("/index", app.indexPageHandler)
		r.Get("/login", app.loginHandler)
		r.Get("/home", app.homeHandler)
	})
	
	return r
}

// HTTPS server only (TLS)
func (app *application) runTLS(mux http.Handler, certFile, keyFile string) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServeTLS(certFile, keyFile)
}

// Placeholder handlers
func (app *application) indexPageHandler(w http.ResponseWriter, r *http.Request) {}
func (app *application) loginHandler(w http.ResponseWriter, r *http.Request)     {}
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request)      {}
