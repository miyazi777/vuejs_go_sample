
export const book = {
  namespaced: true,
  state: {
    title: '',
    price: 0
  },
  mutations: {
    update(state, value) {
      console.log(value);
      state = value
      console.log('title = ' + state.title);
      console.log('price = ' + state.price);
    }
  }
}
