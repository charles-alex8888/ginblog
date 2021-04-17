import Vue from 'vue'
import VueRouter from 'vue-router'
import Router from 'vue-router'

import Login from "../views/Login.vue"
import Mgt from "../views/Mgt.vue"

// 页面路由组件
import Index from "../components/mgt/Index.vue"
import AddArt from "../components/article/AddArt.vue"
import ArtList from "../components/article/ArtList.vue"
import CateList from "../components/category/CateList.vue"
import UserList from "../components/user/UserList.vue"




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
    component: Mgt,
    children: [
      { path: '/index', component: Index},
      { path: '/addart', component: AddArt},
      { path: '/artlist', component: ArtList},
      { path: '/catalist', component: CateList},
      { path: '/userlist', component: UserList},
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to,from,next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path == '/login') return next()
  if (!token && to.path == '/mgt') {
    next('/login')
  }else {
    next()
  }
})

export default router
