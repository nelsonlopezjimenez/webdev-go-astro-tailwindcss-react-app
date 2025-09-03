#!/bin/bash

echo "🏗️  Building Course Management System..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js first."
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go first."
    exit 1
fi

print_status "Checking prerequisites..."
print_success "Node.js version: $(node --version)"
print_success "Go version: $(go version | cut -d' ' -f3)"

# Build Frontend
print_status "Building Astro frontend..."
cd frontend

# Install dependencies if node_modules doesn't exist
if [ ! -d "node_modules" ]; then
    print_status "Installing frontend dependencies..."
    npm install
    if [ $? -ne 0 ]; then
        print_error "Failed to install frontend dependencies"
        exit 1
    fi
fi

# Build the frontend
print_status "Building frontend for production..."
npm run build
if [ $? -ne 0 ]; then
    print_error "Frontend build failed"
    exit 1
fi

print_success "Frontend built successfully"
cd ..

# Build Go Backend with Embedded Frontend
print_status "Building Go backend with embedded frontend..."

# Install Go dependencies
print_status "Installing Go dependencies..."
go mod tidy
if [ $? -ne 0 ]; then
    print_error "Failed to install Go dependencies"
    exit 1
fi

# Build for current platform
print_status "Building executable for current platform..."
go build -ldflags="-s -w" -o course-server main-embedded.go
if [ $? -ne 0 ]; then
    print_error "Go build failed"
    exit 1
fi

print_success "Built: course-server"

# Build for different platforms
print_status "Building for multiple platforms..."

# Windows
print_status "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o course-server-windows.exe main-embedded.go
if [ $? -eq 0 ]; then
    print_success "Built: course-server-windows.exe"
else
    print_warning "Failed to build for Windows"
fi

# macOS
print_status "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o course-server-macos main-embedded.go
if [ $? -eq 0 ]; then
    print_success "Built: course-server-macos"
else
    print_warning "Failed to build for macOS"
fi

# Linux
print_status "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o course-server-linux main-embedded.go
if [ $? -eq 0 ]; then
    print_success "Built: course-server-linux"
else
    print_warning "Failed to build for Linux"
fi

# Create sample lessons if they don't exist
print_status "Setting up sample content..."

if [ ! -d "lessons" ]; then
    mkdir lessons
    print_status "Created lessons directory"
fi

if [ ! -f "lessons/course.yaml" ]; then
    cat > lessons/course.yaml << EOF
title: "Web Application Developer Certificate"
description: "Students will develop effective websites using HTML, client-side scripting, and server-side scripting. Specific emphasis is placed on developing interactive web pages that are used to process data from the Internet or intranets."
duration: "10 weeks"
instructor: "Dr. Nelson Lopez"
requirements:
  - "Build and maintain websites."
  - "Work with stakeholders to create websites."
  - "Research, assess, and appropriately apply emerging technology to support websites as needed in industry."
  - "Comply with the ethics related to the use of copyrighted materials and intellectual property rights."
  - "Demonstrate an entrepreneurial approach to web development sites and pages."
  - "Manage career goals through creating effective resumes/CVs, developing interviewing skills, and setting goals."
EOF
    print_success "Created sample course.yaml"
fi

if [ ! -f "lessons/week1.md" ]; then
    cat > lessons/week1.md << 'EOF'
---
title: "Introduction to Web Development"
description: "Get started with web development fundamentals and set up your development environment."
week: 1
---

# Welcome to Web Development!

## Learning Objectives

By the end of this lesson, you will be able to:

- Understand the basics of web development
- Set up your development environment
- Create your first HTML page
- Understand the role of HTML, CSS, and JavaScript

## What is Web Development?

Web development is the process of creating websites and web applications...

## Getting Started

Let's start with a simple HTML page:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My First Web Page</title>
</head>
<body>
    <h1>Hello, World!</h1>
    <p>Welcome to web development!</p>
</body>
</html>
```

## Next Steps

In the next lesson, we'll dive deeper into HTML structure and semantics.
EOF
    print_success "Created sample week1.md"
fi

# Create distribution package
print_status "Creating distribution package..."
mkdir -p dist

# Copy executables to dist
cp course-server* dist/ 2>/dev/null || true

# Create README for distribution
cat > dist/README.md << EOF
# Course Management System

## Quick Start

1. Run the executable for your platform:
   - Windows: \`course-server-windows.exe\`
   - macOS: \`course-server-macos\`
   - Linux: \`course-server-linux\`

2. Open your browser to: http://localhost:8080

3. Add lessons by creating .md files in the lessons directory

## Usage

\`\`\`bash
# Run with default settings (lessons in ./lessons, port 8080)
./course-server

# Specify lessons directory
./course-server /path/to/lessons

# Specify lessons directory and port
./course-server /path/to/lessons 3000
\`\`\`

## Adding Content

- Create \`.md\` files in the lessons directory
- Use YAML frontmatter for metadata:

\`\`\`markdown
---
title: "Lesson Title"
description: "Lesson description"
week: 1
---

# Your lesson content here...
\`\`\`

- Edit \`course.yaml\` to customize course information
EOF

print_success "Created distribution package in dist/"

# Show file sizes
print_status "Build Summary:"
ls -lh course-server* 2>/dev/null | awk '{print "  " $9 ": " $5}'

echo
print_success "🎉 Build completed successfully!"
print_status "Run './course-server' to start the server"
print_status "Visit http://localhost:8080 to view your course"