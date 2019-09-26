import Vue from 'vue'
import Vuex from 'vuex'

import {counter} from './stores/counter';
import {book} from './stores/book';
import {test} from './stores/test';
import {test2} from './stores/test2';

Vue.use(Vuex)

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    counter: counter,
    book: book,
    test: test,
    test2: test2,
  }
})
