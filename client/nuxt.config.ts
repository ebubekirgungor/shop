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
      Poppins: [400, 500, 700],
    },
  },
  modules: ["@nuxtjs/google-fonts", "@nuxtjs/i18n"],
  i18n: {
    vueI18n: "./i18n.config.ts",
    locales: [
      {
        code: "tr",
        name: "Türkçe",
      },
      {
        code: "en",
        name: "English",
      },
    ],
    strategy: "prefix_except_default",
    defaultLocale: "tr",
    customRoutes: "config",
    pages: {
      "account/index": {
        tr: "/hesap/",
        en: "/account/",
      },
      "account/orders/index": {
        tr: "/hesap/siparisler/",
        en: "/account/orders/",
      },
      "account/orders/[order]": {
        tr: "/hesap/siparisler/[order]",
        en: "/account/orders/[order]",
      },
      "account/addresses": {
        tr: "/hesap/adresler",
        en: "/account/addresses",
      },
      "account/change-password": {
        tr: "/hesap/sifre-degistir",
        en: "/account/change-password",
      },
      "account/favorites": {
        tr: "/hesap/favoriler",
        en: "/account/favorites",
      },
      "account/payment-methods": {
        tr: "/hesap/odeme-yontemleri",
        en: "/account/payment-methods",
      },
      "account/personal-details": {
        tr: "/hesap/kisisel-detaylar",
        en: "/account/personal-details",
      },
      "admin/index": {
        tr: "/yonetim/",
        en: "/admin/",
      },
      "admin/categories/index": {
        tr: "/yonetim/kategoriler/",
        en: "/admin/categories/",
      },
      "admin/products/index": {
        tr: "/yonetim/urunler/",
        en: "/admin/products/",
      },
      "admin/products/[product]": {
        tr: "/yonetim/urunler/[product]",
        en: "/admin/products/[product]",
      },
      "product/[product]": {
        tr: "/urun/[product]",
        en: "/product/[product]",
      },
      cart: {
        tr: "/sepet",
        en: "/cart",
      },
      checkout: {
        tr: "/odeme",
        en: "/checkout",
      },
      login: {
        tr: "/giris",
        en: "/login",
      },
      register: {
        tr: "/kayit",
        en: "/register",
      },
    },
  },
});
