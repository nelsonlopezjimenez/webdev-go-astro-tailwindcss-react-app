# Course Management System

A full-stack course management application built with **Astro** (frontend) and **Go** (backend) that serves a 10-week course with dynamic lesson loading from the filesystem.

## Features

- 📚 **Dynamic Lesson Loading**: Automatically detects and serves new markdown lessons
- 🗓️ **10-Week Course Structure**: Organized weekly lesson progression  
- 📖 **Markdown Support**: Write lessons in markdown with YAML frontmatter
- 🎯 **Course Syllabus**: Complete course overview and progress tracking
- 📱 **Responsive Design**: Works on desktop, tablet, and mobile
- 🔄 **Hot Reloading**: Add lessons without restarting the server

## Technology Stack

### Backend (Go)
- **Gorilla Mux**: HTTP routing
- **YAML**: Configuration parsing
- **Filesystem Watcher**: Dynamic content loading

### Frontend (Astro)
- **React**: Interactive components
- **Tailwind CSS**: Styling
- **Marked**: Markdown parsing

## Quick Start

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn

### 1. Setup Backend

```bash
# Create project directory
mkdir course-app
cd course-app

# Initialize Go module
go mod init course-app

# Create main.go with the provided backend code
# Create go.mod with dependencies

# Install dependencies
go mod tidy

# Create lessons directory
mkdir lessons

# Add sample lessons (week1.md, week2.md, course.yaml)
```

### 2. Setup Frontend

```bash
# Create Astro project
npm create astro@latest frontend -- --template minimal --typescript false

cd frontend

# Install dependencies from package.json
npm install

# Create the Astro project structure with provided files
```

### 3. Run the Application

#### Terminal 1 - Start Backend
```bash
# From project root
go run main.go

# Server starts on http://localhost:8080
# API endpoints:
# - GET /api/course
# - GET /api/lessons  
# - GET /api/lessons/{week}
# - GET /api/syllabus
```

#### Terminal 2 - Start Frontend
```bash
# From frontend directory
cd frontend
npm run dev

# Frontend starts on http://localhost:4321
```

### 4. Access the Application
- **Frontend**: http://localhost:4321
- **Backend API**: http://localhost:8080/api

## Project Structure

```
course-app/
├── main.go              # Go backend server
├── go.mod               # Go dependencies
├── lessons/             # Markdown lesson files
│   ├── course.yaml      # Course configuration
│   ├── week1.md         # Week 1 lesson
│   ├── week2.md         # Week 2 lesson
│   └── frontend/            # Astro frontend
    ├── src/
    │   ├── layouts/
    │   │   └── Layout.astro
    │   ├── pages/
    │   │   ├── index.astro
    │   │   ├── syllabus.astro
    │   │   └── lessons/
    │   │       ├── index.astro
    │   │       └── [week].astro
    │   └── components/
    │       └── MarkdownRenderer.jsx
    ├── astro.config.mjs
    └── package.json
```

## Adding New Lessons

### Method 1: With YAML Frontmatter (Recommended)

Create a new `.md` file in the `lessons/` directory:

```markdown
---
title: "Variables and Data Types"
description: "Learn about different data types and how to use variables effectively"
week: 3
---

# Your Lesson Content

Write your lesson content here in markdown format.

## Code Examples

```python
# Python code examples
name = "Alice"
age = 25
print(f"Hello {name}, you are {age} years old")
```

## Exercises

1. Create a variable for your name
2. Print a greeting message
```

### Method 2: Filename-based (Simple)

Name your file using week patterns:
- `week3.md`
- `week-4.md` 
- `lesson_week_5.md`
- `03-advanced-topics.md`

The system will automatically detect the week number from the filename.

### Method 3: Course Configuration

Edit `lessons/course.yaml` to update course information:

```yaml
title: "Your Course Title"
description: "Course description that appears on the homepage"
duration: "10 weeks"
instructor: "Your Name"
```

## API Endpoints

### GET /api/course
Returns course information from `course.yaml`

```json
{
  "title": "Programming Fundamentals",
  "description": "A comprehensive 10-week course...",
  "duration": "10 weeks", 
  "instructor": "Dr. Sarah Johnson"
}
```

### GET /api/lessons
Returns all available lessons

```json
[
  {
    "week": 1,
    "title": "Introduction to Programming",
    "description": "Get started with programming concepts...",
    "content": "# Welcome to Programming...",
    "created_at": "2024-01-15T10:30:00Z",
    "file_path": "./lessons/week1.md"
  }
]
```

### GET /api/lessons/{week}
Returns specific lesson by week number (1-10)

### GET /api/syllabus  
Returns complete course syllabus with all lessons and course info

## Customization

### Styling
The frontend uses Tailwind CSS. Modify styles in:
- `src/layouts/Layout.astro` - Global styles and layout
- Individual page components for specific styling

### Adding Features
- **Authentication**: Add user login/registration
- **Progress Tracking**: Store user progress in database  
- **Quizzes**: Add interactive quizzes to lessons
- **File Uploads**: Allow students to submit assignments
- **Comments**: Add discussion features to lessons

### Database Integration
To persist data, you can integrate databases:

**Go Backend Options:**
- SQLite (embedded)
- PostgreSQL 
- MySQL
- MongoDB

**Example with SQLite:**
```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// Add to your Go backend for persistent storage
```

## Deployment

### Backend Deployment
```bash
# Build for production
go build -o course-server main.go

# Run in production
./course-server /path/to/lessons
```

### Frontend Deployment
```bash
# Build for production
npm run build

# The dist/ folder contains static files for deployment
# Deploy to Netlify, Vercel, or any static hosting service
```

### Docker Deployment

**Dockerfile for Backend:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o course-server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/course-server .
COPY --from=builder /app/lessons ./lessons
EXPOSE 8080
CMD ["./course-server", "./lessons"]
```

**Dockerfile for Frontend:**
```dockerfile
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## Troubleshooting

### Common Issues

**1. Lessons not appearing:**
- Check that `.md` files are in the `lessons/` directory
- Verify week numbers are 1-10
- Check YAML frontmatter syntax
- Restart the Go server to refresh filesystem scan

**2. CORS errors:**
- Ensure the Go server includes CORS headers
- Check that API_BASE URL is correct in Astro pages
- Verify both servers are running

**3. Markdown not rendering:**
- Check that the MarkdownRenderer component is imported
- Verify `marked` package is installed
- Check for malformed markdown syntax

**4. Styling issues:**
- Ensure Tailwind CSS is properly configured
- Check that styles are applied in the correct scope
- Verify class names are valid Tailwind utilities

### Debug Mode

Enable debug logging in Go:
```go
log.Printf("Found %d lessons", len(server.lessons))
log.Printf("Serving lesson for week %d", week)
```

View browser console for frontend debugging.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

- Create an issue for bug reports
- Check existing issues for common problems
- Review documentation for setup help

---

**Happy Learning!** 🎓

This course management system grows with your content - just add markdown files and watch your course come to life! ...              # Additional weeks
└──