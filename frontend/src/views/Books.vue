<template>
  <div>
    <h1>Books Page</h1>
    <form @submit.prevent="save">
      <input type="text" v-model="book.title" />
      <input type="text" v-model="book.price" />
      <button type="submit">submit</button>
    </form>
  </div>
</template>

<script>
  import _ from 'lodash';

  export default {
    name: "book",
    beforeMount() {
      this.book = _.cloneDeep(this.$store.state.book)
    },
    data() {
      return {
        book: {}
      }
    },
    methods: {
      save() {
        this.$store.commit("book/update", this.book);
      },
      async fetchItem() {
        const res = await axios.get('http://localhost:8000/item/3')
        console.log(res.data);
        this.items = res.data
      }
    }
  }
</script>
