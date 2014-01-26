'use strict';


var EXPRESS_PORT    = 1111,
    LIVERELOAD_PORT = 1112,

    browserify = require('gulp-browserify'),
    concat     = require('gulp-concat'),
    csso       = require('gulp-csso'),
    gulp       = require('gulp'),
    htmlhint   = require('gulp-htmlhint'),
    jshint     = require('gulp-jshint'),
    less       = require('gulp-less'),
    livereload = require('tiny-lr')(),
    recess     = require('gulp-recess'),
    refresh    = require('gulp-livereload'),
    stylish    = require('jshint-stylish');


// ------------------------------------------

gulp.task('website-css', ['website-pages-css', 'website-shared-css']);

gulp.task('website-pages-css', function () {
    return gulp
    .src('src/website/pages/*/main.less')
    .pipe(recess())
    .pipe(less())
    .pipe(csso())
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});

gulp.task('website-shared-css', function () {
    return gulp
    .src('src/website/shared.less')
    .pipe(recess({
        noOverqualifying: false,
        noUniversalSelectors: false,
        strictPropertyOrder: false,
        zeroUnits: false
    }))
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

gulp.task('website-js', ['website-pages-js', 'website-shared-js', 'website-test-js']);

gulp.task('website-jshint', function (cb) {
    return gulp
    .src(['src/website/pages/*/*.js', 'src/website/shared.js'])
    .pipe(jshint('.jshintrc'))
    .pipe(jshint.reporter(stylish));
});

gulp.task('website-pages-js', ['website-jshint'], function () {
    return gulp
    .src('src/website/pages/*/{demo,main}.js')
    .pipe(browserify({
        basedir: 'src',
        detectGlobals: false,
        insertGlobals: false,
        transform: ['rfileify']
    }))
    .pipe(gulp.dest('build/website/pages'))
    .pipe(refresh(livereload));
});

gulp.task('website-shared-js', ['website-jshint'], function () {
    return gulp
    .src([
        './bower_components/jquery/jquery.js',
        './bower_components/bootstrap/dist/js/bootstrap.js',
        './bower_components/canjs/can.jquery.js',
    ])
    .pipe(concat('shared.js'))
    .pipe(gulp.dest('build/website'))
    .pipe(refresh(livereload));
});

gulp.task('website-test-js', ['website-jshint'], function () {
    return gulp
    .src('src/website/{components,pages}/*/test.js')
    .pipe(concat('test.js'))
    .pipe(browserify({
        basedir: 'src',
        detectGlobals: false,
        insertGlobals: false,
        transform: ['rfileify']
    }))
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

    gulp.watch('src/website/variables.less',                       [ 'website-css'        ]);
    gulp.watch('src/website/shared.less',                          [ 'website-shared-css' ]);
    gulp.watch('src/website/{components,pages}/*/*.less',          [ 'website-pages-css'  ]);

    gulp.watch('src/website/test.html',                            [ 'website-test-html'  ]);
    gulp.watch('src/website/{components,pages}/*/*.html',          [ 'website-pages-html' ]);

    gulp.watch('src/website/{components,pages}/*/*.{js,mustache}', [ 'website-pages-js', 'website-test-js' ]);
});

gulp.task('default', ['assets', 'watch']);
