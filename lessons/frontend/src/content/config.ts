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