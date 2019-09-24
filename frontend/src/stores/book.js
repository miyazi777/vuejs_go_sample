
export const book = {
  namespaced: true,
  state: {
    title: 'book1'
  },
  mutations: {
    title(state, title) {
      state.title = title
    },
  }
}
