Here are several ways to exclude folders when using the `tree` command:

## **Using `-I` (Ignore Pattern) - Most Common**

```bash
# Exclude single folder
tree -I "node_modules"

# Exclude multiple folders (pipe-separated)
tree -I "node_modules|.git|dist"

# Exclude folders with wildcards
tree -I "*.git|*node_modules*|build*"
```

## **Common Development Exclusions**

```bash
# Web development project
tree -I "node_modules|.git|dist|build|.next|coverage"

# Go project
tree -I ".git|vendor|bin|*.exe|*.so"

# Your course project
tree -I ".git|node_modules|dist|.astro|coverage|*.log"
```

## **Using `-a` and `--dirsfirst` with `-I`**

```bash
# Show hidden files but exclude certain folders
tree -a -I ".git|node_modules"

# Show directories first, exclude build artifacts
tree --dirsfirst -I "dist|build|node_modules"
```

## **Advanced Patterns**

```bash
# Exclude by file extension and folders
tree -I "*.pyc|__pycache__|.git|*.log"

# Exclude temporary and cache folders
tree -I "*cache*|*tmp*|*temp*|.git"

# Case-insensitive exclusion (if supported)
tree -I "NODE_MODULES|node_modules|Node_Modules"
```

## **For Your Course Project Specifically**

```bash
# Basic exclusion for your Go/Astro project
tree -I "node_modules|.git|dist"

# More comprehensive exclusion
tree -I "node_modules|.git|dist|.astro|coverage|*.log|bin|vendor"

# Show only lesson structure
tree lessons/ -I ".git|*.backup"
```

## **Alternative: Using `find` with `tree`-like Output**

If your `tree` command doesn't support `-I`:

```bash
# Using find to mimic tree behavior
find . -type d \( -name "node_modules" -o -name ".git" -o -name "dist" \) -prune -o -print | sed 's/[^\/]*\//|  /g; s/|  \([^|]\)/└─ \1/'
```

## **Platform-Specific Notes**

### **Linux/Mac (GNU tree)**
```bash
tree -I "pattern"  # Works on most systems
```

### **Windows**
```cmd
# If using Windows tree command (limited options)
tree /F | findstr /V "node_modules"

# Better: Install Unix-like tree
# Using Git Bash, WSL, or install via package manager
```

### **If `-I` Option Doesn't Exist**
```bash
# Use find + grep to exclude
tree | grep -v "node_modules\|\.git\|dist"
```

## **Example Output**

```bash
$ tree -I "node_modules|.git|dist"
.
├── go.mod
├── go.sum
├── main.go
├── lessons/
│   ├── section1-html-css/
│   │   ├── week01-html-basics.md
│   │   └── week02-css-fundamentals.md
│   └── section2-javascript/
│       └── week13-js-introduction.md
├── src/
│   ├── components/
│   └── layouts/
└── astro.config.mjs
```

The `-I` option with pipe-separated patterns is the most reliable and widely supported method across different systems.