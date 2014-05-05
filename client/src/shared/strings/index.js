'use strict';

exports.contains = function (string, substring) {
    string = String(string);
    return string.indexOf(substring) !== -1;
};

exports.hasPrefix = function (string, prefix) {
    string = String(string);
    return string.indexOf(prefix) === 0;
};

exports.hasSuffix = function (string, suffix) {
    string = String(string);
    return string.indexOf(suffix, string.length - suffix.length) !== -1;
};
