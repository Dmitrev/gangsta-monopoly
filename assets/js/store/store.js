import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    players: [],
    currentPlayer: null,
  },
  getters: {
    players(state){
      return state.players;
    },
    currentPlayer(state){
      return state.currentPlayer;
    }
  },
  mutations: {
    setPlayers(state, players){
      state.players = players;
    },
    setCurrentPlayer(state, player) {
      state.currentPlayer = player;
    }
  }
});

export default store;
