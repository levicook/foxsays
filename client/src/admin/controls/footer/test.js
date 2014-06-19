'use strict';

var Footer = require('./');

suite('admin/footer', function () {
	var element, footer;

	setup(function () {
		element = $('<div>');
		footer = new Footer(element);
	});

	test('appended .admin-footer', function () {
		tt(element.find('.admin-footer').is('*'));
	});
});
