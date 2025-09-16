import { defineCollection, z } from 'astro:content';

const docsCollection = defineCollection({
  type: 'content', // Explicitly set content type
  schema: z.object({
    title: z.string(),
    description: z.string().optional(),
    publishDate: z.date().transform((val) => val || new Date()),
    tags: z.array(z.string()).optional(),
    draft: z.boolean().default(false),
  }),
});

export const collections = {
  'docs': docsCollection,
};