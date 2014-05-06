'use strict';

require('../../helpers/global_helpers');

var rhtml = require('rhtml');

module.exports = can.Control.extend({
    defaults: {
        view: can.view.mustache(rhtml('./main.mustache'))
    }
}, {
    init: function (element, options) {
        var Header = require('../../components/header'),
            Footer = require('../../components/footer');

        this.model = {};
        this.helpers = {};

        element.append(options.view(this.model, this.helpers));

        this.header = new Header(element.find('.headerHook'));
        this.footer = new Footer(element.find('.footerHook'));
    }
});
