/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        blue: "#003049",
        red: "#d62828",
        orange: "#f77f00",
        yellow: "#fcbf49",
        beige: "#eae2b7",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
