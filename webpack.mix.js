const { mix } = require('laravel-mix');

mix
  .react('resources/assets/js/index.js', 'dist/js')
  .sass('resources/assets/sass/app.scss', 'dist/css')
  .version()
  .sourceMaps();
