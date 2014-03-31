'use strict';

var assert = window.assert,
    suite = window.suite,
    setup = window.setup,
    test = window.test,
    ForgotPassword = require('./');

suite('website/pages/forgot_password', function () {
    suite('ForgotPassword', function () {
        var element, forgotPassword;

        setup(function () {
            element = document.createElement('div');
            forgotPassword = new ForgotPassword(element);
        });

        test('is a ForgotPassword', function () {
            assert.instanceOf(forgotPassword, ForgotPassword);
        });

        test('appended .forgot_password', function () {
            assert.ok(element.find('.forgot_password').is('*'));
        });

    });
});
