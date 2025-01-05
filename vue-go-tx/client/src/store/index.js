import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    sharedData:{
      username:null
    },
    pages:{
      page:'msg',
      content:'chat',
    },
  },
  mutations: {
    setUser(state, username) {
      state.sharedData.username = username;
    },
    setPage(state, page) {
      state.pages.page = page; // 修改 page
    },
    setContent(state, content) {
      state.pages.content = content; // 修改 content
    },
  },
  actions: {
    setUser(context, username) {
      context.commit('setUser', username);
    },
    updatePage(context, page) {
      context.commit('setPage', page); // 调用 mutation 更新 page
    },
    updateContent(context, content) {
      context.commit('setContent', content); // 调用 mutation 更新 content
    },
  }
});

export default store;