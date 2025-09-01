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
	"time"

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
}

func NewServer(lessonsDir string) *Server {
	return &Server{
		lessonsDir: lessonsDir,
		lessons:    make(map[int]*Lesson),
	}
}

func (s *Server) loadCourseInfo() error {
	courseFile := filepath.Join(s.lessonsDir, "course.yaml")
	if _, err := os.Stat(courseFile); os.IsNotExist(err) {
		// Default course info if file doesn't exist
		s.course = Course{
			Title:       "Programming Fundamentals",
			Description: "A comprehensive 10-week course covering programming fundamentals",
			Duration:    "10 weeks",
			Instructor:  "Course Instructor",
		}
		return nil
	}

	data, err := os.ReadFile(courseFile)
	if err != nil {
		return fmt.Errorf("failed to read course file: %w", err)
	}

	if err := yaml.Unmarshal(data, &s.course); err != nil {
		return fmt.Errorf("failed to parse course file: %w", err)
	}

	return nil
}

func (s *Server) scanLessons() error {
	s.lessons = make(map[int]*Lesson)

	err := filepath.WalkDir(s.lessonsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}

		lesson, err := s.parseLesson(path)
		if err != nil {
			log.Printf("Error parsing lesson %s: %v", path, err)
			return nil // Continue processing other lessons
		}

		if lesson.Week >= 1 && lesson.Week <= 10 {
			s.lessons[lesson.Week] = lesson
		}

		return nil
	})

	return err
}

func (s *Server) parseLesson(filePath string) (*Lesson, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)
	lesson := &Lesson{
		FilePath:  filePath,
		CreatedAt: time.Now(),
	}

	// Check if file has frontmatter
	if strings.HasPrefix(contentStr, "---") {
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) >= 3 {
			// Parse YAML frontmatter
			var metadata LessonMetadata
			if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err == nil {
				lesson.Title = metadata.Title
				lesson.Description = metadata.Description
				lesson.Week = metadata.Week
				lesson.Content = strings.TrimSpace(parts[2])
			}
		}
	}

	// If no frontmatter or parsing failed, try to infer from filename
	if lesson.Week == 0 {
		filename := filepath.Base(filePath)
		if weekNum := extractWeekFromFilename(filename); weekNum > 0 {
			lesson.Week = weekNum
		}
	}

	// Set default title if not provided
	if lesson.Title == "" {
		lesson.Title = fmt.Sprintf("Week %d Lesson", lesson.Week)
	}

	// Use full content if no frontmatter was found
	if lesson.Content == "" {
		lesson.Content = contentStr
	}

	// Get file modification time
	if stat, err := os.Stat(filePath); err == nil {
		lesson.CreatedAt = stat.ModTime()
	}

	return lesson, nil
}

func extractWeekFromFilename(filename string) int {
	// Try to extract week number from filename patterns like:
	// week1.md, week-1.md, 01-lesson.md, lesson_week_1.md, etc.
	lower := strings.ToLower(filename)
	
	// Remove extension
	lower = strings.TrimSuffix(lower, ".md")
	
	// Look for patterns
	patterns := []string{"week", "lesson", "chapter"}
	
	for _, pattern := range patterns {
		if strings.Contains(lower, pattern) {
			// Extract numbers from the string
			var num string
			for _, char := range lower {
				if char >= '0' && char <= '9' {
					num += string(char)
				}
			}
			if weekNum, err := strconv.Atoi(num); err == nil && weekNum >= 1 && weekNum <= 10 {
				return weekNum
			}
		}
	}
	
	return 0
}

func (s *Server) handleCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.course)
}

func (s *Server) handleLessons(w http.ResponseWriter, r *http.Request) {
	// Rescan lessons to pick up any new files
	if err := s.scanLessons(); err != nil {
		http.Error(w, "Failed to scan lessons", http.StatusInternalServerError)
		return
	}

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

	// Rescan lessons to pick up any new files
	if err := s.scanLessons(); err != nil {
		http.Error(w, "Failed to scan lessons", http.StatusInternalServerError)
		return
	}

	lesson, exists := s.lessons[week]
	if !exists {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

func (s *Server) handleSyllabus(w http.ResponseWriter, r *http.Request) {
	// Rescan lessons to get latest data
	if err := s.scanLessons(); err != nil {
		http.Error(w, "Failed to scan lessons", http.StatusInternalServerError)
		return
	}

	syllabus := struct {
		Course  Course             `json:"course"`
		Lessons map[int]*Lesson    `json:"lessons"`
		Weeks   []int             `json:"weeks"`
	}{
		Course:  s.course,
		Lessons: s.lessons,
		Weeks:   make([]int, 0, 10),
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

func (s *Server) setupRoutes() http.Handler {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)

	return corsHandler
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

	server := NewServer(lessonsDir)

	// Load course information
	if err := server.loadCourseInfo(); err != nil {
		log.Fatal("Failed to load course info:", err)
	}

	// Initial scan of lessons
	if err := server.scanLessons(); err != nil {
		log.Fatal("Failed to scan lessons:", err)
	}

	log.Printf("Starting server on :8080")
	log.Printf("Lessons directory: %s", lessonsDir)
	log.Printf("Found %d lessons", len(server.lessons))

	handler := server.setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", handler))
}