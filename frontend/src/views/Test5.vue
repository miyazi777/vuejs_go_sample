<template>
  <div>
    <h1>Test Page5(画面で入力されたものをAPI経由で新規登録)</h1>
    <form @submit.prevent="save">
      <input type="text" v-model="book.name" />
      <input type="text" v-model="book.price" />
      <button type="submit">submit</button>
    </form>
  </div>
</template>

<script>
  export default {
    name: "book",
    data() {
      return {
        book: {}
      }
    },
    methods: {
      save() {
        this.$store.dispatch("test5/save", this.book);
      }
    }
  }
</script>
