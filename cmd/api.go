package main

import (
    "log"
    "net/http"
    "strings"
    "time"
)

func (app *application) mount() http.Handler {
    // 1. Create a native Go ServeMux (Router)
    mux := http.NewServeMux()

    // 2. Build the explicit URL match path
    // e.g., "/.well-known/pki-validation/37B60801A77A091CF0AED6D1ECA6B65C.txt"
    routePath := "/.well-known/pki-validation/" + app.config.textName

    // 3. Register the handler specifically for this path
    mux.HandleFunc(routePath, func(w http.ResponseWriter, req *http.Request) {
        // Guard clause: Ensure it only handles exact matches and GET requests
        if req.Method != http.MethodGet {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        contentReader := strings.NewReader(validationFileContents)
        http.ServeContent(w, req, app.config.textName, time.Now(), contentReader)
    })

    // 4. Wrap the mux with custom lightweight logging middleware
    return app.loggerMiddleware(mux)
}

// A simple native logging middleware to track incoming verification requests
func (app *application) loggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Process the actual request
        next.ServeHTTP(w, r)
        
        // Log the details once complete
        log.Printf("[HTTP] %s %s from %s took %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
    })
}

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