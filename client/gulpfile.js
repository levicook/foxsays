'use strict';


var EXPRESS_PORT    = 1111,
    LIVERELOAD_PORT = 1112,

    gulp       = require('gulp'),
    browserify = require('gulp-browserify'),
    clean      = require('gulp-clean'),
    refresh    = require('gulp-livereload'),
    livereload = require('tiny-lr')(),

    recess = require('gulp-recess'),
    less   = require('gulp-less');



gulp.task('clean', function () {
    gulp
    .src('build', { read: false })
    .pipe(clean());
});


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


gulp.task('website-pages-css', function () {
    gulp.src('src/website/pages/*/main.less')
    .pipe(recess())
    .pipe(less())
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});


gulp.task('website-pages-html', function () {
    gulp.src('src/website/pages/*/{demo,main,test}.html')
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});


gulp.task('website-pages-js', function () {
    gulp.src('src/website/pages/*/{demo,main,test}.js')
    .pipe(browserify({
        basedir: 'src',
        debug: true
    }))
    .on('prebundle', function (bundle) {
        bundle.external('bootstrap');
        bundle.external('can');
        bundle.external('jQuery');
    })
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});


gulp.task('website-shared-css', function () {
    gulp.src('src/website/shared.less')
    .pipe(recess())
    .pipe(less())
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});


gulp.task('website-shared-js', function () {
    gulp.src('src/website/shared.js')
    .pipe(browserify({
        shim: {
            bootstrap: {
                depends: { jQuery: 'jQuery' },
                exports: null,
                path: 'bower_components/bootstrap/dist/js/bootstrap.js'
            },
            can: {
                depends: { jQuery: 'jQuery' },
                exports: 'can',
                path: 'bower_components/canjs/can.jquery.js'
            },
            jQuery: {
                exports: 'jQuery',
                path: 'bower_components/jquery/jquery.js'
            }
        }
    }))
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});


gulp.task('default', [
    'express',
    'livereload',
    'website-pages-css',
    'website-pages-html',
    'website-pages-js',
    'website-shared-css',
    'website-shared-js',
], function () {
    gulp.watch('src/website/shared.less', ['website-shared-css']);
    gulp.watch('src/website/shared.js', ['website-shared-js']);
    gulp.watch('src/website/{components,pages}/*/*.html', ['website-pages-html']);
    gulp.watch('src/website/{components,pages}/*/*.{js,mustache}', ['website-pages-js']);
    gulp.watch('src/website/{components,pages}/*/*.less', ['website-pages-css']);
});
