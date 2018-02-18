import Vue from 'vue';
import Vuex from 'vuex';
import Store from './store/store';

// Vue components
import TurnIndicator from '../vue/turnIndicator';
import ThrowDiceButton from '../vue/throwDiceButton';

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
        TurnIndicator,
        ThrowDiceButton
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
                    case 'your_turn': self.updateMyTurn(); break;
                }
            };
            this.$store.commit('setWs', this.ws);

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
        startGame: function () {
            this.screen = "game";
        },
        updatePlayers: function (players) {
            this.players = players
        },
        updateCurrentPlayer(player){
            // Resets all players "myTurn" property
            this.$store.commit("setMyTurn", false);
            this.$store.commit("setCurrentPlayer", player);
        },
        updateMyTurn(){
            this.$store.commit("setMyTurn", true);
        }
    }
});
