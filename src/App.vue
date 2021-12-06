<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <HelloWorld msg="Welcome to Your Vue.js + TypeScript App" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HelloWorld from "./components/HelloWorld.vue";
import Cat from "./core/intefaces/cat";

export default defineComponent({
  name: "App",
  components: {
    HelloWorld,
  },
  data() {
    return {
      cats: [] as Cat[],
    };
  },

  created() {
    console.log("here");
    fetch("http://localhost:3000/api/cats/v1/get", {
      method: "GET",
    }).then((res) => {
      res.json().then((r) => {
        this.cats = r.cats;
        console.log(this);
      });
    });
  },

  methods: {},
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
