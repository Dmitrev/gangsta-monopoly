var gulp = require('gulp');
var babel = require('gulp-babel');
var sass = require('gulp-sass');
var vueify = require('gulp-vueify');


gulp.task('default', ['js', 'sass'])


gulp.task('js', function () {
    return gulp.src('assets/**/*.js')
        .pipe(babel())
        .pipe(gulp.dest('public/dist'));
});


gulp.task('vueify', function () {
    return gulp.src('components/**/*.vue')
        .pipe(vueify())
        .pipe(gulp.dest('public/dist'));
});

gulp.task('sass', function () {
    return gulp.src('assets/**/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('public/dist'));
});

gulp.task('sass:watch', function () {
    gulp.watch('assets/**/*.scss', ['sass']);
});

gulp.task('js:watch', function () {
    gulp.watch('assets/**/*.js', ['js']);
});
