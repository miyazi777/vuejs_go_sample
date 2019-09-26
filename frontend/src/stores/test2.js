import axios from 'axios';

export const test2 = {
  namespaced: true,
  state: {
    book: {},
  },
  mutations: {
    updateBook (state, data) {
      console.log('mu');
      console.log(data);
      state.book = data
    }
  },
  actions: {
    async getBook(context, id) {
      await axios.get('http://localhost:8000/test2/' + id)
        .then((res) => {
          context.commit('updateBook', res.data);
        });
    }
  }
}
