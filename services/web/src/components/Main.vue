<template>
  <div class="main">
    <img alt="Vue" src="../assets/logo.png" />
    <h1>Welcome to Super Smash Heroes!</h1>
    <h3>Superheroes</h3>
    <ul>
      <li v-for="superhero in superheroes" v-bind:key="superhero.id">
        <p>{{ superhero.full_name }}</p>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

import { Superhero } from "../types";

@Component
export default class Main extends Vue {
  @Prop({ default: [] }) private superheroes!: Superhero[];
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li,
p {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
