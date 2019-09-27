<template>
  <div class="about">
    <h1>Test Page3(配列をリスト表示)</h1>
    <table>
      <tr>
        <td>ID</td>
        <td>名前</td>
        <td>値段</td>
      </tr>
      <tr v-for="book in books" v-bind:key="book.id">
        <td>{{ book.Id }}</td>
        <td>{{ book.Name }}</td>
        <td>{{ book.Price }}</td>

        <td><router-link v-bind:to="{ name : 'test2', params : { id: book.Id }}">データを参照</router-link></td>
      </tr>
    </table>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'

  export default {
    name: "test3",
    created() {
      this.$store.dispatch('test3/getBooks')
    },
    computed: {
      ...mapGetters({
        books: 'test3/books'
      })
      //books() {
      //  return this.$store.state.test3.books
      //}
    }
  }
</script>
