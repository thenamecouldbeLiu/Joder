import Cookies from 'js-cookie'
import { useLoginStore } from './../stores/loginStore'
import { createRouter, createWebHistory, type RouteLocationNormalized } from 'vue-router'
let userStore: any = null
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'HomePage',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/login',
      name: 'LoginPage',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/signUp',
      name: 'SignUp',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/SignUpView.vue')
    }
  ]
})

router.beforeEach((to: RouteLocationNormalized) => {
  if (userStore === null) {
    userStore = useLoginStore()
  }

  if (!Cookies.get('joder_token') && to.path != '/login' && to.path != '/signUp') {
    return { name: 'LoginPage' }
  }
})

export default router
