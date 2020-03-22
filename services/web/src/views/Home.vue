<template>
  <v-container fluid class="container">
    <v-img alt="Vue" src="../assets/logo.png" height="200" contain />
    <v-row class="mx-auto my-12" justify="center">
      <h1 class="display-2">Welcome to Super Smash Heroes!</h1>
      <v-col cols="md-6">
        <v-btn class="my-8" color="primary" block x-large to="/battle"
          >Start Game</v-btn
        >
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Home extends Vue {}
</script>

<style scoped>
.container {
  padding: 100px 20px 20px 20px;
}
</style>
