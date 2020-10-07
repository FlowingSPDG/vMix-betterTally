import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Tally from '../views/Tally.vue'
import api from '../utils/api.vue'

Vue.use(VueRouter)

// Mixin ./utils/api.vue
Vue.mixin(api)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/tally',
    name: 'Tally',
    component: Tally
  }
]

const router = new VueRouter({
  routes
})

export default router
