Now here's how to use the dynamic React component in your Astro page:## Summary: Complete TOC Fetching Solution

Now you have **three different approaches** to fetch TOC items:

### **1. Server-Side Fetching (Astro Build Time)**
- **File**: First artifact - fetches TOC during Astro build
- **Pros**: Fast loading, SEO-friendly
- **Cons**: Requires rebuild when markdown changes

### **2. API Endpoint (Go Server)**
- **File**: Second artifact - Go API that parses markdown files
- **Pros**: Dynamic, works with both dev/prod modes
- **Cons**: Requires API call

### **3. Client-Side React Component**
- **File**: Third & Fourth artifacts - Dynamic React component
- **Pros**: Real-time updates, loading states, retry functionality
- **Cons**: Requires JavaScript, slight delay

## Implementation Steps:

### **1. Add the Go API endpoint:**
```go
// Add to your main.go setupRoutes()
http.HandleFunc("/api/lessons/", handleLessonAPI)
```

### **2. Choose your approach:**

**Option A - Static (Recommended for production):**
Use the first artifact for build-time TOC extraction.

**Option B - Dynamic (Recommended for development):**
Use the React component (artifacts 3 & 4) for real-time TOC fetching.

### **3. Test the API:**
```bash
# Test your API endpoint
curl http://localhost:8080/api/lessons/section1-html-css/week1/toc
```

### **4. TOC Detection Methods:**

The Go server will look for TOC in this order:
1. **Explicit TOC section** in markdown:
   ```markdown
   ## Table of Contents
   - [Introduction](#introduction)
   - [Main Topic](#main-topic)
   ```

2. **Auto-generated from headings** if no explicit TOC found:
   ```markdown
   ## Introduction
   ## Main Topic  
   ## Summary
   ```

**Which approach would you like to implement first?** I recommend starting with the React component for immediate testing, then switching to server-side for production.