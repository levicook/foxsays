'use strict';

var can = window.can,
    rhtml = require('rhtml');

module.exports = can.Control.extend({
    defaults: {
        view: can.view.mustache(rhtml('./main.mustache'))
    }
}, {
    init: function () {
        var control = this,
            element = control.element,
            options = control.options;

        control.model = {};

        control.helpers = {};

        element.append(options.view(control.model, control.helpers));
    }
});
