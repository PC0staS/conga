#!/bin/bash

# Build script for CONGA - compile for multiple platforms

echo "🔨 Building CONGA for multiple platforms...\n"

mkdir build
cd build

# Linux 64-bit
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o conga-linux ..

# Linux ARM64
echo "Building for Linux ARM64..."
GOOS=linux GOARCH=arm64 go build -o conga-linux-arm64 ..

# macOS Intel
echo "Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o conga-macos-intel ..

# macOS Apple Silicon
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o conga-macos-arm64 ..

# Windows 64-bit
echo "Building for Windows (64-bit)..."
GOOS=windows GOARCH=amd64 go build -o conga-windows.exe ..

# Windows 32-bit
echo "Building for Windows (32-bit)..."
GOOS=windows GOARCH=386 go build -o conga-windows-32.exe ..

cd ..

echo "\n✅ Build complete! Binaries in ./build/"
echo "\nFiles generated:"
ls -lh build/