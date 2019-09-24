<template>
  <div>
    <h1>Books Page</h1>
    <form @submit.prevent="save">
      <input type="text" v-model="title"/>
      <button type="submit">submit</button>
    </form>
  </div>
</template>

<script>
  export default {
    name: "book",
    computed: {
      title: {
        get() {
          return this.$store.state.book.title
        },
        set (value) {
          this.$store.commit('book/title', value)
        }
      }
    },
    methods: {
      save() {
        console.log('save');
      }
    }
  }
</script>
