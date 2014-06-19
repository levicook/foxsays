'use strict';

var Footer = require('./');

suite('main/footer', function () {
    var element, footer;

    setup(function () {
        element = $('<div>');
        footer = new Footer(element);
    });

    test('appended .main-footer', function () {
        tt(element.find('.main-footer').is('*'));
    });
});
