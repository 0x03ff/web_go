#!/usr/bin/env bash
set -e

# Always resolve and run from the workspace root directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$SCRIPT_DIR"

# Ensure the local testing challenge directory exists
mkdir -p /tmp/acme-challenges

echo "========================================================"
echo "    Starting local domain-validation testing server     "
echo "========================================================"
echo "  Serving challenge files from: /tmp/acme-challenges"
echo ""
echo "  Endpoints:"
echo "    - http://127.0.0.1:8080/.well-known/acme-challenge/<filename>"
echo "    - http://127.0.0.1:8080/.well-known/pki-validation/<filename>"
echo ""
echo "  Press Ctrl+C to stop."
echo "--------------------------------------------------------"

# Boot the server using Go's source run pattern
go run ./cmd