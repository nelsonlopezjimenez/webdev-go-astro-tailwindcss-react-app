---
title: "Format svg sanitize or allow"
description: "svg tag"
publishDate: 2025-09-16
tags: ["tutorial", "guide", "test"]
draft: false
---
SVG tags not rendering in markdown is a common issue. The problem is likely that your markdown processor (marked.js) is either stripping the SVG tags for security reasons or not properly handling them.

Here are solutions to fix SVG rendering:

## Solution 1: Configure marked.js to allow SVG

Update your MarkdownRenderer to allow HTML including SVG:

```javascript
marked.setOptions({
  breaks: true,
  gfm: true,
  sanitize: false, // Allow HTML tags including SVG
  highlight: function(code, lang) {
    // your existing highlight code
  }
});
```

## Solution 2: Use DOMPurify for safe SVG rendering

Install DOMPurify:
```bash
npm install dompurify
```

Then update your MarkdownRenderer:

```javascript
import DOMPurify from 'dompurify';

const MarkdownRenderer = ({ content }) => {
  useEffect(() => {
    if (content) {
      const html = marked(content);
      // Sanitize but allow SVG
      const cleanHtml = DOMPurify.sanitize(html, {
        ADD_TAGS: ['svg', 'path', 'circle', 'rect', 'line', 'polygon', 'polyline', 'g', 'text'],
        ADD_ATTR: ['viewBox', 'xmlns', 'd', 'fill', 'stroke', 'stroke-width', 'x', 'y', 'width', 'height', 'cx', 'cy', 'r']
      });
      setHtmlContent(cleanHtml);
    }
  }, [content]);
};
```

## Solution 3: Serve SVG as image files

Instead of inline SVG, save your SVG as a file and reference it:

```markdown
<img src="/images/diagram.svg" alt="SVG diagram" width="400" height="300">
```

## Solution 4: Check browser console

The SVG might be there but not visible due to CSS issues. Check if:
- The SVG has proper `viewBox` attribute
- CSS isn't hiding it with `display: none`
- The SVG has explicit width/height

Try this test SVG in your markdown:

```markdown
<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">
  <circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" />
</svg>
```

The most likely fix is Solution 1 - setting `sanitize: false` in marked.js options to allow HTML/SVG tags to render.