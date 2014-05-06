'use strict';

var Home = require('./');

suite('main/home', function () {
    var element, home;

    setup(function () {
        element = $('<div>');
        home = new Home(element);
    });

    test('appended .main-header', function () {
        tt(element.find('.main-header').is('*'));
    });

    test('appended .main-home', function () {
        tt(element.find('.main-home').is('*'));
    });

    test('appended .main-footer', function () {
        tt(element.find('.main-footer').is('*'));
    });

});
