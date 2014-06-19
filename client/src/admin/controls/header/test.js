'use strict';

var Header = require('./');

suite('admin/header', function () {
	var element, header;

	setup(function () {
		element = $('<div>');
		header = new Header(element);
	});

	test('appended .admin-header', function () {
		tt(element.find('.admin-header').is('*'));
	});
});
