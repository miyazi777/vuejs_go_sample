import axios from 'axios';

export const test4 = {
  namespaced: true,
  state: {
    books: {},
    page: {}
  },
  getters: {
    books: state => {
      return state.books
    },
    page: state => {
      return state.page
    }
  },
  mutations: {
    setBooks (state, data) {
      state.books = data
    },
    setPage(state, data) {
      state.page = data
    }
  },
  actions: {
    async getBooks(context, params) {
      await axios.get('http://localhost:8000/test4?page=' + params)
        .then((res) => {
          context.commit('setBooks', res.data.books);
          context.commit('setPage', res.data.page);
        });
    }
  }
}
