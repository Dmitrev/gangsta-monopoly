'use strict';

new Vue({
    el: '#app',

    data: {
        ws: null // websocket
    },
    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
    }
});