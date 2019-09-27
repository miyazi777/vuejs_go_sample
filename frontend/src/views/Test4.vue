<template>
  <div class="about">
    <h1>Test Page4(配列をリスト表示(ページネイト付き))</h1>
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

        <td><router-link v-bind:to="{ name: 'test2', params: { id: book.Id }}">データを参照</router-link></td>
      </tr>
    </table>
    <div v-for="p of page.pageCount" :key="p">
      <p @click="getBooks(p)"><span v-if="currentPage == p">*</span>page{{p}}</p>
    </div>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'

  export default {
    name: "test4",
    created() {
      this.currentPage = this.$route.query.page || 1
      console.log(this.currentPage);
      this.$store.dispatch('test4/getBooks', this.currentPage-1)
    },
    data() {
      return {
        currentPage: 1
      }
    },
    computed: {
      ...mapGetters({
        books: 'test4/books',
        page: 'test4/page',
      })
    },
    methods: {
      getBooks(page) {
        this.currentPage = page
        this.$store.dispatch('test4/getBooks', this.currentPage-1)
      }
    }
  }
</script>
