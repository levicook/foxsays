'use strict';

var rhtml = require('rhtml');

module.exports = can.Control.extend({
    defaults: {
        view: can.view.mustache(rhtml('./main.mustache'))
    }
}, {

    init: function (element, options) {
        var fixtureId = -1;

        this.model = {
            names: options.fixtures.map(function (fixture) {
                return fixture.name;
            }),
            fixtureId: function() {
                return fixtureId += 1;
            }
        };

        element.append(options.view(this.model));
        this.demoHook = element.find('.demoHook');
        this.demoMenu = element.find('.demoMenu');

        can.route(':fixture', { fixture: 0 });
        can.route.ready();

        if (!can.route.attr('fixture')) {
            can.route.attr('fixture', 0);
        }
    },

    '{can.route} fixture': function(_1, _2, newVal) {
        newVal = parseInt(newVal, 10);

        this.demoMenu.find('.demo').
            removeClass('active').
            filter(':eq(' + newVal + ')').
            addClass('active');

        var options  = this.options,
            fixtures = options.fixtures,
            fixture  = fixtures[newVal] || fixtures[0];

        if (can.isFunction(fixture.setup)) {
            fixture.setup(fixture.name, fixture.data);
        }

        window.control = new options.Control(this.demoHook.empty(), fixture.data);
    }

});
