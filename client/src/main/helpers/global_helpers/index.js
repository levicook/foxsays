'use strict';

can.Mustache.registerHelper('link-to', function (route /*, options */) {
	return '#todo-link-to-' + route;
});
