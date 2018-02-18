import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        players: [],
        currentPlayer: null,
        myTurn: false,
        ws: null,
        thrownDice: false
    },
    getters: {
        players(state) {
            return state.players;
        },
        currentPlayer(state) {
            return state.currentPlayer;
        },
        myTurn(state) {
            return state.myTurn;
        },
        ws(state) {
            return state.ws;
        },
        thrownDice(state) {
            return state.thrownDice;
        }
    },
    mutations: {
        setPlayers(state, players) {
            state.players = players;
        },
        setCurrentPlayer(state, player) {
            state.currentPlayer = player;
        },
        setMyTurn(state, myTurn) {
            state.myTurn = myTurn;
        },
        setWs(state, ws) {
            state.ws = ws;
        },
        setThrownDice(state, thrown) {
            state.thrownDice = thrown;
        }
    }
});

export default store;
