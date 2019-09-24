
export const counter = {
  namespaced: true,
  state: {
    count: 5
  },
  mutations: {
    increment(state) {
      state.count++
    },
    decrement(state) {
      state.count--
    }
  },
  actions: {
    increment: ({ commit }) => {
      setTimeout(() => {
        commit('increment')
      }, 1000)
    }
  }
}
