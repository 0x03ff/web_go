#!/usr/bin/env bash
set -e

CHALLENGE_DIR="/tmp/acme-challenges"
mkdir -p "$CHALLENGE_DIR"

echo "=== Add ACME / PKI Challenge File ==="
echo ""

# 1. Ask for the file name
read -p "Enter challenge filename (e.g., 37B60801A77A091CF0AED6D1ECA6B65C.txt): " FILENAME

if [ -z "$FILENAME" ]; then
    echo "Error: Filename cannot be empty."
    exit 1
fi

FILEPATH="$CHALLENGE_DIR/$FILENAME"

# 2. Open nano editor directly for clean pasting
EDITOR_CMD="${EDITOR:-nano}"

echo ""
echo "Opening editor to paste your payload..."
echo "Press Ctrl+O -> Enter to save, then Ctrl+X to exit."
sleep 1

$EDITOR_CMD "$FILEPATH"

echo ""
echo "✔ Successfully created: $FILEPATH"
echo "  Test URL: http://127.0.0.1:8080/.well-known/pki-validation/$FILENAME"