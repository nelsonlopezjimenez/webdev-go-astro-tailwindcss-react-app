---
title: "WEEK 59 - Development vs Production modes" 
description: "WEEK 59 go server for develop/prod mode. This contains three to artifacts and a shell script. It does not match the original ver2 app. But it can read directory for future use if needed"
week: 59
---

# Lesson  59 Content Here

## Prompt
If you add a markdown file directly to the lessons folder:

✅ WILL WORK:
- API endpoints that read from filesystem (os.ReadFile)
- File listing APIs (os.ReadDir) 
- Any Go code that accesses actual filesystem

❌ WON'T WORK:
- Astro pages (they use embedded files from compile time)
- Static file serving from embedded FS
- Any functionality that relies on embedded files

🔄 INCONSISTENT BEHAVIOR:
- Your API might show the new lesson exists
- But your web pages won't display it
- This creates a confusing user experience

💾 TEMPORARY:
- The file will be lost if you rebuild without adding it to source
- It only exists in the runtime folder, not in your project

🔧 RECOMMENDATION:
Instead of adding files directly:
1. Add the file to your source lessons/ folder
2. Rebuild Astro: cd lessons/frontend && npm run build
3. Rebuild Go: go build -o course-server.exe *.go
4. Or use development mode to skip embedding

## // main.go - Updated with development mode support

```go
package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

// Global flag for development mode
var isDevelopment = flag.Bool("dev", false, "Run in development mode (reads files from disk)")
var port = flag.String("port", "8080", "Port to run server on")

func main() {
    flag.Parse()

    // Print mode information
    if *isDevelopment {
        log.Println("🔧 DEVELOPMENT MODE")
        log.Println("   - Reading lesson files from filesystem")
        log.Println("   - Changes to markdown files will be visible immediately")
        log.Println("   - Astro files still need to be rebuilt")
    } else {
        log.Println("🚀 PRODUCTION MODE")
        log.Println("   - Using embedded files")
        log.Println("   - All files are compiled into the executable")
    }

    // Setup routes
    setupRoutes()

    // Start server
    log.Printf("Server starting on http://localhost:%s", *port)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func setupRoutes() {
    // Static files (Astro frontend)
    http.HandleFunc("/", handleStaticFiles)
    
    // API routes
    http.HandleFunc("/api/sections/", handleSectionAPI)
    http.HandleFunc("/api/lessons/", handleLessonAPI)
    
    // Add other routes as needed
    // http.HandleFunc("/admin/", handleAdmin)
}
```

## // static_handlers.go - Handle static files in both modes

```go
package main

import (
    "embed"
    "io/fs"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

// Embedded files for production mode
//go:embed lessons/frontend/dist
var embeddedFiles embed.FS

func handleStaticFiles(w http.ResponseWriter, r *http.Request) {
    if *isDevelopment {
        handleStaticFilesDev(w, r)
    } else {
        handleStaticFilesProd(w, r)
    }
}

func handleStaticFilesDev(w http.ResponseWriter, r *http.Request) {
    // Development mode: serve from filesystem
    distPath := "lessons/frontend/dist"
    
    // Check if dist folder exists
    if _, err := os.Stat(distPath); os.IsNotExist(err) {
        http.Error(w, "Frontend not built. Run: cd lessons/frontend && npm run build", http.StatusNotFound)
        return
    }
    
    // Serve files from filesystem
    fs := http.Dir(distPath)
    handler := http.FileServer(fs)
    
    // Remove any prefix and serve
    http.StripPrefix("", handler).ServeHTTP(w, r)
    
    // Log requests in development
    log.Printf("📁 DEV: Serving %s from filesystem", r.URL.Path)
}

func handleStaticFilesProd(w http.ResponseWriter, r *http.Request) {
    // Production mode: serve from embedded files
    staticFS, err := fs.Sub(embeddedFiles, "lessons/frontend/dist")
    if err != nil {
        http.Error(w, "Embedded files not found", http.StatusInternalServerError)
        return
    }
    
    handler := http.FileServer(http.FS(staticFS))
    handler.ServeHTTP(w, r)
}
```

