<template>
  <div>
    <h1>Test Page(画面表示時にAPIで値を取得し、dataに設定するサンプル)</h1>
    <p>{{ message }}</p>
  </div>
</template>

<script>
  import _ from 'lodash';

  export default {
    name: "test",
    created() {
      this.$store.dispatch('test/getMsg').then(() => {
        this.message = _.cloneDeep(this.$store.state.test.msg)
      });
    },
    data() {
      return {
        message: ''
      }
    },
    //computed: {
    //  msg: function() {
    //    return this.$store.getters['test/msg'];
    //  }
    //},
    //methods: {
    //  getMsg() {
    //    return this.$store.dispatch('test/getMsg').then(() => {
    //      //this.message = this.$store.state.test.msg;
    //      this.message = _.cloneDeep(this.$store.state.test.msg)
    //    });
    //  }
    //}
  }
</script>
