const mix = require('laravel-mix');

// Fix path issue on Windows
mix.setPublicPath('./');

mix.sass('assets/scss/style.scss', 'public/dist/css')
    .js('assets/js/index.js', 'public/dist/js');
