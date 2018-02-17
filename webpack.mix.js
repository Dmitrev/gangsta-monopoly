const mix = require('laravel-mix');


mix.sass('assets/scss/style.scss', 'public/dist/css')
    .js('assets/js/index.js', 'public/dist/js');
