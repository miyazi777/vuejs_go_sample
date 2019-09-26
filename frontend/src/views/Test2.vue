<template>
  <div>
    <h1>Test Page2(指定されたURLパラメータに合致したデータを取得)</h1>
    <p>id : {{ $route.params.id }}</p>
    <p>{{ book.name }}</p>
    <p>{{ book.price }}</p>
  </div>
</template>

<script>
  import _ from 'lodash';

  export default {
    name: "test2",
    created() {
      this.$store.dispatch('test2/getBook', this.$route.params.id).then(() => {
        this.book = _.cloneDeep(this.$store.state.test2.book)
      });
    },
    data() {
      return {
        book: ''
      }
    }
  }
</script>
