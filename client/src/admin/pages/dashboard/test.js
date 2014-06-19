'use strict';

var Dashboard = require('./');

suite('admin/dashboard', function () {
	var element, dashboard;

	setup(function () {
		element = $('<div>');
		dashboard = new Dashboard(element);
	});

	test('appended .admin-header', function () {
		tt(element.find('.admin-header').is('*'));
	});

	test('appended .admin-dashboard', function () {
		tt(element.find('.admin-dashboard').is('*'));
	});

	test('appended .admin-footer', function () {
		tt(element.find('.admin-footer').is('*'));
	});

});
