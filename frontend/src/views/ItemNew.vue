<template>
  <div>
    <h1>ItemNew</h1>
    <form @submit.prevent="save">
      name:
      <input type="text" v-model="item.name" /><br>
      price:
      <input type="text" v-model="item.price" /><br>
      <button type="submit">submit</button>
    </form>
  </div>
</template>

<script>
  import axios from 'axios';

  export default {
    name: "itemNew",
    data() {
      return {
        item: {}
      }
    },
    methods: {
      save() {
        axios.post('http://localhost:8000/item',
          {
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
