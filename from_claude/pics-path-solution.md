The issue is that your image file isn't being served by your Go server. The markdown is processed correctly, but the image path isn't accessible through your current static file configuration.

Here are several solutions:

**Option 1: Put images in the Astro public directory**

Move your image to `lessons/frontend/public/images/`:
```
lessons/frontend/public/
  └── images/
      └── image-1.png
```

Then use this path in your markdown:
```markdown
![development branch](/images/image-1.png)
<img src="/images/image-1.png" alt="development branch" width="500" height="600">
```

**Option 2: Add image serving to your Go server**

Add this to your `main.go` to serve images from lesson directories:

```go
// Add this route in your setupRoutes function
api.HandleFunc("/images/{section}/{week}/{filename}", s.handleLessonImages).Methods("GET")

// Add this method to your Server struct
func (s *Server) handleLessonImages(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    section := vars["section"]
    week := vars["week"]
    filename := vars["filename"]
    
    // Construct the file path
    imagePath := filepath.Join(s.lessonsDir, section, filename)
    
    // Check if file exists
    if _, err := os.Stat(imagePath); os.IsNotExist(err) {
        http.NotFound(w, r)
        return
    }
    
    // Serve the image
    http.ServeFile(w, r, imagePath)
}
```

Then use this path in your markdown:
```markdown
![development branch](/api/images/section1-html-css/week1/image-1.png)
```

**Option 3: Update your static handler to serve lesson images**

Modify your `handleStatic` function to also check lesson directories:

```go
func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/")
    
    if path == "" {
        path = "index.html"
    }
    
    // Check if it's an image request from lessons
    if strings.Contains(path, "image-") && (strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".gif")) {
        // Try to find the image in lesson directories
        lessonDirs := []string{"section1-html-css", "section2-javascript", "section3-backend", "section4-react"}
        for _, dir := range lessonDirs {
            imagePath := filepath.Join(s.lessonsDir, dir, path)
            if _, err := os.Stat(imagePath); err == nil {
                http.ServeFile(w, r, imagePath)
                return
            }
        }
    }
    
    // Your existing static file serving code...
    embedPath := "lessons/frontend/dist/" + path
    // ... rest of your existing handleStatic code
}
```

**Option 4: Use base64 encoded images**

Convert your image to base64 and embed it directly:
```markdown
<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..." alt="development branch" width="500" height="600">
```

**Recommended approach:** Use **Option 1** (put images in Astro's public directory) as it's the simplest and most reliable. The public directory is specifically designed for static assets and will be properly served by both your Go server and Astro dev server.