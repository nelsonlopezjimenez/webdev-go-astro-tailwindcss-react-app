package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

//go:embed all:lessons/frontend/dist
var staticFiles embed.FS

type Course struct {
	Title        string   `json:"title" yaml:"title"`
	Description  string   `json:"description" yaml:"description"`
	Duration     string   `json:"duration" yaml:"duration"`
	Instructor   string   `json:"instructor" yaml:"instructor"`
	Requirements []string `json:"requirements" yaml:"requirements"`
}

type Lesson struct {
	Week        int       `json:"week"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	FilePath    string    `json:"file_path"`
	FileSize    int64     `json:"file_size"`
}

type LessonMetadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Week        int    `yaml:"week"`
}

type Server struct {
	lessonsDir string
	course     Course
	lessons    map[int]*Lesson
	mutex      sync.RWMutex
	watcher    *fsnotify.Watcher
}

func NewServer(lessonsDir string) (*Server, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}

	server := &Server{
		lessonsDir: lessonsDir,
		lessons:    make(map[int]*Lesson),
		watcher:    watcher,
	}

	// Add the lessons directory to the watcher if it exists
	if _, err := os.Stat(lessonsDir); err == nil {
		err = watcher.Add(lessonsDir)
		if err != nil {
			log.Printf("Warning: failed to watch lessons directory: %v", err)
		}
	} else {
		log.Printf("Lessons directory %s not found, creating it", lessonsDir)
		os.MkdirAll(lessonsDir, 0755)
		watcher.Add(lessonsDir)
	}

	return server, nil
}

func (s *Server) startFileWatcher() {
	go func() {
		for {
			select {
			case event, ok := <-s.watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
					if strings.HasSuffix(strings.ToLower(event.Name), ".md") ||
						strings.HasSuffix(strings.ToLower(event.Name), ".yaml") ||
						strings.HasSuffix(strings.ToLower(event.Name), ".yml") {

						log.Printf("🔍 Detected file change: %s", event.Name)
						time.Sleep(100 * time.Millisecond)

						if strings.Contains(event.Name, "course.yaml") || strings.Contains(event.Name, "course.yml") {
							if err := s.loadCourseInfo(); err != nil {
								log.Printf("❌ Error reloading course info: %v", err)
							}
						}

						if strings.HasSuffix(strings.ToLower(event.Name), ".md") {
							if err := s.scanLessons(); err != nil {
								log.Printf("❌ Error rescanning lessons: %v", err)
							} else {
								log.Printf("✅ Lessons updated. Found %d lessons", len(s.lessons))
							}
						}
					}
				}

			case err, ok := <-s.watcher.Errors:
				if !ok {
					return
				}
				log.Printf("❌ File watcher error: %v", err)
			}
		}
	}()
}

func (s *Server) loadCourseInfo() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	courseFile := filepath.Join(s.lessonsDir, "course.yaml")
	if _, err := os.Stat(courseFile); os.IsNotExist(err) {
		courseFile = filepath.Join(s.lessonsDir, "course.yml")
		if _, err := os.Stat(courseFile); os.IsNotExist(err) {
			s.course = Course{
				Title:       "Web Application Developer Certificate",
				Description: "A comprehensive program covering web development fundamentals",
				Duration:    "10 weeks",
				Instructor:  "Course Instructor",
				Requirements: []string{
					"Build and maintain websites.",
					"Work with stakeholders to create websites.",
					"Research, assess, and appropriately apply emerging technology to support websites as needed in industry.",
					"Comply with the ethics related to the use of copyrighted materials and intellectual property rights.",
					"Demonstrate an entrepreneurial approach to web development sites and pages.",
					"Manage career goals through creating effective resumes/CVs, developing interviewing skills, and setting goals.",
				},
			}
			log.Printf("📄 Using default course info (no course.yaml found)")
			return nil
		}
	}

	data, err := os.ReadFile(courseFile)
	if err != nil {
		return fmt.Errorf("failed to read course file: %w", err)
	}

	if err := yaml.Unmarshal(data, &s.course); err != nil {
		return fmt.Errorf("failed to parse course file: %w", err)
	}

	log.Printf("📚 Course info loaded: %s", s.course.Title)
	return nil
}

func (s *Server) scanLessons() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("🔍 Scanning lessons directory: %s", s.lessonsDir)

	newLessons := make(map[int]*Lesson)

	if _, err := os.Stat(s.lessonsDir); os.IsNotExist(err) {
		log.Printf("📁 Lessons directory doesn't exist, skipping scan")
		s.lessons = newLessons
		return nil
	}

	err := filepath.WalkDir(s.lessonsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}

		lesson, err := s.parseLesson(path)
		if err != nil {
			log.Printf("❌ Error parsing lesson %s: %v", path, err)
			return nil
		}

		if lesson.Week >= 1 && lesson.Week <= 100 {
			newLessons[lesson.Week] = lesson
			log.Printf("✅ Added lesson for week %d: %s", lesson.Week, lesson.Title)
		}

		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("📊 Found %d valid lessons", len(newLessons))
	s.lessons = newLessons
	return nil
}

func (s *Server) parseLesson(filePath string) (*Lesson, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)
	lesson := &Lesson{
		FilePath:  filePath,
		CreatedAt: fileInfo.ModTime(),
		FileSize:  fileInfo.Size(),
	}

	if strings.HasPrefix(contentStr, "---") {
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) >= 3 {
			var metadata LessonMetadata
			if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err == nil {
				lesson.Title = metadata.Title
				lesson.Description = metadata.Description
				lesson.Week = metadata.Week
				lesson.Content = strings.TrimSpace(parts[2])
			}
		}
	}

	if lesson.Week == 0 {
		filename := filepath.Base(filePath)
		if weekNum := extractWeekFromFilename(filename); weekNum > 0 {
			lesson.Week = weekNum
		}
	}

	if lesson.Title == "" {
		lesson.Title = fmt.Sprintf("Week %d Lesson", lesson.Week)
	}

	if lesson.Content == "" {
		lesson.Content = contentStr
	}

	return lesson, nil
}

func extractWeekFromFilename(filename string) int {
	lower := strings.ToLower(filename)
	lower = strings.TrimSuffix(lower, ".md")

	patterns := []string{"week", "lesson", "chapter"}

	for _, pattern := range patterns {
		if strings.Contains(lower, pattern) {
			var num string
			for _, char := range lower {
				if char >= '0' && char <= '9' {
					num += string(char)
				}
			}
			if weekNum, err := strconv.Atoi(num); err == nil && weekNum >= 1 && weekNum <= 100 {
				return weekNum
			}
		}
	}

	return 0
}

// API Handlers
func (s *Server) handleCourse(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	course := s.course
	s.mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func (s *Server) handleLessons(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var lessons []*Lesson
	for i := 1; i <= 100; i++ {
		if lesson, exists := s.lessons[i]; exists {
			lessons = append(lessons, lesson)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

func (s *Server) handleLesson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	weekStr := vars["week"]

	week, err := strconv.Atoi(weekStr)
	if err != nil || week < 1 || week > 100 {
		http.Error(w, "Invalid week number", http.StatusBadRequest)
		return
	}

	s.mutex.RLock()
	lesson, exists := s.lessons[week]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

func (s *Server) handleSyllabus(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	syllabus := struct {
		Course      Course          `json:"course"`
		Lessons     map[int]*Lesson `json:"lessons"`
		Weeks       []int           `json:"weeks"`
		LastUpdated time.Time       `json:"last_updated"`
		TotalFiles  int             `json:"total_files"`
	}{
		Course:      s.course,
		Lessons:     s.lessons,
		Weeks:       make([]int, 0),
		LastUpdated: time.Now(),
		TotalFiles:  len(s.lessons),
	}

	var weeks []int
	for week := range s.lessons {
		weeks = append(weeks, week)
	}
	sort.Ints(weeks)
	syllabus.Weeks = weeks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(syllabus)
}

// Serve embedded static files
func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
	// Remove /static prefix
	path := strings.TrimPrefix(r.URL.Path, "/static/")

	// Remove leading slash if present
	path = strings.TrimPrefix(path, "/")

	// Default to index.html for root requests
	if path == "" {
		path = "index.html"
	}

	// Construct the embedded file path
	embedPath := "lessons/frontend/dist/" + path

	// DEBUG: Log what we're trying to serve
	log.Printf("🔍 Requested path: %s", r.URL.Path)
	log.Printf("🔍 Cleaned path: %s", path)
	log.Printf("🔍 Embed path: %s", embedPath)

	// Try to read the file from embedded filesystem
	data, err := staticFiles.ReadFile(embedPath)
	if err != nil {
		log.Printf("❌ File not found: %s, serving index.html fallback", embedPath)
		indexPath := embedPath + "/index.html"
		// **************** to check whether index.html succeed or failed
		log.Printf("🔍 Trying index path: %s", indexPath)

		// If file not found, serve index.html (for SPA routing)
		data, err = staticFiles.ReadFile(indexPath)
		if err != nil {
			log.Printf("❌ Index file not found: %s, serving root index.html fallback", indexPath)

			// If still not found, serve root index.html (for SPA routing)
			data, err = staticFiles.ReadFile("lessons/frontend/dist/index.html")
			if err != nil {
				http.NotFound(w, r)
				return
			}
		} else {
			log.Printf("✅ Index file found: %s", indexPath)

		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write((data))
		return
	}
	// Set appropriate content type
	if strings.HasSuffix(path, ".css") {
		log.Printf("🎨 Setting CSS content type")
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	} else if strings.HasSuffix(path, ".js") {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	} else if strings.HasSuffix(path, ".html") {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	} else if strings.HasSuffix(path, ".json") {
		w.Header().Set("Content-Type", "application/json")
	} else if strings.HasSuffix(path, ".png") {
		w.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		w.Header().Set("Content-Type", "image/jpeg")
	} else if strings.HasSuffix(path, ".svg") {
		w.Header().Set("Content-Type", "image/svg+xml")
	} else if strings.HasSuffix(path, ".ico") {
		w.Header().Set("Content-Type", "image/x-icon")
	} else {
		// Default content type for other files
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	w.Write(data)
}

func (s *Server) setupRoutes() http.Handler {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")

	// Serve static files
	r.PathPrefix("/static/").HandlerFunc(s.handleStatic)

	// ****** The issue is that this current routing setup catches all paths handled by static file
	// ******** handler instead of being processed by SPA routes
	r.PathPrefix("/").Handler(http.HandlerFunc(s.handleStatic)) //

	// r.HandleFunc("/", s.handleStatic)

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)

	return corsHandler
}

func (s *Server) close() error {
	if s.watcher != nil {
		return s.watcher.Close()
	}
	return nil
}

func main() {
	lessonsDir := "./lessons"
	port := ":8080"

	if len(os.Args) > 1 {
		lessonsDir = os.Args[1]
	}
	if len(os.Args) > 2 {
		port = ":" + os.Args[2]
	}

	log.Println("🔍 Checking embedded files:")
	err := fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			log.Printf("📁 Embedded file: %s", path)
		}
		return nil
	})
	if err != nil {
		log.Printf("❌ Error walking embedded files: %v", err)
	}

	// Create lessons directory if it doesn't exist
	if err := os.MkdirAll(lessonsDir, 0755); err != nil {
		log.Printf("Warning: failed to create lessons directory: %v", err)
	}

	server, err := NewServer(lessonsDir)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}
	defer server.close()

	// Load initial data
	if err := server.loadCourseInfo(); err != nil {
		log.Printf("Warning: failed to load course info: %v", err)
	}

	if err := server.scanLessons(); err != nil {
		log.Printf("Warning: failed to scan lessons: %v", err)
	}

	// Start file watcher
	server.startFileWatcher()

	log.Printf("🚀 Course Management System Server")
	log.Printf("📁 Lessons directory: %s", lessonsDir)
	log.Printf("📚 Found %d lessons", len(server.lessons))
	log.Printf("🌐 Server starting on http://localhost%s", port)
	log.Printf("📖 Frontend: http://localhost%s", port)
	log.Printf("🔗 API: http://localhost%s/api", port)

	handler := server.setupRoutes()
	log.Fatal(http.ListenAndServe(port, handler))
}
