'use strict';


var EXPRESS_PORT    = 3081,
    LIVERELOAD_PORT = 3082,

    browserify = require('browserify'),
    cache      = require('gulp-cached'),
    changed    = require('gulp-changed'),
    concat     = require('gulp-concat'),
    cssmin     = require('gulp-cssmin'),
    gulp       = require('gulp'),
    gulpLR     = require('gulp-livereload'),
    jshint     = require('gulp-jshint'),
    less       = require('gulp-less'),
    rename     = require('gulp-rename'),
    rev        = require('gulp-rev'),
    tinyLR     = require('tiny-lr')(),
    uglify     = require('gulp-uglify'),

    strings   = require('./src/shared/strings'),
    hasPrefix = strings.hasPrefix,
    hasSuffix = strings.hasSuffix,

    vendorJS = {
        demo: [
            './bower_components/canjs/can.object.js',
            './bower_components/canjs/can.fixture.js',
        ],
        main: [
            './bower_components/uri.js/src/URI.js',
            './bower_components/canjs/can.jquery.min.js',
            './bower_components/canjs/can.list.sort.js',
        ],
        test: [
            './bower_components/mocha/mocha.js',
            './bower_components/chai/chai.js',
        ]
    };

// ------------------------------------------

function shouldBrowserify(reqPath) {
    return !hasPrefix(reqPath, '/vendor') && hasSuffix(reqPath, '.js');
}

function newBrowserifyFor(reqPath) {
    var b;
    b = browserify('./src' + reqPath);
    b.transform('rfileify');
    return b;
}

// ------------------------------------------

gulp.task('clear', function () {
    process.stdout.write('\u001B[2J\u001B[0;0f');
});

// ------------------------------------------

gulp.task('express', function () {
    var browserify  = require('browserify'),
        connectLR   = require('connect-livereload'),
        express     = require('express'),
        app         = express(),
        serveIndex  = require('serve-index'),
        serveStatic = require('serve-static');

    app.use(function (req, res, next) {
        if (shouldBrowserify(req.path)) {
            res.set('Content-Type', 'application/javascript');
            newBrowserifyFor(req.path).bundle({ detectGlobals: false }).pipe(res);
        } else {
            next();
        }
    });

    app.use(function (req, res, next) {
        if (hasSuffix(req.path, '/demo.html')) {
            res.sendfile('./src/shared/layout/demo.html');
        } else {
            next();
        }
    });

    app.use(function (req, res, next) {
        if (hasSuffix(req.path, '/test.html')) {
            res.sendfile('./src/shared/layout/test.html');
        } else {
            next();
        }
    });

    app.use(connectLR({ port: LIVERELOAD_PORT }));

    app.use(serveIndex('./src', { icons: true }));
    app.use(serveStatic('./src'));
    app.use(serveStatic('./pkg'));
    app.use('/mocha', serveStatic('./bower_components/mocha'));

    app.use(require('errorhandler')())

    app.listen(EXPRESS_PORT);
});

// ------------------------------------------

gulp.task('htmlhint', function () {
    var htmlhint = require("gulp-htmlhint");

    return gulp.src("./src/**/*.{html,mustache}")
    .pipe(htmlhint({ htmlhintrc: '.htmlhintrc' }))
    .pipe(htmlhint.reporter())
    .pipe(gulpLR(tinyLR))
});

// ------------------------------------------

gulp.task('jshint', ['jshint:main', 'jshint:test']);

gulp.task('jshint:main', function() {
    return gulp
    .src(['./src/**/*.js', '!./src/**/test.js'])
    .pipe(cache('js'))
    .pipe(jshint('.main.jshintrc'))
    .pipe(jshint.reporter('jshint-stylish'))
    .pipe(gulpLR(tinyLR))
});

gulp.task('jshint:test', function() {
    return gulp
    .src('./src/**/test.js')
    .pipe(cache('js'))
    .pipe(jshint('.test.jshintrc'))
    .pipe(jshint.reporter('jshint-stylish'))
    .pipe(gulpLR(tinyLR))
});

// ------------------------------------------

gulp.task('pkg:vendor-main.js', function() {
    return gulp
    .src(vendorJS.main)
    .pipe(concat('vendor-main.js'))
    .pipe(gulp.dest('./pkg'))
});

gulp.task('pkg:vendor-demo.js', function() {
    return gulp.src(vendorJS.main.concat(vendorJS.demo))
    .pipe(concat('vendor-demo.js'))
    .pipe(gulp.dest('./pkg'))
});

gulp.task('pkg:vendor-test.js', function() {
    return gulp.src(vendorJS.main.concat(vendorJS.demo).concat(vendorJS.test))
    .pipe(concat('vendor-test.js'))
    .pipe(gulp.dest('./pkg'))
});

// ------------------------------------------

gulp.task('pkg:less', function() {
    return gulp
    .src('./src/**/*.less')
    .pipe(changed('./pkg', { extension: '.css' }))
    .pipe(less())
    .pipe(gulp.dest('./pkg'))
    .pipe(gulpLR(tinyLR))
});

// ------------------------------------------

gulp.task('dist:css', function () {
    return gulp
    .src('./pkg/*/pages/**/*.css')
    .pipe(gulp.dest('./dist/assets'))
    .pipe(cssmin())
    .pipe(rename({suffix: '.min'}))
    .pipe(gulp.dest('./dist/assets'))
});

gulp.task('dist:pages:js', function () {
    return gulp
    .src('./pkg/*/pages/**/main.js')
    .pipe(gulp.dest('./dist/assets'))
    .pipe(uglify({ outSourceMap: true }))
    .pipe(rename({ suffix: '.min' }))
    .pipe(gulp.dest('./dist/assets'))
});

gulp.task('dist:vendor-main.js', ['pkg:vendor-main.js'], function () {
    return gulp
    .src('./pkg/vendor-main.js')
    .pipe(gulp.dest('./dist/assets'))
    .pipe(uglify({ outSourceMap: true }))
    .pipe(rename({ suffix: '.min' }))
    .pipe(gulp.dest('./dist/assets'))
});

gulp.task('dist', [
    'dist:css',
    'dist:pages:js',
    'dist:vendor-main.js',
    'pkg:vendor-demo.js', // \__ So we don't muck up our dev environment.
    'pkg:vendor-test.js', // /
], function () {
    return gulp
    .src([
        './dist/assets/**/*.min.css',
        './dist/assets/**/*.min.js',
        './dist/assets/vendor-main.js'
    ])
    .pipe(rev())
    .pipe(gulp.dest('./dist/assets'))
    .pipe(rev.manifest())
    .pipe(gulp.dest('./dist/assets'))
});

// ------------------------------------------

gulp.task('livereload', function () {
    tinyLR.listen(LIVERELOAD_PORT);
});

// ------------------------------------------

gulp.task('watch', ['pkg'], function () {
    gulp.watch('./src/**/*.js', ['clear', 'jshint']);
    gulp.watch('./src/**/*.less', ['clear', 'pkg:less']);
    gulp.watch('./src/**/*.{html,mustache}', ['clear', 'htmlhint']);
});

// ------------------------------------------

gulp.task('pkg', [
    'pkg:less',
    'pkg:vendor-demo.js',
    'pkg:vendor-main.js',
    'pkg:vendor-test.js',
]);

gulp.task('default', [
    'express',
    'livereload',
    'watch',
], function () {
    console.log('[express] listening at http://127.0.0.1:' + EXPRESS_PORT);
    console.log('[livereload] listening at http://127.0.0.1:' + LIVERELOAD_PORT);
});
