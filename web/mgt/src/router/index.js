import Vue from 'vue'
import VueRouter from 'vue-router'
import Router from 'vue-router'

import Login from "../views/Login.vue"
import Mgt from "../views/Mgt.vue"

Vue.use(VueRouter)
Vue.use(Router)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/mgt',
    name: 'mgt',
    component: Mgt
  }
]

const router = new VueRouter({
  routes
})

export default router
