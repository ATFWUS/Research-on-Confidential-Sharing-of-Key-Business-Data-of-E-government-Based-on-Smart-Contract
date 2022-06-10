import Vue from 'vue'
import Vuex from 'vuex'
import tab from './tab'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    menuActive: "1"
  },
  mutations: {
  },
  actions: {
  },
  modules: {
	   tab,
	   
  }
})
