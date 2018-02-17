import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    players: []
  },
  getters: {
    players(state){
      return state.players;
    }
  },
  mutations: {
    setPlayers(state, players){
      state.players = players;
    }
  }
});

export default store;
