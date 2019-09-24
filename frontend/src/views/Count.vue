<template>
  <div>
    <h1>Count Page</h1>
    {{$store.state.counter.count}}
    <button @click="increment">+</button>
    <button @click="decrement">-</button>
  </div>
</template>

<script>
  export default {
    name: "counter",
    methods: {
      increment() {
        this.$store.commit("counter/increment");
      },
      decrement() {
        this.$store.commit("counter/decrement");
      }
    }
  }
</script>
