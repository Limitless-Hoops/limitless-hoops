#!/bin/bash

set -e

GO_VERSION="1.24.2"
DOCKER_VERSION="28.1.1"

ARCH="amd64"
OS="$(uname -s)"

install_go() {
  echo "🔧 Installing Go $GO_VERSION..."

  if go version 2>/dev/null | grep "$GO_VERSION" >/dev/null; then
    echo "✅ Go $GO_VERSION is already installed"
    return
  fi

  echo "📦 Downloading Go $GO_VERSION..."
  curl -LO "https://go.dev/dl/go${GO_VERSION}.${1}-${ARCH}.tar.gz"

  sudo rm -rf /usr/local/go
  sudo tar -C /usr/local -xzf "go${GO_VERSION}.${1}-${ARCH}.tar.gz"
  rm "go${GO_VERSION}.${1}-${ARCH}.tar.gz"

  if ! grep -q "/usr/local/go/bin" ~/.bashrc 2>/dev/null; then
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
  fi
  export PATH=$PATH:/usr/local/go/bin

  echo "✅ Go $GO_VERSION installed"
}

install_docker_linux() {
  echo "🐳 Installing Docker Engine $DOCKER_VERSION (Linux)..."

  if docker version 2>/dev/null | grep "$DOCKER_VERSION" >/dev/null; then
    echo "✅ Docker $DOCKER_VERSION already installed"
    return
  fi

  sudo apt-get update
  sudo apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

  sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL "https://download.docker.com/linux/$(. /etc/os-release && echo "$ID")/gpg"
    sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
  sudo chmod a+r /etc/apt/keyrings/docker.gpg

  echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/$(. /etc/os-release && echo "$ID") \
    $(lsb_release -cs) stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

  sudo apt-get update
  sudo apt-get install -y \
    docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

  echo "✅ Docker installed"
}

install_docker_macos() {
  echo "🐳 Installing Docker Desktop (macOS)..."

  if command -v docker > /dev/null; then
    echo "✅ Docker already installed"
    return
  fi

  echo "🚨 Docker Desktop must be installed manually on macOS:"
  echo "👉 Visit: https://www.docker.com/products/docker-desktop/"
}

# === OS Handling ===

if [[ "$OS" == "Darwin" ]]; then
  echo "🖥 Detected macOS"
  install_go "darwin"
  install_docker_macos
elif [[ "$OS" == "Linux" ]]; then
  echo "🐧 Detected Linux"
  install_go "linux"
  install_docker_linux
else
  echo "❌ Unsupported OS: $OS"
  exit 1
fi
