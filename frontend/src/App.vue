<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link> |
      <router-link to="/count">Count</router-link> |
      <router-link to="/books">Books</router-link> |
      <router-link to="/test">Test</router-link> |
      <router-link v-bind:to="{ name : 'test2', params : { id: this.id }}">Test2</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>
  //import _ from 'lodash';

  export default {
    data() {
      return {
        id: 2
      }
    },
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
