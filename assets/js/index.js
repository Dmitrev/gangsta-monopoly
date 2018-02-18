import Vue from 'vue';
import Vuex from 'vuex';
import Store from './store/store';

// Vue components
import TurnIndicator from '../vue/turnIndicator'
Vue.use(Vuex);

new Vue({
    el: '#app',

    data: {
        ws: null, // websocket
        screen: "start",
        players: [],
        myTurn: false
    },
    components: {
        TurnIndicator
    },
    store: Store,
    methods:{
        // Linked to the
        start: function () {
            var self = this;
            this.ws = new WebSocket('ws://' + window.location.host + '/ws');

            this.ws.onmessage = function (event) {
                var json = JSON.parse(event.data);
                console.log(json);
                switch(json.action) {
                    case 'register': self.register(); break;
                    case 'register_ok': self.goToLobby(); break;
                    case 'position_update': self.updatePlayers(json.data); break;
                    case 'game_started': self.startGame(); break;
                    case 'next_turn': self.updateCurrentPlayer(json.data); break;
                }
            }
        },
        goToLobby: function () {
            this.screen = "lobby";
        },
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
        ready: function () {
            this.ws.send(JSON.stringify({
                action: "ready"
            }));
        },
        throwDice: function(){
          this.ws.send(JSON.stringify({
              action: "throw_dice"
          }));
        },
        startGame: function () {
            this.screen = "game";
        },
        updatePlayers: function (players) {
            this.players = players
        },
        updateCurrentPlayer(player){
            this.$store.commit("setCurrentPlayer", player);
        }
    }
});
