#!/bin/bash
set -e

CERT_DIR="./nginx/certs"
DOMAIN="limitlesshoops.dev"

echo "🔐 Setting up mkcert certificates for $DOMAIN..."

# Install mkcert if not available
if ! command -v mkcert &> /dev/null; then
  echo "📦 Installing mkcert..."
  if [[ "$(uname)" == "Darwin" ]]; then
    brew install mkcert
  elif [[ "$(uname)" == "Linux" ]]; then
    sudo apt install -y libnss3-tools
    curl -L https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64 -o mkcert
    chmod +x mkcert && sudo mv mkcert /usr/local/bin/
  else
    echo "❌ Unsupported OS"
    exit 1
  fi
fi

# Initialize the local CA
mkcert -install

# Create the certs directory
mkdir -p "$CERT_DIR"

# Generate cert and key
mkcert -key-file "$CERT_DIR/$DOMAIN.key" \
       -cert-file "$CERT_DIR/$DOMAIN.crt" \
       "$DOMAIN"

echo "✅ Certificates created:"
echo "  $CERT_DIR/$DOMAIN.crt"
echo "  $CERT_DIR/$DOMAIN.key"
