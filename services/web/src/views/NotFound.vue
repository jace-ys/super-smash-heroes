<template>
  <v-container fluid class="container">
    <h1 class="title">Page Not Found</h1>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class NotFound extends Vue {}
</script>

<style scoped>
.container {
  padding: 20px;
}
</style>
