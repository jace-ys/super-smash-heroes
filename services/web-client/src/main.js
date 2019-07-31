import Vue from 'vue';
import axios from 'axios';

Vue.config.productionTip = false

axios.get("/api/superheroes/1").then(response => console.log(response.data));
