const cssnano = require('cssnano')({
  preset: 'default'
});

module.exports = {
  plugins: [
    "@tailwindcss/postcss",
    require('autoprefixer'),
    ...process.env.NODE_ENV === 'production'
      ? [cssnano]
      : []
  ]
};
