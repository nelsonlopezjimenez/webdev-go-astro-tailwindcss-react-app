# Multiple go files

## Prompt
Go server with astro embedded. I want to add functionality. Do I have to do it in the main.go file? can be imported from another file.go file?

## Claude 
Yes, absolutely! You can (and should) organize your Go server functionality into separate files. This is much better than cramming everything into `main.go`. Here are several approaches:## Recommended Approach for Your Project

For your course server, I recommend the **simple same-package approach** (the third artifact). Here's why:

### Benefits:
1. **No import path complexity** - All files in same package
2. **Easy to build** - Just `go build *.go` 
3. **Clean organization** - Related functions grouped in files
4. **Embedded assets work easily** - No path issues with `//go:embed`

### File Organization:
```
├── main.go              # Entry point only
├── routes.go            # Route setup
├── api_handlers.go      # Existing API endpoints  
├── static_handlers.go   # Static file serving (embedded Astro)
├── new_features.go      # Your new functionality
├── progress_service.go  # Student progress tracking
├── search_service.go    # Search functionality
└── utils.go            # Shared utilities
```

### Adding New Functionality:

Just create a new `.go` file with functions, then register routes in `routes.go`:

```go
// In your new file: analytics.go
package main

func handleAnalyticsAPI(w http.ResponseWriter, r *http.Request) {
    // Your analytics functionality
}

// In routes.go, add:
http.HandleFunc("/api/analytics/", handleAnalyticsAPI)
```

### Build Command:
```bash
go build -o course-server.exe *.go
```

**What specific functionality do you want to add?** I can show you exactly how to implement it with this structure.

## Go Project Structure with Separate Files
```sh
// Project structure for organized Go server

// File: main.go
package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize routes
    setupRoutes()
    
    // Start server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// File: routes.go
package main

import (
    "net/http"
)

func setupRoutes() {
    // Static files (embedded Astro)
    setupStaticRoutes()
    
    // API routes
    setupAPIRoutes()
    
    // Custom functionality routes
    setupCustomRoutes()
}

func setupStaticRoutes() {
    // Your existing embedded Astro static file serving
    http.HandleFunc("/", serveStaticFiles)
}

func setupAPIRoutes() {
    http.HandleFunc("/api/sections/", handleSectionAPI)
    http.HandleFunc("/api/lessons/", handleLessonAPI)
    http.HandleFunc("/api/progress/", handleProgressAPI)
}

func setupCustomRoutes() {
    http.HandleFunc("/custom/", handleCustomFunction)
    http.HandleFunc("/admin/", handleAdminFunction)
}

// File: handlers.go
package main

import (
    "encoding/json"
    "net/http"
)

func handleSectionAPI(w http.ResponseWriter, r *http.Request) {
    // Your existing section API logic
}

func handleLessonAPI(w http.ResponseWriter, r *http.Request) {
    // New lesson API functionality
}

func handleProgressAPI(w http.ResponseWriter, r *http.Request) {
    // Student progress tracking
}

// File: custom_handlers.go
package main

import (
    "net/http"
)

func handleCustomFunction(w http.ResponseWriter, r *http.Request) {
    // Your new custom functionality here
}

func handleAdminFunction(w http.ResponseWriter, r *http.Request) {
    // Admin panel functionality
}

// File: static.go
package main

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed lessons/frontend/dist
var staticFiles embed.FS

func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
    // Your existing static file serving logic
    staticFS, _ := fs.Sub(staticFiles, "lessons/frontend/dist")
    http.FileServer(http.FS(staticFS)).ServeHTTP(w, r)
}
```

