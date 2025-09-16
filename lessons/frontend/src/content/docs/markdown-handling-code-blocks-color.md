---
title: "Markdown handling code blocks color"
description: "The problem is that your MarkdownRenderer component is doing basic string replacement for markdown syntax, but it's not handling code blocks with language specification properly."
publishDate: 2025-09-16
tags: ["markdown", "code blocks", "setDate"]
draft: false
---

Perfect! You solved the main rendering issue by using your existing `MarkdownRenderer` component instead of Astro's `Content` component. Now the code block rendering issue is much clearer to fix.

The problem is that your `MarkdownRenderer` component is doing basic string replacement for markdown syntax, but it's not handling code blocks with language specification properly.

Looking at your code blocks, they should be:

```yaml
```

But your simple markdown converter is likely just doing basic replacements. Here's how to fix your `MarkdownRenderer` component:

**Update your `MarkdownRenderer.jsx` to handle code blocks properly:**

```javascript
// In your MarkdownRenderer component
function convertMarkdownToHTML(markdown) {
  return markdown
    // Handle fenced code blocks with language
    .replace(/```(\w+)?\n([\s\S]*?)```/g, (match, lang, code) => {
      const language = lang || '';
      const escapedCode = code
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;');
      
      return `<pre><code class="language-${language}">${escapedCode}</code></pre>`;
    })
    // Headers
    .replace(/^### (.*$)/gim, '<h3>$1</h3>')
    .replace(/^## (.*$)/gim, '<h2>$1</h2>')
    .replace(/^# (.*$)/gim, '<h1>$1</h1>')
    // Bold and italic
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.*?)\*/g, '<em>$1</em>')
    // Inline code (but not in code blocks)
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    // Links
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a href="$2">$1</a>')
    // Line breaks
    .replace(/\n\n/g, '</p><p>')
    .replace(/\n/g, '<br>');
}
```

**Or better yet, use a proper markdown library:**

Install a markdown parser:
```bash
npm install marked
```

Then update your `MarkdownRenderer`:

```javascript
import { marked } from 'marked';

export default function MarkdownRenderer({ content }) {
  const htmlContent = marked(content, {
    highlight: function(code, lang) {
      // Basic syntax highlighting
      return `<code class="language-${lang}">${code}</code>`;
    }
  });

  return (
    <div 
      className="prose prose-lg max-w-none"
      dangerouslySetInnerHTML={{ __html: htmlContent }}
    />
  );
}
```

This should properly render your code blocks with syntax highlighting and correct formatting.