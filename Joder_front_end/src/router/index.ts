
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
    }
  ]
})

/*之後再修
router.beforeEach((to: RouteLocationNormalized) => {
  if (userStore === null){
     userStore = useLoginStore()
  }
  
  console.log(userStore.isLogin)
  if (!userStore.isLogin && to.path != '/login') {
    return { name: 'LoginPage' }
  }
  console.log(userStore.isLogin)
})*/ 


export default router
