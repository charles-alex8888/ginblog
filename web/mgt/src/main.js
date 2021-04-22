import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import router from 'vue-router'
import './assets/css/style.css'

import axios from 'axios'
import './plugin/antui'
// import { config } from 'vue/types/umd'
axios.defaults.baseURL = 'http://localhost:3000/api/v1'

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
