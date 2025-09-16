---
title: "Content collection steps"
description: "Approach 1: Astro Content Collections. Move your markdown files to src/content/ directory. Define a schema in src/content/config.ts."
publishDate: 2025-01-01
tags: ["collection", "config", "test"]
draft: false
---

# Approach 1: 

## Astro Content Collections (Recommended). 

1. Move your markdown files to src/content/ directory in your Astro project.
1. Define a content collection schema in src/content/config.ts 
1. Astro will automatically process and include them in the build 1. 
Benefits: Built-in frontmatter validation, type safety, automatic routing 
1. Drawback: Requires moving files into the Astro project structure

Here's how to implement Approach 1 (Astro Content Collections) for your external markdown files:

## Step 1: Move Your Markdown Files

Move your markdown files from their current location to:
```
lessons/frontend/src/content/docs/
```

Create the directory structure if it doesn't exist. You can organize them however makes sense:
```
lessons/frontend/src/content/
├── docs/
│   ├── file1.md
│   ├── file2.md
│   ├── special-topic.md
│   └── reference-guide.md
└── config.ts
```

## Step 2: Define Content Collection Schema

Create `lessons/frontend/src/content/config.ts`:

```typescript
import { defineCollection, z } from 'astro:content';

const docsCollection = defineCollection({
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    publishDate: z.date().optional(),
    tags: z.array(z.string()).optional(),
    draft: z.boolean().default(false),
  }),
});

export const collections = {
  'docs': docsCollection,
};
```

## Step 3: Add Frontmatter to Your Markdown Files

Update each markdown file to include YAML frontmatter:

```markdown
---
title: "Your Document Title"
description: "Brief description of the content"
publishDate: 2024-01-01
tags: ["tutorial", "guide"]
draft: false
---

# Your existing markdown content starts here
```

## Step 4: Create Dynamic Routes

Create `lessons/frontend/src/pages/docs/[...slug].astro`:

```astro
---
import { getCollection } from 'astro:content';
import Layout from '../../layouts/Layout.astro';

export async function getStaticPaths() {
  const docs = await getCollection('docs');
  return docs.map((doc) => ({
    params: { slug: doc.slug },
    props: { doc },
  }));
}

const { doc } = Astro.props;
const { Content } = await doc.render();
---

<Layout title={doc.data.title}>
  <main>
    <h1>{doc.data.title}</h1>
    {doc.data.description && <p>{doc.data.description}</p>}
    <Content />
  </main>
</Layout>
```

## Step 5: Create an Index Page

Create `lessons/frontend/src/pages/docs/index.astro` to list all documents:

```astro
---
import { getCollection } from 'astro:content';
import Layout from '../../layouts/Layout.astro';

const docs = await getCollection('docs');
---

<Layout title="Documentation">
  <main>
    <h1>Documentation</h1>
    <ul>
      {docs.map((doc) => (
        <li>
          <a href={`/docs/${doc.slug}`}>{doc.data.title}</a>
          {doc.data.description && <p>{doc.data.description}</p>}
        </li>
      ))}
    </ul>
  </main>
</Layout>
```

## Benefits of This Approach:

- **Type Safety**: Astro validates frontmatter against your schema
- **Automatic Processing**: All markdown files are automatically included in builds
- **SEO Friendly**: Generates static pages with proper metadata
- **Hot Reloading**: Changes to markdown files trigger rebuilds in development
- **Flexible Organization**: You can organize files in subdirectories within `src/content/docs/`

## Next Steps:

1. Move your markdown files to the content directory
2. Add frontmatter to each file
3. Test the setup with `npm run dev`
4. Your files will be accessible at URLs like `/docs/your-filename`

This approach integrates cleanly with your existing Astro setup and ensures all your markdown content is properly included in the dist folder.