import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./web/templates/**/*.{js,ts,html,gohtml}"],
  theme: {
    extend: {
      container: {
        center: true,
        padding: {
          DEFAULT: "1rem",
          md: "2rem",
        }
      }
    },
  },
  plugins: [typography],
}

