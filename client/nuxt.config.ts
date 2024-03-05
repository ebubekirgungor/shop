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
  googleFonts: {
    families: {
      Poppins: [400, 500],
    },
  },
  modules: ["@nuxtjs/google-fonts"],
});
