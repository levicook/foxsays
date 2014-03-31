'use strict';

var header = require('../../components/member_header/index.js'),
    footer = require('../../components/member_footer/index.js');

module.exports = {

    controller: function () {
        this.header = new header.controller();
        this.footer = new footer.controller();
    },

    view: function (ctrl) {
        this.header = new header.view(ctrl.header);
        this.footer = new footer.view(ctrl.footer);

        return m('.forgot_password', [
            m('.header', [header]),
            m('h2', 'Forgot Password'),
            m('.footer', [footer])
        ]);
    }
};
