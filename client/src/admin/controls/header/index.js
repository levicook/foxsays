'use strict';

var rhtml = require('rhtml');

module.exports = can.Control.extend({
	defaults: {
		view: can.stache(rhtml('./main.mustache'))
	}
}, {
	init: function (element, options) {
		// validateOptions(options);

		element.append(options.view());
	}
});
