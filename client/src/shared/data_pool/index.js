'use strict';

module.exports = can.Construct.extend({

    init: function (bodyOrData) {
        if (bodyOrData === document.body) {
            this.rawData = JSON.parse(document.getElementById('dataPool').text);
        } else {
            this.rawData = bodyOrData || {};
        }

        this.cache = {};
    },

    getFirstPerson: function () {
        if (!this.cache.firstPerson) {
            var FirstPerson = require('../models/first_person');
            this.cache.firstPerson = new FirstPerson(this.rawData.firstPerson);
        }
        return this.cache.firstPerson;
    }

    //  update: function (data) {
    //  }

});
