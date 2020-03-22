<template>
  <v-app id="app">
    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script lang="ts">
import axios from "axios";
import { Component, Vue } from "vue-property-decorator";

import { Superhero } from "./types";

@Component
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
}
</style>
