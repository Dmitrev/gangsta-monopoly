'use strict';

new Vue({
    el: '#app',

    data: {
        ws: null // websocket
    },
    
    methods:{
      register: function () {
        var name = prompt("What's your name?", "Player");

        if (name.length !== 0) {
            this.ws.send(JSON.stringify({
                action: "register",
                data: name
            }));
        }
        else{
            this.register(); // Retry
        }
      },
      registerOk: function(){
        alert('Register complete!');
      }
    },
    created: function () {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');

        this.ws.onmessage = function (event) {
            var json = JSON.parse(event.data);
            console.log(json);

            switch(json.action) {
                case 'register': self.register(); break;
                case 'register_ok': self.registerOk(); break;
            }
        }
    }
});