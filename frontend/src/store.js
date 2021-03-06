import Vue from 'vue'
import Vuex from 'vuex'

import {counter} from './stores/counter';
import {book} from './stores/book';
import {test} from './stores/test';
import {test2} from './stores/test2';
import {test3} from './stores/test3';
import {test4} from './stores/test4';
import {test5} from './stores/test5';

Vue.use(Vuex)

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    counter: counter,
    book: book,
    test: test,
    test2: test2,
    test3: test3,
    test4: test4,
    test5: test5,
  }
})
