---
title: "Vibrant vivid code block colors"
description: "Vibrant vivid colors"
publishDate: 2025-09-16
tags: ["tutorial", "guide", "test"]
draft: false
---
You're right - the first image (Astro's rendering) has much more vibrant colors compared to the second (Prism.js). The difference is that Astro uses the "github-dark" theme which has brighter, more saturated colors.

**Option 1: Use a more vibrant Prism theme**

Replace the theme import in your MarkdownRenderer:

```javascript
// Remove this:
// import 'prismjs/themes/prism-dark.css';

// Use one of these instead:
import 'prismjs/themes/prism-tomorrow.css';
// or
import 'prismjs/themes/prism-okaidia.css';
// or
import 'prismjs/themes/prism-solarizedlight.css';
```

**Option 2: Create custom CSS to match Astro's colors**

Add this CSS to override Prism's colors:

```css
/* Add to your component or global CSS */
.markdown-content pre[class*="language-"] {
  background: #24292e !important;
  color: #e1e4e8 !important;
}

.markdown-content .token.comment {
  color: #6a737d !important;
}

.markdown-content .token.keyword {
  color: #f97583 !important;
}

.markdown-content .token.string {
  color: #9ecbff !important;
}

.markdown-content .token.function {
  color: #b392f0 !important;
}

.markdown-content .token.number {
  color: #79b8ff !important;
}

.markdown-content .token.operator {
  color: #f97583 !important;
}

.markdown-content .token.punctuation {
  color: #e1e4e8 !important;
}
```

**Option 3: Switch to Shiki (same as Astro uses)**

Replace Prism with Shiki for exact color matching:

```bash
npm install shiki
```

```javascript
import { getHighlighter } from 'shiki';

const MarkdownRenderer = ({ content }) => {
  const [htmlContent, setHtmlContent] = useState('');

  useEffect(() => {
    const highlightCode = async () => {
      const highlighter = await getHighlighter({
        theme: 'github-dark',
        langs: ['typescript', 'javascript', 'yaml', 'bash']
      });

      marked.setOptions({
        highlight: (code, lang) => {
          if (lang && highlighter.getLoadedLanguages().includes(lang)) {
            return highlighter.codeToHtml(code, { lang, theme: 'github-dark' });
          }
          return code;
        }
      });

      setHtmlContent(marked(content));
    };

    if (content) {
      highlightCode();
    }
  }, [content]);
};
```

Option 3 with Shiki will give you exactly the same colors as Astro since they use the same highlighting engine and theme.


Here's a practical comparison of Shiki vs Prism.js for your use case:

## Shiki

**Pros:**
- **Exact color matching**: Uses the same themes as VS Code and Astro
- **Superior accuracy**: TextMate grammars provide more precise syntax parsing
- **Theme consistency**: Wide selection of authentic VS Code themes
- **Future-proof**: Actively maintained by the VS Code team
- **Better language support**: More comprehensive language definitions

**Cons:**
- **Larger bundle size**: Includes WASM runtime and language grammars
- **Async initialization**: Requires await for setup, complicates component logic
- **Performance overhead**: Slower initial load due to WASM compilation
- **Complexity**: More setup code required

## Prism.js

**Pros:**
- **Smaller bundle**: Lightweight, only includes what you import
- **Synchronous**: No async setup required, simpler implementation
- **Fast rendering**: Pure JavaScript, no WASM overhead
- **Mature ecosystem**: Extensive plugin system and themes
- **Simple integration**: Drop-in replacement with minimal code changes

**Cons:**
- **Limited themes**: Fewer high-quality themes available
- **Less accurate parsing**: Regex-based, can miss edge cases
- **Manual theme matching**: Requires custom CSS to match specific color schemes
- **Maintenance concerns**: Less active development compared to Shiki

## Recommendation for Your Project

Given that you want to match Astro's exact colors and you're already using Astro (which uses Shiki), I'd recommend **Shiki**. The bundle size increase is minimal for an educational platform, and you'll get perfect visual consistency with your lesson pages.

However, if bundle size and simplicity are priorities, stick with **Prism.js** and use the custom CSS I provided earlier to match the colors - the visual difference will be negligible to users.

For your content management system, the consistency argument favors Shiki since you're already paying the cost in your Astro build.