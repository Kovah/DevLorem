const colors = require('tailwindcss/colors');

module.exports = {
  purge: [
    './templates/index.html',
  ],
  theme: {
    extend: {
      colors: {
        teal: colors.teal,
        gray: colors.gray
      }
    }
  },
  variants: {},
  plugins: [],
}
