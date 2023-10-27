export default defineNuxtConfig({
  devtools: { enabled: true },
  app: {
    head: {
      title: "Shop",
    },
  },
  css: ["~/assets/css/main.css"],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: ["@pinia/nuxt", "@pinia-plugin-persistedstate/nuxt"],
});
