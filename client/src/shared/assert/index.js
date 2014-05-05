'use strict';

module.exports = function () {
    var args = can.makeArray(arguments),
        object = args.shift();

    args.forEach(function (property) {
        if (object[property] === undefined) {
            throw new Error(['missing property: ', '"', property, '"'].join(''));
        }
    });
};
