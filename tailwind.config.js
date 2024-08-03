/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go", ],
  safelist: [],
  // theme: {
  //   extend: {},
  // },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark"]
  }
}

