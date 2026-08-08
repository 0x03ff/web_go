#!/bin/bash

# Exit immediately if go run encounters an error
set -e

echo "Starting local domain-validation testing server..."
echo "Target endpoint: http://127.0.0.1:8080/.well-known/pki-validation/ or http://127.0.0.1:8080/.well-known/acme-challenge/"
echo "Press Ctrl+C to stop."
echo "--------------------------------------------------------"

# Boot the server using Go's source run pattern
go run ./cmd