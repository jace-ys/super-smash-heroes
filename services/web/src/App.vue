<template>
  <div id="app">
    <Main v-bind:superheroes="superheroes" />
  </div>
</template>

<script lang="ts">
import axios from "axios";
import { Component, Vue } from "vue-property-decorator";

import Main from "./components/Main.vue";
import { Superhero } from "./types";

@Component({
  components: {
    Main
  }
})
export default class App extends Vue {
  superheroes: Superhero[] = [];

  mounted() {
    axios.get("/api/superheroes").then(response => {
      console.log(response.data);
      this.superheroes = response.data.superheroes as Superhero[];
    });
  }
}
</script>

<style>
html,
body {
  height: 100vh;
  width: 100vw;
  margin: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding: 20px;
  margin-top: 60px;
}
</style>
