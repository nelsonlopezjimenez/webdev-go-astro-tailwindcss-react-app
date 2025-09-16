import { defineCollection, z } from 'astro:content';

const docsCollection = defineCollection({
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    publishDate: z.date().default(() => new Date()), // Defaults to current date
    tags: z.array(z.string()).optional(),
    draft: z.boolean().default(false),
  }),
});

export const collections = {
  'docs': docsCollection,
};