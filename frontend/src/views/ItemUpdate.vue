<template>
  <div>
    <h1>ItemUpdate</h1>
    <form @submit.prevent="update">
      name:
      <input type="text" v-model="item.name" /><br>
      price:
      <input type="text" v-model="item.price" /><br>
      <button type="submit">update</button>
    </form>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "itemUpdate",
    created() {
      axios.get('http://localhost:8000/item/' + this.$route.params.id).then((res) => {
        this.item = res.data
      });
    },
    data() {
      return {
        item: ''
      }
    },
    methods: {
      update() {
        axios.patch('http://localhost:8000/item',
          {
            id: this.item.id,
            name: this.item.name,
            price: parseInt(this.item.price)
          }
        ).then(res => {
          console.log(res);
        })
      }
    }
  }
</script>
