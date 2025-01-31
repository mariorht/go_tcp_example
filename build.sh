#!/bin/bash

# Directorio de builds
BUILD_DIR="builds"
LINUX_DIR="$BUILD_DIR/linux"
WINDOWS_DIR="$BUILD_DIR/windows"
MAC_DIR="$BUILD_DIR/mac"

# Crear las carpetas de builds
mkdir -p "$LINUX_DIR" "$WINDOWS_DIR" "$MAC_DIR"

# Compilar para Linux
echo "🔧 Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/server" ./server/main.go
GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/client" ./client/main.go

# Compilar para Windows
echo "🔧 Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/server.exe" ./server/main.go
GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/client.exe" ./client/main.go

# Compilar para macOS
echo "🔧 Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o "$MAC_DIR/server" ./server/main.go
GOOS=darwin GOARCH=amd64 go build -o "$MAC_DIR/client" ./client/main.go

echo "✅ Build completed! Check the 'builds/' directory."
