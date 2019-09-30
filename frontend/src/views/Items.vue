<template>
  <div>
    <h1>Items Page</h1>
    <div><router-link to="/item/new">新規作成</router-link></div>
    <table>
      <tr>
        <td>ID</td>
        <td>名前</td>
      </tr>
      <tr v-for="item in items" v-bind:key="item.id">
        <td>{{ item.id }}</td>
        <td>{{ item.name }}</td>

        <td><router-link v-bind:to="{ name: 'itemDetail', params: { id: item.id }}">データを参照</router-link></td>
        <td><router-link v-bind:to="{ name: 'itemUpdate', params: { id: item.id }}">データを編集</router-link></td>
      </tr>
    </table>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "items",
    data() {
      return {
        items: []
      }
    },
    methods: {
      async fetchItems() {
        const res = await axios.get('http://localhost:8000/items')
        console.log(res.data);
        this.items = res.data
      }
    },
    watch: {
      $route: {
        async handler() {
          await this.fetchItems()
        },
        immediate: true
      }
    }
  }
</script>
