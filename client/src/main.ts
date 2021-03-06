import Vue from 'vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import router from './router'
import App from './App.vue'
import store from '@/store'
const vueConfig = require('vue-config');

const isDev = process.env.NODE_ENV === 'development';

let key = process.env.VUE_APP_TWITTER_KEY;
try {
  key = require('../../token/auth.json').twitter;
} catch (e) {
  console.log(e);
}
key = key.split('.');

Vue.use(Vuetify);
Vue.use(vueConfig, {
  API: process.env.API_ENDPOINT,
  isDev,
  providers: {
    auth: {
      twitter: {
        clientId: key[0],
        clientSecret: key[1],
      }
    }
  },
});

new Vue({
  router,
  store,
  render: function (h) { return h(App) }
}).$mount('#app');
