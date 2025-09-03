**No, you cannot add lessons to the embedded executable without recompiling.** The embedded files are compiled into the binary at build time and become immutable.

However, your current server architecture already solves this problem perfectly:

## How Your System Works

1. **Static Frontend**: Embedded in the executable (requires recompile to change)
2. **Dynamic Lesson Content**: Read from the filesystem at runtime from the `./lessons` directory
3. **File Watcher**: Automatically detects new `.md` files and updates lesson list

## Adding New Lessons (Current Setup)

You can add lessons without recompiling by:

```bash
# Add a new lesson file
echo "---
title: New Lesson Title  
description: Lesson description
week: 12
---

# Lesson Content Here" > ./lessons/week12.md

# Server automatically detects and loads it
```

The server's file watcher will:
- Detect the new file
- Parse the metadata
- Update the lesson list
- Serve it via `/api/lessons` and `/api/lessons/12`

## What Requires Recompiling

Only changes to:
- Frontend UI (Astro components, styles, layouts)
- Server code (Go handlers, routing)
- Static assets (images, fonts embedded in frontend)

## Best of Both Worlds

Your architecture gives you:
- **Fast deployment**: Single executable with embedded frontend
- **Content flexibility**: Dynamic lesson loading from filesystem
- **Live updates**: File watcher for immediate content changes
- **No runtime dependencies**: No need for Node.js or complex setups

This is actually an ideal setup for a course management system - the UI is stable and embedded, while the content remains flexible and updatable.