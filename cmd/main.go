package main

import (
	_ "embed"
	"log"
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

	// The precise validation filename provided by Certificate Authority (CA) like ZeroSSL
	TEXT_NAME = "37B60801A77A091CF0AED6D1ECA6B65C.txt"
	//    TEXT_NAME = ".txt"

)

/*
	please change with the file name down below
	such that the binary contain the text file
*/

//go:embed 37B60801A77A091CF0AED6D1ECA6B65C.txt
var validationFileContents string

type config struct {
	addr     string
	textName string
}

type application struct {
	config config
}

func main() {
	if TEXT_NAME == "" || HTTP_ADDR == "" {
		log.Fatal("Configuration Error: Configuration constants cannot be empty.")
	}

	cfg := config{
		addr:     HTTP_ADDR,
		textName: TEXT_NAME,
	}

	app := &application{config: cfg}
	mux := app.mount()

	log.Printf(" HTTP server starting on %s", cfg.addr)
	if err := app.run(mux); err != nil {
		log.Fatalf(" HTTP server failed: %v", err)
	}
}
