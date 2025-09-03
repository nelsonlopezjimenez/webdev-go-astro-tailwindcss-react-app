Here are the commands to cross-compile your Go application for macOS:

## From Windows to macOS

**For Intel Macs (x86_64):**
```bash
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o course-server-macos-amd64 embedded_server.go
```

**For Apple Silicon Macs (ARM64/M1/M2):**
```bash
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -o course-server-macos-arm64 embedded_server.go
```

**Universal binary (both architectures):**
```bash
# Build both versions first
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o course-server-amd64 embedded_server.go

set GOARCH=arm64
go build -ldflags="-s -w" -o course-server-arm64 embedded_server.go

# Combine into universal binary (requires macOS to run lipo)
lipo -create -output course-server-macos course-server-amd64 course-server-arm64
```

## From Linux/macOS to macOS

**For Intel Macs:**
```bash
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o course-server-macos-amd64 embedded_server.go
```

**For Apple Silicon Macs:**
```bash
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o course-server-macos-arm64 embedded_server.go
```

## Build Script for Multiple Platforms

Create `build-multi.sh`:
```bash
#!/bin/bash
set -e

# Build Astro frontend first
cd lessons/frontend
npm run build
cd ../..

# Build for multiple platforms
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o course-server-windows-amd64.exe embedded_server.go

echo "Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o course-server-macos-amd64 embedded_server.go

echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o course-server-macos-arm64 embedded_server.go

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o course-server-linux-amd64 embedded_server.go

echo "All builds complete!"
ls -la course-server-*
```

## Notes

- The resulting macOS binary will be executable on macOS without needing Go installed
- ARM64 version is for M1/M2 Macs (Apple Silicon)
- AMD64 version is for Intel-based Macs
- Most modern Macs can run the Intel version through Rosetta, but ARM64 is more efficient on Apple Silicon
- The embedded files (Astro frontend) are included in all cross-compiled binaries

## Testing Cross-Compilation

You can verify the build worked by checking the file:
```bash
file course-server-macos-amd64
# Should show: Mach-O 64-bit executable x86_64
```