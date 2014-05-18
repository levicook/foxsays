'use strict';

var Header = require('./');

suite('main/header', function () {
    var element, header;

    setup(function () {
        element = $('<div>');
        header = new Header(element);
    });

    test('appended .main-header', function () {
        tt(element.find('.main-header').is('*'));
    });
});
