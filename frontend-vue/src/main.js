import Vue from 'vue'
import App from './App.vue'

import { routes } from './routes';

import VueRouter from 'vue-router';
import { BootstrapVue } from 'bootstrap-vue';

Vue.config.productionTip = false


// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue)
Vue.use(VueRouter)

const router = new VueRouter({
  routes
});

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