## // lesson_api.go - API handlers with development mode support
```go
package main

import (
    "embed"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
)

// Embedded lesson files for production
//go:embed lessons/section1-html-css lessons/section2-javascript lessons/section3-backend lessons/section4-react
var embeddedLessons embed.FS

// Lesson structure
type Lesson struct {
    Week        int    `json:"week"`
    Available   bool   `json:"available"`
    Title       string `json:"title"`
    Description string `json:"description"`
    FilePath    string `json:"file_path"`
}

func handleSectionAPI(w http.ResponseWriter, r *http.Request) {
    // Enable CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "OPTIONS" {
        return
    }

    // Parse URL path
    path := strings.TrimPrefix(r.URL.Path, "/api/sections/")
    parts := strings.Split(path, "/")

    if len(parts) < 2 || parts[1] != "lessons" {
        http.Error(w, "Invalid API endpoint", http.StatusNotFound)
        return
    }

    sectionName := parts[0]
    lessons := getSectionLessons(sectionName)

    // Log in development mode
    if *isDevelopment {
        log.Printf("🔍 DEV: Getting lessons for section %s, found %d lessons", sectionName, len(lessons))
    }

    json.NewEncoder(w).Encode(lessons)
}

func handleLessonAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    // Parse lesson request
    path := strings.TrimPrefix(r.URL.Path, "/api/lessons/")
    
    // Handle different lesson API endpoints
    switch {
    case strings.Contains(path, "/content"):
        handleLessonContent(w, r, path)
    case strings.Contains(path, "/list"):
        handleLessonList(w, r, path)
    default:
        http.Error(w, "Invalid lesson API endpoint", http.StatusNotFound)
    }
}

func getSectionLessons(sectionName string) []Lesson {
    var lessons []Lesson
    
    // Define week ranges for each section
    weekRanges := map[string][]int{
        "section1-html-css":   {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
        "section2-javascript": {13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
        "section3-backend":    {25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36},
        "section4-react":      {37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48},
    }

    weeks, exists := weekRanges[sectionName]
    if !exists {
        return lessons
    }

    // Check lessons based on mode
    if *isDevelopment {
        lessons = getLessonsFromFilesystem(sectionName, weeks)
    } else {
        lessons = getLessonsFromEmbedded(sectionName, weeks)
    }

    return lessons
}

func getLessonsFromFilesystem(sectionName string, weeks []int) []Lesson {
    var lessons []Lesson
    sectionPath := filepath.Join("lessons", sectionName)
    
    // Create directory if it doesn't exist (this is why empty folder gets created!)
    os.MkdirAll(sectionPath, 0755)
    
    for _, week := range weeks {
        lessonFile := filepath.Join(sectionPath, fmt.Sprintf("week%d.md", week))
        
        available := false
        title := fmt.Sprintf("Week %d", week)
        description := "Lesson content and exercises"

        // Check if file exists on filesystem
        if info, err := os.Stat(lessonFile); err == nil && !info.IsDir() {
            available = true
            if extractedTitle := extractTitleFromFile(lessonFile, true); extractedTitle != "" {
                title = extractedTitle
            }
            
            if *isDevelopment {
                log.Printf("📄 DEV: Found lesson file %s", lessonFile)
            }
        }

        lessons = append(lessons, Lesson{
            Week:        week,
            Available:   available,
            Title:       title,
            Description: description,
            FilePath:    lessonFile,
        })
    }

    return lessons
}

func getLessonsFromEmbedded(sectionName string, weeks []int) []Lesson {
    var lessons []Lesson
    
    for _, week := range weeks {
        lessonPath := fmt.Sprintf("lessons/%s/week%d.md", sectionName, week)
        
        available := false
        title := fmt.Sprintf("Week %d", week)
        description := "Lesson content and exercises"

        // Check if file exists in embedded FS
        if _, err := embeddedLessons.ReadFile(lessonPath); err == nil {
            available = true
            if extractedTitle := extractTitleFromFile(lessonPath, false); extractedTitle != "" {
                title = extractedTitle
            }
        }

        lessons = append(lessons, Lesson{
            Week:        week,
            Available:   available,
            Title:       title,
            Description: description,
            FilePath:    lessonPath,
        })
    }

    return lessons
}

func handleLessonContent(w http.ResponseWriter, r *http.Request, path string) {
    // Extract section and week from path
    // Example: section1-html-css/week1/content
    parts := strings.Split(strings.TrimSuffix(path, "/content"), "/")
    if len(parts) < 2 {
        http.Error(w, "Invalid lesson path", http.StatusBadRequest)
        return
    }

    sectionName := parts[0]
    weekName := parts[1]
    
    var content []byte
    var err error

    if *isDevelopment {
        // Read from filesystem
        filePath := filepath.Join("lessons", sectionName, weekName+".md")
        content, err = os.ReadFile(filePath)
        if err != nil {
            http.Error(w, "Lesson not found on filesystem", http.StatusNotFound)
            return
        }
        log.Printf("📖 DEV: Serving lesson content from %s", filePath)
    } else {
        // Read from embedded files
        embeddedPath := fmt.Sprintf("lessons/%s/%s.md", sectionName, weekName)
        content, err = embeddedLessons.ReadFile(embeddedPath)
        if err != nil {
            http.Error(w, "Lesson not found in embedded files", http.StatusNotFound)
            return
        }
    }

    w.Header().Set("Content-Type", "text/markdown")
    w.Write(content)
}

func extractTitleFromFile(filePath string, isFilesystem bool) string {
    var content []byte
    var err error

    if isFilesystem {
        content, err = os.ReadFile(filePath)
    } else {
        content, err = embeddedLessons.ReadFile(filePath)
    }

    if err != nil {
        return ""
    }

    // Look for YAML frontmatter title
    yamlTitleRegex := regexp.MustCompile(`(?m)^title:\s*["']?([^"'\n]+)["']?`)
    if match := yamlTitleRegex.FindStringSubmatch(string(content)); len(match) > 1 {
        return strings.TrimSpace(match[1])
    }

    // Look for first h1 heading
    h1Regex := regexp.MustCompile(`(?m)^#\s+(.+)$`)
    if match := h1Regex.FindStringSubmatch(string(content)); len(match) > 1 {
        return strings.TrimSpace(match[1])
    }

    return ""
}
```

## Shell script to run with previous code
```sh
# Development Scripts and Usage

