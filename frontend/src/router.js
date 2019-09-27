import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/count',
      name: 'count',
      component: () => import('./views/Count.vue')
    },
    {
      path: '/books',
      name: 'book',
      component: () => import('./views/Books.vue')
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('./views/Test.vue')
    },
    {
      path: '/test2/:id',
      name: 'test2',
      component: () => import('./views/Test2.vue')
    },
    {
      path: '/test3',
      name: 'test3',
      component: () => import('./views/Test3.vue')
    },
    {
      path: '/test4',
      name: 'test4',
      component: () => import('./views/Test4.vue')
    }
  ]
})
