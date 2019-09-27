import axios from 'axios';

export const test3 = {
  namespaced: true,
  state: {
    books: {}
  },
  getters: {
    books: state => {
      return state.books
    }
  },
  mutations: {
    setBooks (state, data) {
      state.books = data
    }
  },
  actions: {
    async getBooks(context) {
      await axios.get('http://localhost:8000/test3')
        .then((res) => {
          context.commit('setBooks', res.data);
        });
    }
  }
}