# File: dev-start.sh (for easy development)
#!/bin/bash
echo "🔧 Starting Course Server in Development Mode"
echo ""

# Check if frontend is built
if [ ! -d "lessons/frontend/dist" ]; then
    echo "📦 Building Astro frontend first..."
    cd lessons/frontend
    npm run build
    cd ../..
    echo "✅ Frontend built successfully"
    echo ""
fi

echo "🚀 Starting Go server in development mode..."
echo "   📝 Lesson files will be read from filesystem"
echo "   🔄 Add/edit markdown files and see changes immediately"
echo "   🌐 Server will run at http://localhost:8080"
echo ""
echo "Press Ctrl+C to stop"
echo ""

# Start in development mode
go run *.go -dev

# File: build-prod.sh (for production build)
#!/bin/bash
echo "🚀 Building Course Server for Production"
echo ""

echo "📦 Building Astro frontend..."
cd lessons/frontend
npm run build
cd ../..

echo "🔨 Building Go executable..."
go build -ldflags="-s -w" -o course-server.exe *.go

echo "✅ Production build complete!"
echo "   📁 Executable: course-server.exe"
echo "   💾 All files embedded in executable"
echo ""
echo "To run: ./course-server.exe"

# File: quick-rebuild.sh (for when you modify Astro files)
#!/bin/bash
echo "🔄 Quick Astro rebuild..."
cd lessons/frontend
npm run build
cd ../..
echo "✅ Astro rebuilt! Restart your dev server to see changes."

# Windows batch files equivalent:

# File: dev-start.bat
@echo off
echo 🔧 Starting Course Server in Development Mode
echo.

if not exist "lessons\frontend\dist" (
    echo 📦 Building Astro frontend first...
    cd lessons\frontend
    npm run build
    cd ..\..
    echo ✅ Frontend built successfully
    echo.
)

echo 🚀 Starting Go server in development mode...
echo    📝 Lesson files will be read from filesystem
echo    🔄 Add/edit markdown files and see changes immediately
echo    🌐 Server will run at http://localhost:8080
echo.
echo Press Ctrl+C to stop
echo.

go run *.go -dev

# File: build-prod.bat
@echo off
echo 🚀 Building Course Server for Production
echo.

echo 📦 Building Astro frontend...
cd lessons\frontend
npm run build
cd ..\..

echo 🔨 Building Go executable...
go build -ldflags="-s -w" -o course-server.exe *.go

echo ✅ Production build complete!
echo    📁 Executable: course-server.exe
echo    💾 All files embedded in executable
echo.
echo To run: course-server.exe

# Usage Examples:

# Development workflow:
# 1. Start development server
chmod +x dev-start.sh
./dev-start.sh

# 2. Add a new lesson file (will be visible immediately)
echo "# Week 20: Advanced JavaScript" > lessons/section2-javascript/week20.md

# 3. Edit existing lesson (changes visible immediately)
nano lessons/section2-javascript/week13.md

# 4. If you modify Astro components, rebuild frontend:
cd lessons/frontend && npm run build && cd ../..

# Production deployment:
# 1. Build production executable
chmod +x build-prod.sh
./build-prod.sh

# 2. Deploy the executable (contains everything)
scp course-server.exe user@server:/path/to/deployment/
ssh user@server "/path/to/deployment/course-server.exe"

# Testing both modes:
# Development (reads from disk):
go run *.go -dev -port 8080

# Production (embedded files):
go build -o test-server.exe *.go && ./test-server.exe -port 8081

# Compare behaviors by visiting:
# http://localhost:8080 (development)
# http://localhost:8081 (production)
```