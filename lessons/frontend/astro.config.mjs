import { defineConfig } from 'astro/config';
import react from '@astrojs/react';
import tailwind from '@astrojs/tailwind';
import node from '@astrojs/node';

export default defineConfig({
  integrations: [react(), tailwind()],
  output:'static',
  outDir: 'dist'
  // output: 'hybrid',
  // adapter: node ({
  //   mode: 'standalone'
  // }),
  // server: {
  //   proxy: {
  //     '/api': 'http://localhost:8080'
  //   }
  // }
});