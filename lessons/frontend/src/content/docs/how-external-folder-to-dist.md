---
title: "How to add external folder to dist"
description: "How to add external folder to dist folder approaches"
publishDate: 2025-01-01
tags: ["tutorial", "guide", "test"]
draft: false
---

Here are the main approaches to include your external markdown files in the Astro dist folder:

**Approach 1: Astro Content Collections (Recommended)**
- Move your markdown files to `src/content/` directory in your Astro project
- Define a content collection schema in `src/content/config.ts`
- Astro will automatically process and include them in the build
- Benefits: Built-in frontmatter validation, type safety, automatic routing
- Drawback: Requires moving files into the Astro project structure

**Approach 2: Build-time Copy Script**
- Create a build script that copies markdown files to `src/pages/` or `public/` before Astro builds
- Use Node.js fs operations or shell commands in package.json scripts
- Files get processed as regular Astro pages or static assets
- Benefits: Keeps original file location, flexible file organization
- Drawback: Requires custom build pipeline

**Approach 3: Astro Integration/Plugin**
- Create a custom Astro integration that watches and copies files during development
- Hook into Astro's build process to include files automatically
- Benefits: Seamless integration, hot reloading support
- Drawback: More complex setup, requires Astro API knowledge

**Approach 4: Symlink/Junction Approach**
- Create symbolic links from your external folder to `src/content/` or `src/pages/`
- Astro treats them as regular project files
- Benefits: No file duplication, real-time updates
- Drawback: Platform-specific, potential Git issues

**Approach 5: Go Server Direct Serving**
- Let your Go server serve the markdown files directly without involving Astro
- Create API endpoints that return processed markdown as HTML
- Use client-side rendering to display content
- Benefits: No build-time dependency, dynamic content
- Drawback: Not included in static build, requires server

**Best Suggestion: Approach 2 (Build-time Copy Script)**
This gives you the most flexibility while keeping your files organized outside the frontend. You can maintain your current folder structure and have a script that copies/processes files before each build, ensuring they're included in the dist folder.
