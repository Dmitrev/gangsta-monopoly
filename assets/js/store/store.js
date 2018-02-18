import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    players: [],
    currentPlayer: null,
    myTurn: false,
    ws: null,
  },
  getters: {
    players(state){
        return state.players;
    },
    currentPlayer(state){
        return state.currentPlayer;
    },
    myTurn(state){
        return state.myTurn;
    },
    ws(state){
        return state.ws;
    }
  },
  mutations: {
    setPlayers(state, players){
      state.players = players;
    },
    setCurrentPlayer(state, player) {
      state.currentPlayer = player;
    },
    setMyTurn(state, myTurn){
        state.myTurn = myTurn;
    },
    setWs(state, ws){
        state.ws = ws;
    }
  }
});

export default store;
