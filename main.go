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
	"regexp"
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
	Section     string    `json:"section"`      // New field
	SectionName string    `json:"section_name"` // New field
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
	Section     string `yaml:"section"`
}

type Section struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	WeekStart   int       `json:"week_start"`
	WeekEnd     int       `json:"week_end"`
	Lessons     []*Lesson `json:"lessons"`
}

type Server struct {
	lessonsDir string
	course     Course
	lessons    map[int]*Lesson     // Keep existing week-based mapping
	sections   map[string]*Section // New section-based mapping
	mutex      sync.RWMutex
	watcher    *fsnotify.Watcher
}

// Add these structs after your existing structs (after Section struct)

type TOCItem struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Level    int       `json:"level"`
	Children []TOCItem `json:"children,omitempty"`
}

type TOCResponse struct {
	TOCItems []TOCItem `json:"tocItems"`
	Source   string    `json:"source"` // "markdown", "default", or "error"
	Week     int       `json:"week"`
	Section  string    `json:"section"`
}

type InstructorInfo struct {
	Name             string   `json:"name"`
	TelephoneNumbers []string `json:"telephone_numbers"`
	Emails           []string `json:"emails"`
	PreferredContact string   `json:"preferred_contact"`
	ResponseTime     string   `json:"response_time"`
}

func getInstructorInfo() InstructorInfo {
	return InstructorInfo{
		Name:             "Nelson Lopez",
		TelephoneNumbers: []string{"TRU: x42478", "TRU: x44215"},
		Emails:           []string{"ndlopezjimenez@doc1.wa.gov", "nelson.lopez-jimenez@edmonds.edu"},
		PreferredContact: "Kiosk messaging system",
		ResponseTime:     "Within 24-hours on weekdays",
	}
}

func NewServer(lessonsDir string) (*Server, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create file watcher: %w", err)
	}

	server := &Server{
		lessonsDir: lessonsDir,
		lessons:    make(map[int]*Lesson),
		sections:   make(map[string]*Section),
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

						log.Printf("Detected file change: %s", event.Name)
						time.Sleep(100 * time.Millisecond)

						if strings.Contains(event.Name, "course.yaml") || strings.Contains(event.Name, "course.yml") {
							if err := s.loadCourseInfo(); err != nil {
								log.Printf("Error reloading course info: %v", err)
							}
						}

						if strings.HasSuffix(strings.ToLower(event.Name), ".md") {
							if err := s.scanLessons(); err != nil {
								log.Printf("Error rescanning lessons: %v", err)
							} else {
								log.Printf("Lessons updated. Found %d lessons in %d sections", len(s.lessons), len(s.sections))
							}
						}
					}
				}

			case err, ok := <-s.watcher.Errors:
				if !ok {
					return
				}
				log.Printf("File watcher error: %v", err)
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
				Duration:    "48 weeks (4 sections)",
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
			log.Printf("Using default course info (no course.yaml found)")
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

	log.Printf("Course info loaded: %s", s.course.Title)
	return nil
}

