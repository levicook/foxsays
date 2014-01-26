'use strict';


var EXPRESS_PORT = 1111,
    LIVERELOAD_PORT = 1112,

    gulp = require('gulp'),
    gulpBrowserify = require('gulp-browserify'),
    gulpClean = require('gulp-clean'),
    refresh = require('gulp-livereload'),
    livereload = require('tiny-lr')();



gulp.task('clean', function () {
    gulp
    .src('build', { read: false })
    .pipe(gulpClean());
});


gulp.task('express', function () {
    var express = require('express'),
        connectLiveReload = require('connect-livereload'),
        app = express();

    app.use(connectLiveReload({ port: LIVERELOAD_PORT }));
    app.use(express.directory(__dirname + '/build'));
    app.use(express.static(__dirname + '/build'));
    app.use(express.errorHandler());

    app.listen(EXPRESS_PORT);
});


gulp.task('livereload', function () {
    livereload.listen(LIVERELOAD_PORT);
});


gulp.task('website-vendor-js', function () {
    gulp
    .src('src/website/vendor.js')
    .pipe(gulpBrowserify({
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


gulp.task('website-pages-js', function () {
    gulp
    .src('src/website/pages/*/{demo,main,test}.js')
    .pipe(gulpBrowserify({
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


gulp.task('default', [
    'clean',
    'website-vendor-js',
    'website-pages-js',
    'express',
    'livereload',
], function () {
    gulp.watch('src/website/vendor.js', ['website-vendor-js']);
    gulp.watch('src/website/{components,pages}/*/*.js', ['website-pages-js']);
});
