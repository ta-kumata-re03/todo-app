import { defineNuxtConfig } from 'nuxt3'

export default defineNuxtConfig({
  srcDir: 'src',
  meta: {
    link: [
      { rel: 'stylesheet', href: 'https://unpkg.com/modern-css-reset/dist/reset.min.css' },
    ],
  },
})