func (s *Server) scanLessons() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("Scanning lessons directory: %s", s.lessonsDir)

	newLessons := make(map[int]*Lesson)
	newSections := make(map[string]*Section)

	// Initialize sections
	sectionConfigs := map[string]struct {
		name      string
		weekStart int
		weekEnd   int
	}{
		"section1-html-css":   {"HTML/CSS Fundamentals", 1, 12},
		"section2-javascript": {"JavaScript Programming", 13, 24},
		"section3-backend":    {"Backend Development", 25, 36},
		"section4-react":      {"React & Frontend", 37, 48},
		"section5":            {"Section 5", 49, 60},
	}

	for sectionID, config := range sectionConfigs {
		newSections[sectionID] = &Section{
			ID:          sectionID,
			Name:        config.name,
			Description: fmt.Sprintf("Weeks %d-%d", config.weekStart, config.weekEnd),
			WeekStart:   config.weekStart,
			WeekEnd:     config.weekEnd,
			Lessons:     []*Lesson{},
		}
	}

	if _, err := os.Stat(s.lessonsDir); os.IsNotExist(err) {
		log.Printf("Lessons directory doesn't exist, skipping scan")
		s.lessons = newLessons
		s.sections = newSections
		return nil
	}

	// Scan section directories
	for sectionID, section := range newSections {
		sectionPath := filepath.Join(s.lessonsDir, sectionID)

		if _, err := os.Stat(sectionPath); os.IsNotExist(err) {
			log.Printf("Section directory %s doesn't exist, skipping", sectionID)
			continue
		}

		err := filepath.WalkDir(sectionPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
				return nil
			}

			lesson, err := s.parseLesson(path, sectionID, section.Name, section.WeekStart)
			if err != nil {
				log.Printf("Error parsing lesson %s: %v", path, err)
				return nil
			}

			if lesson.Week >= section.WeekStart && lesson.Week <= section.WeekEnd {
				newLessons[lesson.Week] = lesson
				section.Lessons = append(section.Lessons, lesson)
				log.Printf("Added lesson for week %d in %s: %s", lesson.Week, sectionID, lesson.Title)
			}

			return nil
		})

		if err != nil {
			log.Printf("Error scanning section %s: %v", sectionID, err)
		}

		// Sort lessons within section
		sort.Slice(section.Lessons, func(i, j int) bool {
			return section.Lessons[i].Week < section.Lessons[j].Week
		})
	}

	// Also scan for legacy lessons (week1.md, week2.md, etc. in root lessons dir)
	legacyFiles, err := filepath.Glob(filepath.Join(s.lessonsDir, "week*.md"))
	if err == nil {
		for _, filePath := range legacyFiles {
			lesson, err := s.parseLesson(filePath, "", "", 1)
			if err != nil {
				log.Printf("Error parsing legacy lesson %s: %v", filePath, err)
				continue
			}

			if lesson.Week >= 1 && lesson.Week <= 48 {
				// Determine which section this lesson belongs to
				for sectionID, section := range newSections {
					if lesson.Week >= section.WeekStart && lesson.Week <= section.WeekEnd {
						lesson.Section = sectionID
						lesson.SectionName = section.Name
						newLessons[lesson.Week] = lesson
						section.Lessons = append(section.Lessons, lesson)
						log.Printf("Added legacy lesson for week %d: %s", lesson.Week, lesson.Title)
						break
					}
				}
			}
		}
	}

	log.Printf("Found %d valid lessons across %d sections", len(newLessons), len(newSections))
	s.lessons = newLessons
	s.sections = newSections
	return nil
}

