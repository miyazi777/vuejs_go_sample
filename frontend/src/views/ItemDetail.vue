<template>
  <div>
    <h1>Item Detail</h1>
    <p>Id : {{ this.item.id }}</p>
    <p>Name : {{ this.item.name }}</p>
    <p>Price : {{ this.item.price }}</p>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "itemDetail",
    created() {
      axios.get('http://localhost:8000/item/' + this.$route.params.id).then((res) => {
        this.item = res.data
      });
    },
    data() {
      return {
        item: ''
      }
    }
  }
</script>
