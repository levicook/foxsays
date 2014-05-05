'use strict';

var rhtml = require('rhtml');

module.exports = can.Control.extend({
    defaults: {
        view: can.view.mustache(rhtml('./main.mustache'))
    }
}, {
    init: function (element, options) {
        // validateOptions(options);

        element.append(options.view());
    }
});
