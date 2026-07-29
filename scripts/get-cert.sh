#!/usr/bin/env bash
set -e

# 1. Ensure a domain name was provided
DOMAIN="${1:-}"

if [ -z "$DOMAIN" ]; then
    echo "Usage: sudo $0 <domain-name>"
    echo "Example: sudo $0 example.com"
    exit 1
fi

CHALLENGE_DIR="/tmp/acme-challenges"
mkdir -p "$CHALLENGE_DIR"

echo "========================================================"
echo " Requesting SSL Certificate for: $DOMAIN"
echo "========================================================"

# 2. Run Certbot with inline bash hooks
sudo certbot certonly --manual \
  --preferred-challenges http \
  --manual-auth-hook "mkdir -p $CHALLENGE_DIR && echo \$CERTBOT_VALIDATION > $CHALLENGE_DIR/\$CERTBOT_TOKEN" \
  --manual-cleanup-hook "rm -f $CHALLENGE_DIR/\$CERTBOT_TOKEN" \
  -d "$DOMAIN"

echo ""
echo "✔ Certificate issuance complete for $DOMAIN!"