## Advanced Go Project Structure with Packages
```sh
// Advanced structure using separate packages

// Project structure:
// ├── main.go
// ├── go.mod
// ├── handlers/
// │   ├── api.go
// │   ├── lessons.go
// │   └── static.go
// ├── models/
// │   ├── lesson.go
// │   └── section.go
// ├── services/
// │   ├── lesson_service.go
// │   └── progress_service.go
// └── utils/
//     ├── file_utils.go
//     └── response_utils.go

// File: main.go
package main

import (
    "log"
    "net/http"
    "./handlers"  // Import your local packages
)

func main() {
    // Setup routes using handlers package
    handlers.SetupRoutes()
    
    log.Println("Course server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// File: handlers/api.go
package handlers

import (
    "encoding/json"
    "net/http"
    "../models"
    "../services"
)

func HandleSectionAPI(w http.ResponseWriter, r *http.Request) {
    sections := services.GetAllSections()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sections)
}

func HandleLessonAPI(w http.ResponseWriter, r *http.Request) {
    // Lesson-specific API logic
}

// File: handlers/static.go
package handlers

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed ../lessons/frontend/dist
var staticFiles embed.FS

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
    staticFS, _ := fs.Sub(staticFiles, "lessons/frontend/dist")
    http.FileServer(http.FS(staticFS)).ServeHTTP(w, r)
}

func SetupRoutes() {
    http.HandleFunc("/", ServeStaticFiles)
    http.HandleFunc("/api/sections/", HandleSectionAPI)
    http.HandleFunc("/api/lessons/", HandleLessonAPI)
}

// File: models/lesson.go
package models

type Lesson struct {
    Week        int    `json:"week"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Available   bool   `json:"available"`
    Section     string `json:"section"`
}

