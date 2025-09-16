---
title: "Slug doc body data slug"
description: "In Astro, the `Content` from `doc.render()` is a component, not a string or object with accessible content."
publishDate: 2025-09-16
tags: ["astro", "collection", "slug"]
draft: false
---
In Astro, the `Content` from `doc.render()` is a component, not a string or object with accessible content. You can't directly access its HTML content through console.log or JavaScript inspection.

Here's what you're working with:

```astro
---
const { doc } = Astro.props;
const { Content } = await doc.render();

// This won't show you the HTML content:
console.log("Content component:", Content); // Shows component function

// To debug the actual markdown content, use:
console.log("Raw markdown body:", doc.body);
console.log("Frontmatter data:", doc.data);
console.log("Document slug:", doc.slug);
---
```

**If you need to access the rendered HTML content:**

You can't directly get the HTML string from the Content component, but you can:

1. **Access the raw markdown:**
```astro
---
console.log("Raw markdown:", doc.body);
---
```

2. **Render to string (server-side only):**
```astro
---
import { experimental_AstroContainer as AstroContainer } from 'astro/container';

// This is more complex and not commonly needed
const container = await AstroContainer.create();
const result = await container.renderToString(Content);
console.log("Rendered HTML:", result);
---
```

3. **Debug the rendered output in the browser:**
```astro
<div class="prose prose-lg max-w-none" id="content-debug">
  <Content />
</div>

<script>
  // This runs client-side after rendering
  document.addEventListener('DOMContentLoaded', () => {
    const contentDiv = document.getElementById('content-debug');
    console.log("Rendered HTML:", contentDiv.innerHTML);
  });
</script>
```

The most practical approach for debugging is to inspect `doc.body` (raw markdown) and `doc.data` (frontmatter) rather than trying to access the rendered HTML.