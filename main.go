package main

import (
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

type Course struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Duration    string `json:"duration" yaml:"duration"`
	Instructor  string `json:"instructor" yaml:"instructor"`
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

	// Add the lessons directory to the watcher
	err = watcher.Add(lessonsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to watch lessons directory: %w", err)
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

				// Only react to write and create events for .md and .yaml files
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
					if strings.HasSuffix(strings.ToLower(event.Name), ".md") ||
						strings.HasSuffix(strings.ToLower(event.Name), ".yaml") ||
						strings.HasSuffix(strings.ToLower(event.Name), ".yml") {

						log.Printf("🔍 Detected file change: %s", event.Name)

						// Small delay to ensure file write is complete
						time.Sleep(100 * time.Millisecond)

						// Reload course info and lessons
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
		// Try .yml extension
		courseFile = filepath.Join(s.lessonsDir, "course.yml")
		if _, err := os.Stat(courseFile); os.IsNotExist(err) {
			// Default course info if file doesn't exist
			s.course = Course{
				Title:       "Programming Fundamentals",
				Description: "A comprehensive 10-week course covering programming fundamentals",
				Duration:    "10 weeks",
				Instructor:  "Course Instructor",
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
	fileCount := 0

	err := filepath.WalkDir(s.lessonsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("❌ Error walking path %s: %v", path, err)
			return err
		}

		// Skip directories
		if d.IsDir() {
			log.Printf("📁 Skipping directory: %s", path)
			return nil
		}

		// Log all files found
		log.Printf("📄 Found file: %s", path)
		fileCount++

		// Check if it's a markdown file
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			log.Printf("⏭️  Skipping non-markdown file: %s", path)
			return nil
		}

		log.Printf("📝 Processing markdown file: %s", path)

		lesson, err := s.parseLesson(path)
		if err != nil {
			log.Printf("❌ Error parsing lesson %s: %v", path, err)
			return nil // Continue processing other lessons
		}

		log.Printf("📖 Parsed lesson - Week: %d, Title: %s", lesson.Week, lesson.Title)

		if lesson.Week >= 1 && lesson.Week <= 10 {
			newLessons[lesson.Week] = lesson
			log.Printf("✅ Added lesson for week %d: %s", lesson.Week, lesson.Title)
		} else {
			log.Printf("⚠️  Skipping lesson with invalid week number %d: %s", lesson.Week, path)
		}

		return nil
	})

	if err != nil {
		log.Printf("❌ Error during directory walk: %v", err)
		return err
	}

	log.Printf("📊 Scan complete - Files found: %d, Markdown files: %d, Valid lessons: %d",
		fileCount, fileCount-1 /* subtract directories */, len(newLessons))

	s.lessons = newLessons
	return nil
}

func (s *Server) parseLesson(filePath string) (*Lesson, error) {
	log.Printf("🔍 Parsing lesson file: %s", filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Get file info
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

	log.Printf("📄 File size: %d bytes", fileInfo.Size())

	// Check if file has frontmatter
	if strings.HasPrefix(contentStr, "---") {
		log.Printf("🏷️  Found YAML frontmatter in %s", filePath)
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) >= 3 {
			// Parse YAML frontmatter
			var metadata LessonMetadata
			if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err == nil {
				lesson.Title = metadata.Title
				lesson.Description = metadata.Description
				lesson.Week = metadata.Week
				lesson.Content = strings.TrimSpace(parts[2])
				log.Printf("✅ YAML parsed - Week: %d, Title: %s", metadata.Week, metadata.Title)
			} else {
				log.Printf("❌ Failed to parse YAML frontmatter: %v", err)
			}
		} else {
			log.Printf("⚠️  Malformed frontmatter in %s", filePath)
		}
	} else {
		log.Printf("📝 No frontmatter found, trying filename detection")
	}

	// If no frontmatter or parsing failed, try to infer from filename
	if lesson.Week == 0 {
		filename := filepath.Base(filePath)
		weekNum := extractWeekFromFilename(filename)
		log.Printf("🔍 Filename '%s' -> Week number: %d", filename, weekNum)
		if weekNum > 0 {
			lesson.Week = weekNum
		}
	}

	// Set default title if not provided
	if lesson.Title == "" {
		lesson.Title = fmt.Sprintf("Week %d Lesson", lesson.Week)
		log.Printf("📝 Using default title: %s", lesson.Title)
	}

	// Use full content if no frontmatter was found
	if lesson.Content == "" {
		lesson.Content = contentStr
		log.Printf("📄 Using full file content (no frontmatter)")
	}

	if lesson.Week == 0 {
		log.Printf("❌ Could not determine week number for %s", filePath)
		return nil, fmt.Errorf("could not determine week number for %s", filePath)
	}

	log.Printf("✅ Successfully parsed lesson: Week %d, Title: %s", lesson.Week, lesson.Title)
	return lesson, nil
}

func extractWeekFromFilename(filename string) int {
	log.Printf("🔍 Extracting week from filename: %s", filename)

	// Try to extract week number from filename patterns like:
	// week1.md, week-1.md, 01-lesson.md, lesson_week_1.md, etc.
	lower := strings.ToLower(filename)

	// Remove extension
	lower = strings.TrimSuffix(lower, ".md")
	log.Printf("   After removing .md: %s", lower)

	// Look for patterns
	patterns := []string{"week", "lesson", "chapter"}

	for _, pattern := range patterns {
		if strings.Contains(lower, pattern) {
			log.Printf("   Found pattern '%s'", pattern)
			// Extract numbers from the string
			var num string
			for _, char := range lower {
				if char >= '0' && char <= '9' {
					num += string(char)
				}
			}
			log.Printf("   Extracted number string: %s", num)
			if weekNum, err := strconv.Atoi(num); err == nil && weekNum >= 1 && weekNum <= 10 {
				log.Printf("   ✅ Week number: %d", weekNum)
				return weekNum
			}
		}
	}

	log.Printf("   ❌ No valid week number found")
	return 0
}

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

	// Convert map to slice and sort by week
	var lessons []*Lesson
	for i := 1; i <= 10; i++ {
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
	if err != nil || week < 1 || week > 10 {
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
		Weeks:       make([]int, 0, 10),
		LastUpdated: time.Now(),
		TotalFiles:  len(s.lessons),
	}

	// Add available weeks in order
	var weeks []int
	for week := range s.lessons {
		weeks = append(weeks, week)
	}
	sort.Ints(weeks)
	syllabus.Weeks = weeks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(syllabus)
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	status := struct {
		Status         string    `json:"status"`
		LessonsDir     string    `json:"lessons_dir"`
		TotalLessons   int       `json:"total_lessons"`
		LastScanned    time.Time `json:"last_scanned"`
		AvailableWeeks []int     `json:"available_weeks"`
	}{
		Status:         "healthy",
		LessonsDir:     s.lessonsDir,
		TotalLessons:   len(s.lessons),
		LastScanned:    time.Now(),
		AvailableWeeks: make([]int, 0, len(s.lessons)),
	}

	for week := range s.lessons {
		status.AvailableWeeks = append(status.AvailableWeeks, week)
	}
	sort.Ints(status.AvailableWeeks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (s *Server) setupRoutes() http.Handler {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")
	api.HandleFunc("/status", s.handleStatus).Methods("GET")

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)

	return corsHandler
}

func (s *Server) close() error {
	return s.watcher.Close()
}

func main() {
	lessonsDir := "./lessons"
	if len(os.Args) > 1 {
		lessonsDir = os.Args[1]
	}

	// Create lessons directory if it doesn't exist
	if err := os.MkdirAll(lessonsDir, 0755); err != nil {
		log.Fatal("Failed to create lessons directory:", err)
	}

	server, err := NewServer(lessonsDir)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}
	defer server.close()

	// Load course information
	if err := server.loadCourseInfo(); err != nil {
		log.Fatal("Failed to load course info:", err)
	}

	// Initial scan of lessons
	if err := server.scanLessons(); err != nil {
		log.Fatal("Failed to scan lessons:", err)
	}

	// Start file watcher
	server.startFileWatcher()

	log.Printf("🚀 Server starting on :8080")
	log.Printf("📁 Lessons directory: %s", lessonsDir)
	log.Printf("📚 Found %d lessons", len(server.lessons))
	log.Printf("👀 File watcher active - will detect new lessons automatically")
	log.Printf("🔗 Available endpoints:")
	log.Printf("   GET /api/course - Course information")
	log.Printf("   GET /api/lessons - All lessons")
	log.Printf("   GET /api/lessons/{week} - Specific lesson")
	log.Printf("   GET /api/syllabus - Complete syllabus")
	log.Printf("   GET /api/status - Server status")

	handler := server.setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
