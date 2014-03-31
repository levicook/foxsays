'use strict';


var EXPRESS_PORT    = 2222,
    LIVERELOAD_PORT = 2223,

    concat     = require('gulp-concat'),
    csso       = require('gulp-csso'),
    glob       = require('glob'),
    gulp       = require('gulp'),
    htmlhint   = require('gulp-htmlhint'),
    jshint     = require('gulp-jshint'),
    less       = require('gulp-less'),
    livereload = require('tiny-lr')(),
    refresh    = require('gulp-livereload'),
    stylish    = require('jshint-stylish');


// ------------------------------------------

gulp.task('website-css', ['website-pages-css', 'website-shared-css']);

gulp.task('website-pages-css', function () {
    return gulp
    .src('src/website/pages/*/main.less')
    .pipe(less())
    .pipe(csso())
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});

gulp.task('website-shared-css', function () {
    return gulp
    .src('src/website/shared.less')
    .pipe(less())
    .pipe(csso())
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});

// ------------------------------------------

gulp.task('website-html', ['website-pages-html', 'website-test-html']);

gulp.task('website-pages-html', function () {
    return gulp
    .src('src/website/pages/*/{demo,main}.html')
    .pipe(htmlhint())
    .pipe(htmlhint.reporter())
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});

gulp.task('website-test-html', function () {
    return gulp
    .src('src/website/test.html')
    .pipe(htmlhint())
    .pipe(htmlhint.reporter())
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});

// ------------------------------------------

gulp.task('website-js', []);

gulp.task('website-jshint', function (cb) {
    return gulp
    .src(['src/website/pages/*/*.js', 'src/website/shared.js'])
    .pipe(jshint('.jshintrc'))
    .pipe(jshint.reporter(stylish));
});

gulp.task('website-shared-js', function () {
    return gulp
    .src([
        './bower_components/mithril/mithril.min.js',
    ])
    .pipe(concat('shared.js'))
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});

// ------------------------------------------

gulp.task('assets', [
    'website-css',
    'website-html',
    'website-js'
]);

// ------------------------------------------

gulp.task('express', function () {
    var express = require('express'),
        connectLiveReload = require('connect-livereload'),
        app = express();

    app.use(connectLiveReload({ port: LIVERELOAD_PORT }));
    app.use(express.directory(__dirname + '/build'));
    app.use(express.static(__dirname + '/build'));
    app.use(express.static(__dirname));
    app.use(express.errorHandler());

    app.listen(EXPRESS_PORT);
});

gulp.task('livereload', function () {
    livereload.listen(LIVERELOAD_PORT);
});

gulp.task('watch', ['express', 'livereload'], function () {
    gulp.watch('src/website/variables.less',               [ 'website-css'                         ]);
    gulp.watch('src/website/shared.less',                  [ 'website-shared-css'                  ]);
    gulp.watch('src/website/{components,pages}/*/*.less',  [ 'website-pages-css'                   ]);
    gulp.watch('src/website/test.html',                    [ 'website-test-html'                   ]);
    gulp.watch('src/website/{components,pages}/*/*.html',  [ 'website-pages-html'                  ]);
    gulp.watch('src/website/{components,pages}/*/*.js',    [ 'website-test-js' ]);
});

gulp.task('default', ['assets', 'watch']);
