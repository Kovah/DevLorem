const colors = require('tailwindcss/colors');

module.exports = {
  content: [
    './templates/index.html',
  ],
  theme: {
    extend: {
      colors: {
        teal: colors.teal,
        gray: colors.gray
      }
    }
  }
}