type Section struct {
    ID          string   `json:"id"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Weeks       []int    `json:"weeks"`
    Lessons     []Lesson `json:"lessons"`
}

// File: services/lesson_service.go
package services

import (
    "os"
    "path/filepath"
    "../models"
)

func GetLessonsBySection(sectionName string) []models.Lesson {
    var lessons []models.Lesson
    
    sectionPath := filepath.Join("lessons", sectionName)
    files, err := os.ReadDir(sectionPath)
    if err != nil {
        return lessons
    }
    
    for _, file := range files {
        if filepath.Ext(file.Name()) == ".md" {
            // Parse lesson from file
            lesson := parseLessonFromFile(filepath.Join(sectionPath, file.Name()))
            lessons = append(lessons, lesson)
        }
    }
    
    return lessons
}

func parseLessonFromFile(filePath string) models.Lesson {
    // Implementation to parse lesson metadata from markdown file
    return models.Lesson{}
}

// File: utils/response_utils.go
package utils

import (
    "encoding/json"
    "net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(data)
}

func SendErrorResponse(w http.ResponseWriter, message string, code int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(map[string]string{"error": message})
}
```

## Simple Same-Package Approach (Recommended by Claude)
```sh
// Simple approach: Multiple files, same package (recommended for your project)

// File: main.go
package main

import (
    "log"
    "net/http"
)

func main() {
    setupAllRoutes()
    
    log.Println("Course server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// File: routes.go
package main

import "net/http"

func setupAllRoutes() {
    // Static files (your existing Astro app)
    http.HandleFunc("/", handleStaticFiles)
    
    // API endpoints
    http.HandleFunc("/api/sections/", handleSectionAPI)
    http.HandleFunc("/api/lessons/", handleLessonAPI)
    
    // New functionality you want to add
    http.HandleFunc("/api/progress/", handleProgressAPI)
    http.HandleFunc("/api/search/", handleSearchAPI)
    http.HandleFunc("/admin/", handleAdminPanel)
}

// File: api_handlers.go
package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

func handleSectionAPI(w http.ResponseWriter, r *http.Request) {
    // Your existing section API code
}

func handleLessonAPI(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/api/lessons/")
    
    switch {
    case strings.HasSuffix(path, "/content"):
        handleLessonContent(w, r)
    case strings.HasSuffix(path, "/progress"):
        handleLessonProgress(w, r)
    default:
        handleLessonList(w, r)
    }
}

func handleLessonContent(w http.ResponseWriter, r *http.Request) {
    // Serve lesson markdown content
}

func handleLessonProgress(w http.ResponseWriter, r *http.Request) {
    // Track student progress
}

func handleLessonList(w http.ResponseWriter, r *http.Request) {
    // List lessons for a section
}

// File: new_features.go
package main

import (
    "encoding/json"
    "net/http"
)

func handleProgressAPI(w http.ResponseWriter, r *http.Request) {
    // New feature: Student progress tracking
    switch r.Method {
    case "GET":
        getStudentProgress(w, r)
    case "POST":
        updateStudentProgress(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleSearchAPI(w http.ResponseWriter, r *http.Request) {
    // New feature: Search through lessons
    query := r.URL.Query().Get("q")
    results := searchLessons(query)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

func handleAdminPanel(w http.ResponseWriter, r *http.Request) {
    // New feature: Admin panel
    if !isAuthorized(r) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Serve admin interface
}

// File: progress_service.go
package main

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type StudentProgress struct {
    StudentID        string `json:"student_id"`
    CompletedWeeks   []int  `json:"completed_weeks"`
    CurrentWeek      int    `json:"current_week"`
    TotalTimeSpent   int    `json:"total_time_spent"`
    LastAccessDate   string `json:"last_access_date"`
}

func getStudentProgress(w http.ResponseWriter, r *http.Request) {
    studentID := r.URL.Query().Get("student_id")
    progress := loadStudentProgress(studentID)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(progress)
}

func updateStudentProgress(w http.ResponseWriter, r *http.Request) {
    var progress StudentProgress
    if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    saveStudentProgress(progress)
    w.WriteHeader(http.StatusOK)
}

func loadStudentProgress(studentID string) StudentProgress {
    // Load from file or database
    progressFile := filepath.Join("data", "progress", studentID+".json")
    data, err := os.ReadFile(progressFile)
    if err != nil {
        return StudentProgress{StudentID: studentID}
    }
    
    var progress StudentProgress
    json.Unmarshal(data, &progress)
    return progress
}

func saveStudentProgress(progress StudentProgress) error {
    os.MkdirAll(filepath.Join("data", "progress"), 0755)
    progressFile := filepath.Join("data", "progress", progress.StudentID+".json")
    
    data, _ := json.MarshalIndent(progress, "", "  ")
    return os.WriteFile(progressFile, data, 0644)
}

// File: search_service.go
package main

import (
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

type SearchResult struct {
    Week        int    `json:"week"`
    Section     string `json:"section"`
    Title       string `json:"title"`
    Excerpt     string `json:"excerpt"`
    MatchCount  int    `json:"match_count"`
}

func searchLessons(query string) []SearchResult {
    var results []SearchResult
    
    // Search through all lesson files
    lessonDirs := []string{"section1-html-css", "section2-javascript", "section3-backend", "section4-react"}
    
    for _, dir := range lessonDirs {
        sectionPath := filepath.Join("lessons", dir)
        files, _ := os.ReadDir(sectionPath)
        
        for _, file := range files {
            if strings.HasSuffix(file.Name(), ".md") {
                result := searchInFile(filepath.Join(sectionPath, file.Name()), query, dir)
                if result.MatchCount > 0 {
                    results = append(results, result)
                }
            }
        }
    }
    
    return results
}

func searchInFile(filePath, query, section string) SearchResult {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return SearchResult{}
    }
    
    text := string(content)
    matches := regexp.MustCompile(`(?i)`+regexp.QuoteMeta(query)).FindAllString(text, -1)
    
    // Extract week number from filename
    weekRegex := regexp.MustCompile(`week(\d+)\.md`)
    weekMatch := weekRegex.FindStringSubmatch(filepath.Base(filePath))
    week := 0
    if len(weekMatch) > 1 {
        week = parseInt(weekMatch[1])
    }
    
    // Create excerpt around first match
    excerpt := createExcerpt(text, query)
    
    return SearchResult{
        Week:       week,
        Section:    section,
        Title:      extractTitle(text),
        Excerpt:    excerpt,
        MatchCount: len(matches),
    }
}

func createExcerpt(text, query string) string {
    // Implementation to create excerpt around search term
    return ""
}

func extractTitle(content string) string {
    // Extract title from markdown content
    return ""
}

func parseInt(s string) int {
    // Convert string to int safely
    return 0
}
```