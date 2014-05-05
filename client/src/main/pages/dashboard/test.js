'use strict';

var Dashboard = require('./');

suite('main/dashboard', function () {
    var element, dashboard;

    setup(function () {
        element = $('<div>');
        dashboard = new Dashboard(element);
    });

    test('appended .main-header', function () {
        tt(element.find('.main-header').is('*'));
    });

    test('appended .main-dashboard', function () {
        tt(element.find('.main-dashboard').is('*'));
    });

    test('appended .main-footer', function () {
        tt(element.find('.main-footer').is('*'));
    });

});
