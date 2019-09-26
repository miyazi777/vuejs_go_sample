import axios from 'axios';

export const test = {
  namespaced: true,
  state: {
    msg: '',
  },
  //getters: {
  //  msg(state) {
  //    return state.msg
  //  }
  //},
  mutations: {
    updateMsg (state, data) {
      state.msg = data.msg
    }
  },
  actions: {
    async getMsg(context) {
      await axios.get('http://localhost:8000/test')
        .then((res) => {
          context.commit('updateMsg', res.data);
        });
    }
  }
}
