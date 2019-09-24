import Vue from 'vue'
import Vuex from 'vuex'

import {counter} from './stores/counter';
import {book} from './stores/book';

Vue.use(Vuex)

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    counter: counter,
    book: book
  }
})
