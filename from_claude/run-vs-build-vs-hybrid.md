Great question! Yes, you can build and serve statically, but there's an important caveat with the current setup.

## 🏗️ **Build Methods:**

### **Method 1: Static Build** (Current Issue)
```bash
cd frontend
npm run build
npm run preview
```

**❌ Problem**: Static builds happen at build-time, so new `.md` files added after building won't appear until you rebuild.

### **Method 2: Development Server** (What you're using)
```bash
npm run dev
```

**✅ Works**: Fetches fresh data from Go API on each page load.

## 🔧 **Solutions for Static Building:**

### **Option A: Hybrid Build (Recommended)**

Update your `astro.config.mjs` to make only the lesson pages server-rendered:Now mark dynamic pages as server-rendered by adding `export const prerender = false;` to lesson pages:### **Option B: Build Script with Auto-rebuild**## 🚀 **Deployment Options:**

### **Option A: Hybrid Mode (Recommended)**
```bash
cd frontend
npm install @astrojs/node
npm run build
npm run preview
# Static pages + server-rendered lessons
```

### **Option B: Full Static with Auto-rebuild**
```bash
cd frontend
npm install chokidar
node build-watch.js
# Rebuilds automatically when lessons change
```

### **Option C: Simple Static (Manual)**
```bash
cd frontend
npm run build  # Build once
npm run preview  # Serve static files

# Add new lesson
echo "new lesson" > ../lessons/week8.md

# Rebuild to see changes  
npm run build
```

## 📊 **Comparison:**

| Method | New Lessons | Performance | Complexity |
|--------|-------------|-------------|------------|
| `npm run dev` | ✅ Instant | Slower | Simple |
| Hybrid Build | ✅ Instant | Fast | Medium |
| Auto-rebuild | ✅ ~1-2s delay | Fast | Medium |
| Static Build | ❌ Manual rebuild | Fastest | Simple |

## 🎯 **Recommendation:**

Use **Hybrid Mode** - it gives you the best of both worlds:
- Static pages for performance
- Server-rendered lesson pages for dynamic content
- No rebuild needed for new lessons

Try the hybrid approach first! 🚀