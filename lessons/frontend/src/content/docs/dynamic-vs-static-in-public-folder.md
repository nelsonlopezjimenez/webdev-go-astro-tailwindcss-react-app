---
title: "Dynamic vs static generated content"
description: "In the long term what is better a dynamically generated content built into dist folder or static content in the public folder or any other folder"
publishDate: 2025-01-01
tags: ["tutorial", "guide", "test"]
draft: false
---
For a long-term educational platform, **dynamically generated content built into the dist folder** is the better choice. Here's the analysis:

## Dynamically Generated (Content Collections) - Best Choice

**Advantages:**
- **Performance**: Static HTML generation means faster page loads
- **SEO**: Search engines can crawl and index all content
- **Reliability**: No server dependency for content delivery
- **Scalability**: CDN-friendly, can handle high traffic
- **Version Control**: Content changes are tracked with code
- **Type Safety**: Schema validation prevents content errors
- **Build-time Processing**: Markdown gets optimized, images processed, links validated

**Disadvantages:**
- Requires rebuild for content changes
- Larger build times with extensive content

## Static Content in Public Folder

**Advantages:**
- Simple to implement
- Direct file access

**Disadvantages:**
- **No processing**: Raw markdown files, no HTML conversion
- **Poor SEO**: Search engines can't properly index markdown
- **No validation**: Broken links or malformed content go undetected
- **Limited functionality**: No dynamic features, table of contents, or cross-references
- **Poor user experience**: Users need markdown viewers

## Server-Side Dynamic Content

**Advantages:**
- Real-time updates
- No build step for content changes

**Disadvantages:**
- **Server dependency**: Goes down if server fails
- **Performance**: Slower than static files
- **Complexity**: Requires caching strategies
- **Cost**: Higher hosting costs for dynamic servers

## Recommendation for Educational Content

Use **Content Collections** because:

1. **Student Experience**: Fast-loading, properly formatted lessons
2. **Instructor Workflow**: Content validation catches errors early
3. **Platform Reliability**: Works even if your Go server has issues
4. **Future-Proofing**: Easy to add features like search, cross-references, progress tracking
5. **Deployment Flexibility**: Can deploy to any static hosting service

The rebuild requirement is manageable with automation - you can set up CI/CD to automatically rebuild when instructors update content, giving you both the benefits of static generation and reasonable content update workflows.