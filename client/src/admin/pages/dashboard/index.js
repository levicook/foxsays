'use strict';

require('../../helpers/global_helpers');

var rhtml = require('rhtml');

var Footer = require('../../controls/footer');
var Header = require('../../controls/header');



module.exports = can.Control.extend({
	defaults: {
		view: can.stache(rhtml('./main.mustache'))
	}
}, {
	init: function (element, options) {
		// validateOptions(options);

		this.model = {};
		this.helpers = {};

		element.append(options.view(this.model, this.helpers));

		this.header = new Header(element.find('.headerHook'));

		this.footer = new Footer(element.find('.footerHook'));
	}
});
