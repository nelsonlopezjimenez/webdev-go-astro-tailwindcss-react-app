---
title: "Frontmatter set current date"
description: "Setting date: static, build script, template placeholder, dynamic, git-based"
publishDate: 2025-09-16
tags: ["frontmatter", "date", "setDate"]
draft: false
---

# Table of content

1. [Static](#method-1-static-current-date-manual)
1. [Automated](#method-2-using-build-tools-automated)
1. [Template placeholder](#method-3-template-placeholder)
1. [Dynamic](#method-4-astro-content-collections-dynamic)
1. [git-based](#method-5-git-based-date)

There are several ways to set a date to the current date in markdown frontmatter:

## Method 1: Static Current Date (Manual)

```yaml
---
title: "Your Document Title"
publishDate: 2025-01-17
---
```

## Method 2: Using Build Tools (Automated)

If you're using a build script or CI/CD pipeline, you can inject the current date:

```bash
# In a build script
current_date=$(date +%Y-%m-%d)
sed -i "s/publishDate: auto/publishDate: $current_date/" your-file.md
```

## Method 3: Template Placeholder

Use a placeholder that gets replaced during build:

```yaml
---
title: "Your Document Title"
publishDate: "{{CURRENT_DATE}}"
---
```

Then replace it with a script or build tool.

## Method 4: Astro Content Collections (Dynamic)

Since you're using Astro Content Collections, you can handle this in your content schema:

```typescript
// In src/content/config.ts
const docsCollection = defineCollection({
  schema: z.object({
    title: z.string(),
    publishDate: z.date().default(() => new Date()), // Defaults to current date
    // ... other fields
  }),
});
```

## Method 5: Git-based Date

Use the file's Git creation or modification date:

```typescript
// In your Astro page
import { execSync } from 'child_process';

const gitDate = execSync(`git log -1 --format=%ci ${filePath}`).toString().trim();
```

## Recommended Approach

For your use case, **Method 4** is best since you're using Astro Content Collections. You can either:

1. **Default to current date** if no date is provided
2. **Manually set dates** when creating files
3. **Use Git dates** for automatic tracking

The schema default will automatically set the current date when no `publishDate` is specified in the frontmatter.