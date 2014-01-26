'use strict';

var assert = window.assert,
    suite = window.suite,
    setup = window.setup,
    test = window.test,
    $ = window.jQuery,
    PublicHome = require('./website/pages/public_home');

suite('website/pages/public_home', function () {
    suite('PublicHome', function () {
        var element, publicHome;

        setup(function () {
            element = $('<div/>');
            publicHome = new PublicHome(element, {});
        });

        test('is a PublicHome', function () {
            assert.instanceOf(publicHome, PublicHome);
        });

    });
});
