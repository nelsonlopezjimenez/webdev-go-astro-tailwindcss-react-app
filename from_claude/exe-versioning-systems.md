Here are several ways to track executable versions for your Go application:

## 1. Build-Time Version Information

Add version info during compilation using `-ldflags`:

```go
package main

import (
    "fmt"
    "runtime/debug"
)

var (
    version   = "dev"
    commit    = "none"
    date      = "unknown"
    builtBy   = "unknown"
)

func printVersion() {
    fmt.Printf("Version: %s\n", version)
    fmt.Printf("Commit: %s\n", commit)
    fmt.Printf("Built: %s\n", date)
    fmt.Printf("Built by: %s\n", builtBy)
}

// Add to your main() function
func main() {
    if len(os.Args) > 1 && os.Args[1] == "--version" {
        printVersion()
        return
    }
    // ... rest of your code
}
```

Build with version info:
```bash
VERSION=$(git describe --tags --always --dirty)
COMMIT=$(git rev-parse --short HEAD)
DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')
USER=$(whoami)

go build -ldflags="-s -w -X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE -X main.builtBy=$USER" -o course-server.exe
```

## 2. Build Script with Auto-Versioning

Create `build.sh`:
```bash
#!/bin/bash
set -e

# Get version from git
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "v0.0.0-unknown")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')
USER=$(whoami)

echo "Building version: $VERSION"

# Build Astro frontend
cd lessons/frontend
npm run build
cd ../..

# Build Go server with version info
go build -ldflags="-s -w -X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE -X main.builtBy=$USER" -o course-server.exe

echo "Build complete: course-server.exe"
echo "Version: $VERSION"
```

## 3. Add Version Endpoint

Add to your server:
```go
func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
    info := map[string]string{
        "version":    version,
        "commit":     commit,
        "build_date": date,
        "built_by":   builtBy,
        "go_version": runtime.Version(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(info)
}

// Add to setupRoutes():
api.HandleFunc("/version", s.handleVersion).Methods("GET")
```

## 4. Git-Based Versioning

Use git tags for version numbers:
```bash
# Create version tags
git tag v1.0.0
git tag v1.1.0

# Build script automatically uses latest tag
VERSION=$(git describe --tags --abbrev=0)
```

## 5. Startup Version Logging

Add to your main() function:
```go
log.Printf("Course Management System Server")
log.Printf("Version: %s", version)
log.Printf("Commit: %s", commit) 
log.Printf("Built: %s", date)
log.Printf("Lessons directory: %s", lessonsDir)
```

## 6. Executable Naming with Version

Include version in filename:
```bash
# Build with versioned name
VERSION=$(git describe --tags --always)
go build -ldflags="-s -w -X main.version=$VERSION" -o course-server-$VERSION.exe
```

## Recommended Approach

Combine methods 1, 2, and 3 for comprehensive versioning:
- Build-time version injection
- Automated build script
- HTTP version endpoint for runtime checking

This gives you version info in logs, via command line flags, and through your API.