import axios from 'axios';

export const test5 = {
  namespaced: true,
  state: {
    name: '',
    price: 0
  },
  mutations: {
    save(state, data) {
      state = data
    }
  },
  actions: {
    save(context, data) {
      context.commit('save', data);
      axios.post('http://localhost:8000/test5', {
        Name: data.name,
        Price: parseInt(data.price)
      }).then(res => {
        console.log('response');
        console.log(res.data);
      })
    }
  }
}
