'use strict';

var DemoControl = require('./control.js');

exports.run = function (Control, fixtures) {
    window.demo = new DemoControl(document.getElementById('demo'), {
        Control: Control,
        fixtures: fixtures
    });
};
