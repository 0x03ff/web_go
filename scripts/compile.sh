#!/usr/bin/env bash

# Exit immediately on error, undefined variables, or pipe failures
set -euo pipefail

# Always run relative to project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$SCRIPT_DIR"

echo "=========================================="
echo "      Go Cross-Compiler CLI Wizard        "
echo "=========================================="

# 1. Select Operating System
echo "Select Target OS:"
echo "1) Linux"
echo "2) macOS (Darwin)"
echo "3) Windows"
read -p "Enter choice [1-3]: " os_choice

case $os_choice in
    1) GOOS="linux" ;;
    2) GOOS="darwin" ;;
    3) GOOS="windows" ;;
    *) echo "Invalid OS selection. Exiting."; exit 1 ;;
esac

# 2. Select Architecture
echo -e "\nSelect Target Architecture:"
echo "1) amd64 (64-bit Intel/AMD)"
echo "2) arm64 (Apple Silicon, TV Boxes, Raspberry Pi)"
read -p "Enter choice [1-2]: " arch_choice

case $arch_choice in
    1) GOARCH="amd64" ;;
    2) GOARCH="arm64" ;;
    *) echo "Invalid Architecture selection. Exiting."; exit 1 ;;
esac

# 3. Handle File Extension for Windows Binaries
EXT=""
if [ "$GOOS" = "windows" ]; then
    EXT=".exe"
fi

# 4. Ensure output directory exists
mkdir -p bin

# 5. Set Output Name dynamically based on choices
OUTPUT_NAME="bin/web_go-${GOOS}-${GOARCH}${EXT}"

echo -e "\n------------------------------------------"
echo "Compiling for: GOOS=$GOOS | GOARCH=$GOARCH"
echo "Output file:   $OUTPUT_NAME"
echo "------------------------------------------"

# Execute the Go build command with stripped debugging symbols for minimal binary size
CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$OUTPUT_NAME" ./cmd

echo "✔ Build successful!"
ls -lh "$OUTPUT_NAME"