<template>
  <v-container fluid class="container">
    <v-row>
      <v-col v-for="(player, index) in players" :key="player.alterEgo" class="pa-0">
        <h1 :class="`p${index + 1}-header`">P{{ index + 1 }} - {{ player.alterEgo }}</h1>
        <v-img :src="player.imageUrl" :alt="player.alterEgo" height="700"></v-img>
      </v-col>
    </v-row>
    <v-row>
      <v-col
        v-for="superhero in superheroes"
        :key="superhero.alterEgo"
        :class="playerSelectedClass(superhero)"
        class="pa-0"
        cols="1"
      >
        <a>
          <v-img
            v-on:click="playerSelect(superhero)"
            :src="superhero.imageUrl"
            :alt="superhero.alterEgo"
            aspect-ratio="1"
          />
        </a>
      </v-col>
    </v-row>
    <v-row>
      <v-col class="pa-0">
        <v-btn v-on:click="battle" :loading="loading" color="primary" x-large tile block>Battle!</v-btn>
      </v-col>
      <v-col class="pa-0">
        <v-btn color="secondary" x-large tile block>Winner: {{ getWinner() }}</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import axios from "axios";
import { Component, Vue } from "vue-property-decorator";

import { Superhero, PlayerTurn, Winner } from "../types";

const MAX_PLAYERS = Object.keys(PlayerTurn).length / 2;

@Component
export default class Battle extends Vue {
  superheroes: Superhero[] = [];
  players: Superhero[] = [];
  turn = PlayerTurn.PlayerOne;
  winner = Winner.None;
  loading = false;

  mounted() {
    axios.get("/api/superheroes").then(response => {
      this.superheroes = response.data.superheroes as Superhero[];
      this.players = this.superheroes.slice(0, 2);
    });
  }

  playerSelect(superhero: Superhero): void {
    if (this.winner) {
      this.winner = Winner.None;
    }

    if (!this.players.includes(superhero)) {
      Vue.set(this.players, this.turn, superhero);
      if (this.turn < MAX_PLAYERS - 1) {
        this.turn++;
      } else {
        this.turn = PlayerTurn.PlayerOne;
      }
    }
  }

  playerSelectedClass(superhero: Superhero): string {
    const index = this.players.findIndex(
      (player: Superhero) => player && player.id === superhero.id
    );
    if (index >= 0) {
      return `p${index + 1}-selected`;
    }
    return "";
  }

  battle() {
    const data = {
      playerOne: {
        intelligence: this.players[0].intelligence,
        strength: this.players[0].strength,
        speed: this.players[0].speed,
        durability: this.players[0].durability,
        power: this.players[0].power,
        combat: this.players[0].combat
      },
      playerTwo: {
        intelligence: this.players[1].intelligence,
        strength: this.players[1].strength,
        speed: this.players[1].speed,
        durability: this.players[1].durability,
        power: this.players[1].power,
        combat: this.players[1].combat
      }
    };

    this.loading = true;
    setTimeout(() => {
      axios.post("/api/battle", data).then(response => {
        this.winner = response.data.winner;
        this.loading = false;
      });
    }, 3000);
  }

  getWinner() {
    if (this.winner) {
      return `P${this.winner} - ${this.players[this.winner - 1].alterEgo}`;
    }
  }
}
</script>

<style scoped>
.container {
  padding: 0;
}
.p1-header {
  color: white;
  background-color: red;
  padding-top: 15px;
}
.p1-selected {
  border: 3px solid red;
}
.p2-header {
  color: white;
  background-color: blue;
  padding-top: 15px;
}
.p2-selected {
  border: 3px solid blue;
}
</style>
