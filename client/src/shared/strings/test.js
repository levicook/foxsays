'use strict';

var strings = require('./');

suite('shared/strings', function () {

    test('contains', function () {
        tt(strings.contains('fox', 'f'));
        tt(strings.contains('fox', 'o'));
        tt(strings.contains('fox', 'x'));
        ff(strings.contains('fox', 'z'));
    });

    test('hasPrefix', function () {
        tt(strings.hasPrefix('foo', 'f'));
        tt(strings.hasPrefix('foo', 'fo'));
        tt(strings.hasPrefix('foo', 'foo'));
        ff(strings.hasPrefix('foo', 'x'));
    });

    test('hasSuffix', function () {
        tt(strings.hasSuffix('foo', 'foo'));
        tt(strings.hasSuffix('foo', 'oo'));
        tt(strings.hasSuffix('foo', 'o'));
        ff(strings.hasSuffix('foo', 'x'));
    });

});
