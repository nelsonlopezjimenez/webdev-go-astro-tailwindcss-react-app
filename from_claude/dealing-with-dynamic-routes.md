The error is because Astro requires `getStaticPaths()` for dynamic routes like `[week].astro`. However, since we're fetching data at runtime from the Go API, we should use server-side rendering instead. Let me fix this by updating the Astro config and the dynamic route.Now I need to add the Node.js adapter for server-side rendering:Actually, let me provide a simpler solution. Instead of switching to server-side rendering, let's keep it static and provide the required `getStaticPaths()` function:Now let me add the required `getStaticPaths()` function to the dynamic route:Perfect! Now I've fixed the issue by:

## ✅ **Fixed the Dynamic Route Issue:**

1. **Added `getStaticPaths()` function** - This tells Astro which dynamic routes to pre-generate
2. **Pre-defined all 10 week paths** - The function returns paths for weeks 1-10
3. **Kept static site generation** - No need for server-side rendering complexity

## 🔧 **How it works:**

- **Build time**: Astro generates static pages for `/lessons/1`, `/lessons/2`, etc.
- **Runtime**: Each page fetches lesson data from your Go API
- **Fallback**: Pages gracefully handle missing lessons with error messages

## 🚀 **Next Steps:**

1. **Install dependencies:**
   ```bash
   cd frontend
   npm install
   ```

2. **Make sure your Go server is running:**
   ```bash
   go run main.go
   ```

3. **Start the Astro dev server:**
   ```bash
   npm run dev
   ```

Now you should be able to access:
- `http://localhost:4321/lessons/1` - Week 1 lesson
- `http://localhost:4321/lessons/2` - Week 2 lesson
- etc.

The dynamic route will now work properly and fetch lesson content from your Go backend API!