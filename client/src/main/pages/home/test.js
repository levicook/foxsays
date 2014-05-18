'use strict';

var Page = require('./');

suite('main/home', function () {
	var element, page;

	setup(function () {
		element = $('<div>');
		page = new Page(element);
	});

	test('appended .main-header', function () {
		tt(element.find('.main-header').is('*'));
	});

	test('appended .main-home', function () {
		tt(element.find('.main-home').is('*'));
	});

	test('appended .main-footer', function () {
		tt(element.find('.main-footer').is('*'));
	});

});