func (s *Server) parseLesson(filePath, sectionID, sectionName string, weekOffset int) (*Lesson, error) {
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
		FilePath:    filePath,
		CreatedAt:   fileInfo.ModTime(),
		FileSize:    fileInfo.Size(),
		Section:     sectionID,
		SectionName: sectionName,
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
				if metadata.Section != "" {
					lesson.Section = metadata.Section
				}
			}
		}
	}

	if lesson.Week == 0 {
		filename := filepath.Base(filePath)
		if weekNum := extractWeekFromFilename(filename); weekNum > 0 {
			// For section files, adjust week number based on section
			if sectionID != "" && weekOffset > 1 {
				lesson.Week = weekOffset + weekNum - 1
			} else {
				lesson.Week = weekNum
			}
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
			if weekNum, err := strconv.Atoi(num); err == nil && weekNum >= 1 && weekNum <= 48 {
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
	for i := 1; i <= 60; i++ {
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
	if err != nil || week < 1 || week > 60 {
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

// New section-based handlers
func (s *Server) handleSections(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var sections []*Section
	sectionOrder := []string{"section1-html-css", "section2-javascript", "section3-backend", "section4-react"}

	for _, sectionID := range sectionOrder {
		if section, exists := s.sections[sectionID]; exists {
			sections = append(sections, section)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sections)
}

func (s *Server) handleSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["section"]

	s.mutex.RLock()
	section, exists := s.sections[sectionID]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(section)
}

func (s *Server) handleSectionLesson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["section"]
	weekStr := vars["week"]

	week, err := strconv.Atoi(weekStr)
	if err != nil || week < 1 || week > 12 {
		http.Error(w, "Invalid week number", http.StatusBadRequest)
		return
	}

	s.mutex.RLock()
	section, exists := s.sections[sectionID]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	// Convert section week (1-12) to global week
	globalWeek := section.WeekStart + week - 1

	s.mutex.RLock()
	lesson, exists := s.lessons[globalWeek]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

// Add this new handler to your main.go

func (s *Server) handleSectionSyllabus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["section"]

	s.mutex.RLock()
	section, exists := s.sections[sectionID]
	s.mutex.RUnlock()

	if !exists {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	// Section-specific course information
	sectionInfo := map[string]struct {
		CourseCode    string   `json:"course_code"`
		Credits       string   `json:"credits"`
		Prerequisites string   `json:"prerequisites"`
		Description   string   `json:"description"`
		Objectives    []string `json:"objectives"`
		Topics        []string `json:"topics"`
		Assessment    []string `json:"assessment"`
		Resources     []string `json:"resources"`
	}{
		"section1-html-css": {
			CourseCode:    "CIS 241 ART 225 CIS 291",
			Credits:       "11.0 Credits",
			Prerequisites: "CIS 100 or instructor permission",
			Description:   "Website development using current HTML languages, approached from a source code perspective. Covers tags, forms, linked objects, current CSS, frames, tables, and an introduction to scripting.",
			Objectives: []string{
				"Perform content design and technical analysis on web applications and websites",
				"Use current HTML to develop, debug, maintain, and document web applications and websites",
				"Compare and contrast different browsers' effects on current HTML documents",
				"Use current HTML forms, iframes, and tables",
				"Create current HTML style through inline, embedded, and Cascading Style Sheets",
			},
			Topics: []string{
				"HTML5 semantic elements and document structure",
				"CSS fundamentals: selectors, properties, and values",
				"CSS layout techniques: Flexbox and Grid",
				"Responsive web design and media queries",
				"Web accessibility principles and best practices",
				"Form creation and validation",
				"CSS animations and transitions",
				"Browser compatibility and testing",
				"Version control with Git and GitHub",
				"Web development tools and workflow",
				"Performance optimization basics",
				"Final project: Complete responsive website",
			},
			Assessment: []string{
				"Weekly coding assignments (40%)",
				"Midterm project: Multi-page website (20%)",
				"Final project: Responsive portfolio site (25%)",
				"Lab exercises and participation (15%)",
			},
			Resources: []string{
				"MDN Web Docs - HTML/CSS Reference",
				"W3Schools - HTML/CSS Tutorials",
				"Can I Use - Browser compatibility tables",
				"CSS-Tricks - CSS techniques and guides",
				"GitHub - Version control and project hosting",
			},
		},
		"section2-javascript": {
			CourseCode:    "CIS 242",
			Credits:       "5.0 Credits",
			Prerequisites: "CIS 241 with a minimum grade of 2.5 or instructor permission",
			Description:   "Students will explore embedding, inline and external scripts, functions, form validation, loops, conditional statements, strings, numbers, and DHTML. Introduction to JavaScript Frameworks.",
			Objectives: []string{
				"Use object-oriented client-side scripting with well-formed web pages",
				"Recognize client-side variables and data types and operations",
				"Write client-side functions, event handlers, and control structures",
				"Verify form data through scripting validation",
				"Save state information through hidden fields, query-strings, and cookies",
				"List concepts of server-side programming and Node.js",
			},
			Topics: []string{
				"JavaScript fundamentals: variables, data types, operators",
				"Functions and scope",
				"DOM manipulation and event handling",
				"Control structures: loops and conditionals",
				"Arrays and objects",
				"Form validation and user input handling",
				"Asynchronous JavaScript: callbacks, promises, async/await",
				"ES6+ features: arrow functions, destructuring, modules",
				"Local storage and session management",
				"Introduction to JavaScript frameworks",
				"Debugging and testing techniques",
				"Final project: Interactive web application",
			},
			Assessment: []string{
				"Weekly programming exercises (35%)",
				"Midterm exam: JavaScript fundamentals (20%)",
				"Interactive web app project (30%)",
				"Lab work and code reviews (15%)",
			},
			Resources: []string{
				"MDN Web Docs - JavaScript Reference",
				"JavaScript.info - Modern JavaScript tutorial",
				"W3Schools - JavaScript tutorials and examples",
				"CodePen - JavaScript code playground",
				"Chrome DevTools - Debugging and testing",
			},
		},
		"section3-backend": {
			CourseCode:    "CIS 243",
			Credits:       "5.0 Credits",
			Prerequisites: "CIS 242 with a minimum grade of 2.5 or instructor permission",
			Description:   "Server-side scripting fundamentals including functions, logical structure, database connectivity, Object-Oriented principles, relational databases, and web frameworks.",
			Objectives: []string{
				"Understand difference between client-side and server-side scripting",
				"Use appropriate script types to complete interactive websites with data repositories",
				"Use Model, View, Controller (MVC) principles and architecture",
				"Use operators including logical operators and variables in scripting language",
				"Create procedures and reusable code in scripting language",
				"Create websites using web frameworks",
			},
			Topics: []string{
				"Node.js runtime and npm package management",
				"Express.js framework and routing",
				"Database design and MongoDB integration",
				"RESTful API development",
				"Authentication and authorization",
				"Middleware and error handling",
				"Data validation and sanitization",
				"File uploads and processing",
				"Environment configuration and deployment",
				"Testing strategies for backend applications",
				"Security best practices",
				"Final project: Full-stack CRUD application",
			},
			Assessment: []string{
				"API development assignments (40%)",
				"Database design project (20%)",
				"Full-stack application (25%)",
				"Technical documentation and testing (15%)",
			},
			Resources: []string{
				"Node.js Documentation",
				"Express.js Official Guide",
				"MongoDB University courses",
				"Postman - API testing and documentation",
				"Heroku/Netlify - Deployment platforms",
			},
		},
		"section4-react": {
			CourseCode:    "CIS 244",
			Credits:       "5.0 Credits",
			Prerequisites: "CIS 241 with a minimum grade of 2.5 or instructor permission",
			Description:   "Students learn to work with open-source JavaScript frameworks including React, AngularJS, Vue.js, and other commonly used frameworks to create and update website content.",
			Objectives: []string{
				"Determine business model of websites (B2B, B2C, e-commerce, social networking)",
				"Compare and contrast top JavaScript frameworks",
				"Develop and implement content using JavaScript frameworks",
				"Develop responsive and accessible websites using current technologies",
				"Create ongoing plan to maintain and update websites",
			},
			Topics: []string{
				"React fundamentals: components, JSX, props",
				"State management with hooks (useState, useEffect, useContext)",
				"Component lifecycle and side effects",
				"Event handling and forms in React",
				"React Router for single-page applications",
				"State management with Redux or Context API",
				"API integration and data fetching",
				"Testing React components",
				"Performance optimization techniques",
				"Deployment and build optimization",
				"Modern React patterns and best practices",
				"Capstone project: Production-ready React application",
			},
			Assessment: []string{
				"Component-building exercises (35%)",
				"Mid-term project: Multi-page React app (25%)",
				"Capstone project: Full-featured application (30%)",
				"Code quality and documentation (10%)",
			},
			Resources: []string{
				"React Official Documentation",
				"Create React App - Development environment",
				"React Router documentation",
				"Redux Toolkit - State management",
				"Vercel/Netlify - React deployment platforms",
			},
		},
	}

	info, exists := sectionInfo[sectionID]
	if !exists {
		info = sectionInfo["section1-html-css"] // Default fallback
	}

	// Combine section data with syllabus info
	response := struct {
		*Section
		SyllabusInfo interface{} `json:"syllabus_info"`
	}{
		Section:      section,
		SyllabusInfo: info,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleSyllabus(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	syllabus := struct {
		Course         Course              `json:"course"`
		InstructorInfo InstructorInfo      `json:"instructor_info"` // Add this
		Lessons        map[int]*Lesson     `json:"lessons"`
		Sections       map[string]*Section `json:"sections"`
		Weeks          []int               `json:"weeks"`
		LastUpdated    time.Time           `json:"last_updated"`
		TotalFiles     int                 `json:"total_files"`
	}{
		Course:         s.course,
		InstructorInfo: getInstructorInfo(), // Add this
		Lessons:        s.lessons,
		Sections:       s.sections,
		Weeks:          make([]int, 0),
		LastUpdated:    time.Now(),
		TotalFiles:     len(s.lessons),
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

// Keep your existing handleStatic function as-is
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
	log.Printf("Requested path: %s", r.URL.Path)
	log.Printf("Cleaned path: %s", path)
	log.Printf("Embed path: %s", embedPath)

	// Try to read the file from embedded filesystem
	data, err := staticFiles.ReadFile(embedPath)
	if err != nil {
		log.Printf("File not found: %s, serving index.html fallback", embedPath)
		indexPath := embedPath + "/index.html"
		log.Printf("Trying index path: %s", indexPath)

		// If file not found, serve index.html (for SPA routing)
		data, err = staticFiles.ReadFile(indexPath)
		if err != nil {
			log.Printf("Index file not found: %s, serving root index.html fallback", indexPath)

			// If still not found, serve root index.html (for SPA routing)
			data, err = staticFiles.ReadFile("lessons/frontend/dist/index.html")
			if err != nil {
				http.NotFound(w, r)
				return
			}
		} else {
			log.Printf("Index file found: %s", indexPath)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
		return
	}

	// Set appropriate content type
	if strings.HasSuffix(path, ".css") {
		log.Printf("Setting CSS content type")
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

	// Create API subrouter FIRST
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")
	api.HandleFunc("/sections", s.handleSections).Methods("GET")
	api.HandleFunc("/sections/{section}", s.handleSection).Methods("GET")
	// Add this route in your setupRoutes function
	api.HandleFunc("/sections/{section}/syllabus", s.handleSectionSyllabus).Methods("GET")

	// ###### This would allow both URL patterns:
	//       /api/sections/section1-html-css/week/5 (original)
	// ##### /api/sections/section1-html-css/5 (shorter)

	// api.HandleFunc("/sections/{section}/{week:[0-9]+}", s.handleSectionLesson).Methods("GET")
	api.HandleFunc("/sections/{section}/week/{week:[0-9]+}", s.handleSectionLesson).Methods("GET")

	api.HandleFunc("/sections/{section}/week/{week:[0-9]+}/toc", s.handleLessonTOC).Methods("GET")
	api.HandleFunc("/sections/{section}/week/{week:[0-9]+}/content", s.handleLessonContent).Methods("GET")

	// Debug log
	log.Println("API routes registered")

	// Static files handler - MUST be after API routes
	r.PathPrefix("/").HandlerFunc(s.handleStatic)

	// CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)

	return corsHandler
}
func (s *Server) setupRoutesLegasy2() http.Handler {
	r := mux.NewRouter()

	// API routes MUST be first
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")
	api.HandleFunc("/sections", s.handleSections).Methods("GET")
	api.HandleFunc("/sections/{section}", s.handleSection).Methods("GET")
	api.HandleFunc("/sections/{section}/week/{week:[0-9]+}", s.handleSectionLesson).Methods("GET")

	// DEBUG: Print registered routes
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Printf("🔗 Registered route: %s", pathTemplate)
		}
		return nil
	})

	// Static handler MUST be last
	// r.PathPrefix("/").HandlerFunc(s.handleStatic)

	// Static files - use a more specific pattern that excludes /api
	r.PathPrefix("/static/").HandlerFunc(s.handleStatic)
	// Root
	r.PathPrefix("/").HandlerFunc(s.handleStatic) // Everything except /api

	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)
}

func (s *Server) setupRoutesLegacy() http.Handler {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Legacy routes (maintain backward compatibility)
	api.HandleFunc("/course", s.handleCourse).Methods("GET")
	api.HandleFunc("/lessons", s.handleLessons).Methods("GET")
	api.HandleFunc("/lessons/{week:[0-9]+}", s.handleLesson).Methods("GET")
	api.HandleFunc("/syllabus", s.handleSyllabus).Methods("GET")

	// New section-based routes
	api.HandleFunc("/sections", s.handleSections).Methods("GET")
	api.HandleFunc("/sections/{section}", s.handleSection).Methods("GET")
	api.HandleFunc("/sections/{section}/week/{week:[0-9]+}", s.handleSectionLesson).Methods("GET")

	// Serve static files and SPA routes
	// r.PathPrefix("/").HandlerFunc(s.handleStatic)
	r.PathPrefix("/").Handler(http.HandlerFunc(s.handleStatic))

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(r)

	return corsHandler
}

// Add these methods to your Server struct (add after your existing methods)

func (s *Server) handleLessonTOC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["section"]
	weekStr := vars["week"]

	week, err := strconv.Atoi(weekStr)
	if err != nil || week < 1 {
		http.Error(w, "Invalid week number", http.StatusBadRequest)
		return
	}

	// Get the lesson content
	s.mutex.RLock()
	section, sectionExists := s.sections[sectionID]
	s.mutex.RUnlock()

	if !sectionExists {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	// Convert section week to global week
	globalWeek := section.WeekStart + week - 1

	s.mutex.RLock()
	lesson, lessonExists := s.lessons[globalWeek]
	s.mutex.RUnlock()

	response := TOCResponse{
		Week:    week,
		Section: sectionID,
		Source:  "error",
	}

	if !lessonExists {
		response.TOCItems = s.getDefaultTOCItems()
		response.Source = "default"
	} else {
		// Extract TOC from lesson content
		tocItems := s.extractTOCFromContent(lesson.Content)
		if len(tocItems) == 0 {
			response.TOCItems = s.getDefaultTOCItems()
			response.Source = "default"
		} else {
			response.TOCItems = tocItems
			response.Source = "markdown"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleLessonContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["section"]
	weekStr := vars["week"]

	week, err := strconv.Atoi(weekStr)
	if err != nil || week < 1 {
		http.Error(w, "Invalid week number", http.StatusBadRequest)
		return
	}

	s.mutex.RLock()
	section, sectionExists := s.sections[sectionID]
	s.mutex.RUnlock()

	if !sectionExists {
		http.Error(w, "Section not found", http.StatusNotFound)
		return
	}

	// Convert section week to global week
	globalWeek := section.WeekStart + week - 1

	s.mutex.RLock()
	lesson, lessonExists := s.lessons[globalWeek]
	s.mutex.RUnlock()

	if !lessonExists {
		http.Error(w, "Lesson content not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(lesson.Content))
}

// Add these helper methods

func (s *Server) extractTOCFromContent(content string) []TOCItem {
	var tocItems []TOCItem
	lines := strings.Split(content, "\n")
	inTOCSection := false

	// Method 1: Look for explicit Table of Contents section
	for _, line := range lines {
		// Look for TOC header
		if matched, _ := regexp.MatchString(`(?i)^##?\s*(table of contents|contents|toc)\s*$`, line); matched {
			inTOCSection = true
			continue
		}

		// Stop when we hit another major section
		if inTOCSection {
			if matched, _ := regexp.MatchString(`^##?\s+[^[]+$`, line); matched {
				break
			}
		}

		// Extract TOC links from explicit TOC section
		if inTOCSection {
			if matched, _ := regexp.MatchString(`^\s*[-*]\s*\[([^\]]+)\]\(#([^)]+)\)`, line); matched {
				re := regexp.MustCompile(`^\s*[-*]\s*\[([^\]]+)\]\(#([^)]+)\)`)
				matches := re.FindStringSubmatch(line)
				if len(matches) >= 3 {
					title := matches[1]
					id := matches[2]

					// Estimate level from indentation
					indentation := len(line) - len(strings.TrimLeft(line, " \t"))
					level := (indentation / 2) + 2 // Convert to heading level
					if level > 6 {
						level = 6
					}

					tocItems = append(tocItems, TOCItem{
						ID:    id,
						Title: title,
						Level: level,
					})
				}
			}
		}
	}

	// Method 2: If no explicit TOC found, extract from headings
	if len(tocItems) == 0 {
		tocItems = s.extractTOCFromHeadings(content)
	}

	return tocItems
}

func (s *Server) extractTOCFromHeadings(content string) []TOCItem {
	var tocItems []TOCItem
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if matched, _ := regexp.MatchString(`^(#{1,6})\s+(.+)$`, line); matched {
			re := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
			matches := re.FindStringSubmatch(line)
			if len(matches) >= 3 {
				hashes := matches[1]
				title := matches[2]
				level := len(hashes)

				// Skip h1 (main title) and empty titles
				if level <= 1 || strings.TrimSpace(title) == "" {
					continue
				}

				// Generate ID from title
				id := s.generateIDFromTitle(title)

				tocItems = append(tocItems, TOCItem{
					ID:    id,
					Title: s.cleanTitle(title),
					Level: level,
				})
			}
		}
	}

	return tocItems
}

func (s *Server) generateIDFromTitle(title string) string {
	// Remove markdown formatting
	re1 := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	title = re1.ReplaceAllString(title, "$1") // Bold
	re2 := regexp.MustCompile(`\*([^*]+)\*`)
	title = re2.ReplaceAllString(title, "$1") // Italic
	re3 := regexp.MustCompile("`([^`]+)`")
	title = re3.ReplaceAllString(title, "$1") // Code
	re4 := regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)
	title = re4.ReplaceAllString(title, "$1") // Links

	// Convert to lowercase and replace spaces/special chars with hyphens
	id := strings.ToLower(title)
	re5 := regexp.MustCompile(`[^\w\s-]`)
	id = re5.ReplaceAllString(id, "")
	re6 := regexp.MustCompile(`\s+`)
	id = re6.ReplaceAllString(id, "-")
	re7 := regexp.MustCompile(`-+`)
	id = re7.ReplaceAllString(id, "-")
	id = strings.Trim(id, "-")

	return id
}

func (s *Server) cleanTitle(title string) string {
	// Remove markdown formatting for display
	re1 := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	title = re1.ReplaceAllString(title, "$1") // Bold
	re2 := regexp.MustCompile(`\*([^*]+)\*`)
	title = re2.ReplaceAllString(title, "$1") // Italic
	re3 := regexp.MustCompile("`([^`]+)`")
	title = re3.ReplaceAllString(title, "$1") // Code
	re4 := regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)
	title = re4.ReplaceAllString(title, "$1") // Links

	return strings.TrimSpace(title)
}

func (s *Server) getDefaultTOCItems() []TOCItem {
	return []TOCItem{
		{ID: "learning-objectives", Title: "Learning Objectives", Level: 2},
		{ID: "introduction", Title: "Introduction", Level: 2},
		{ID: "main-concepts", Title: "Main Concepts", Level: 2},
		{ID: "practical-examples", Title: "Practical Examples", Level: 2},
		{ID: "hands-on-practice", Title: "Hands-on Practice", Level: 2},
		{ID: "review-summary", Title: "Review & Summary", Level: 2},
		{ID: "assignments", Title: "Assignments", Level: 2},
		{ID: "resources", Title: "Additional Resources", Level: 2},
	}
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

	log.Println("Checking embedded files:")
	err := fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			log.Printf("Embedded file: %s", path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking embedded files: %v", err)
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

	log.Printf("Course Management System Server")
	log.Printf("Lessons directory: %s", lessonsDir)
	log.Printf("Found %d lessons in %d sections", len(server.lessons), len(server.sections))
	log.Printf("Server starting on http://localhost%s", port)
	log.Printf("Frontend: http://localhost%s", port)
	log.Printf("API: http://localhost%s/api", port)
	log.Printf("New Section API: http://localhost%s/api/sections", port)

	handler := server.setupRoutes()
	log.Fatal(http.ListenAndServe(port, handler))
}
