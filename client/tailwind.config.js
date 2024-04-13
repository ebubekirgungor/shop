/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
  ],
  theme: {
    extend: {
      fontFamily: {
        poppins: ["Poppins", "sans-serif"],
      },
      gridTemplateColumns: {
        auto: "repeat(auto-fill, minmax(200px, 1fr))",
        auto_box: "repeat(auto-fill, minmax(14rem, 1fr))",
        auto_box_orders: "repeat(auto-fill, minmax(6rem, 1fr))",
        auto_box_checkout: "repeat(auto-fill, minmax(13rem, 1fr))",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
