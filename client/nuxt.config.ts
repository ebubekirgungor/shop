export default defineNuxtConfig({
  devtools: { enabled: true },
  devServer: {
    https: {
      key: "../certs/RootCA.key",
      cert: "../certs/RootCA.pem",
    },
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE,
      apiKey: process.env.API_KEY,
    },
  },
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
