package main

import (
	"log"

	"github.com/0x03ff/web_go/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load("internal/env/.env"); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Configure addresses from environment variables
	httpsAddr := env.GetString("HTTPS_ADDR", "192.168.0.10:8086")

	// Create application config
	cfg := config{
		addr: httpsAddr,
	}

	app := &application{config: cfg}
	mux := app.mount()

	// Start HTTPS server (this will block the main thread)
	log.Printf(" HTTPS server starting on %s", httpsAddr)
	if err := app.run(mux); err != nil {
		log.Fatalf(" HTTPS server failed: %v", err)
	}
}
