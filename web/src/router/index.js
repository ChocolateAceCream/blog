import { createRouter, createWebHistory } from 'vue-router'
import { routeMiddleware } from './middleware'
import { auth } from './auth'
import Layout from '@/views/layout'
import { ElMessage } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
import { useRouterStore } from '@/stores/routerStore'
const routes = [
  {
    // path: '/', name: 'baseRoute', component: Layout, meta: { title: '主页', requireAuth: true },
    path: '/', name: 'baseRoute', redirect: '/home', component: Layout, meta: { title: '主页', requireAuth: true },
    // children: [
    //   // { path: '', name: 'home', component: () => import('@/views/home'), },
    //   { path: '/home', name: 'home', component: () => import('@/views/home'), meta: { title: '主页', requireAuth: true } },
    // ],

  },
  {
    path: '/404', name: '404', component: () => import('@/views/error/index.vue'), meta: { title: '404', requireAuth: false },
  },
  auth,
]
const wrappedRouter = routeMiddleware(routes)

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_APP_BASE_URL),
  routes: wrappedRouter
})

let flag = 0
router.beforeEach(async(to, from) => {
  const isAuthenticated = useSessionStore().isAuthenticated
  console.log('-----flag-------', flag)
  console.log('-----isAuthenticated-------', isAuthenticated)
  console.log('-----!isAuthenticated && to.meta.requireAuth-------', to)
  if (isAuthenticated) {
    if (from.name == null) {
      if (flag === 0) {
        const routerStore = useRouterStore()
        await routerStore.setAsyncRouter()
        flag++
        console.log('-----router-------', router)
        console.log('-----to.matched.length-------', to.matched.length)
        return { ...to, replace: true }
      }
      if (to.matched.length) {
        return true
      } else {
        router.push('/404')
      }
    } else {
    }
  } else {
    if (to.meta.requireAuth && to.matched.length > 0 || (to.matched.length === 0)) {
      ElMessage({
        message: '认证已过期，请重新登录',
        type: 'error',
        duration: 5 * 1000
      })

      return {
        path: '/auth/login',
        query: { redirect: to.fullPath },
      }
    }
  }
})

router.beforeResolve(async to => {
  return true
})

router.afterEach(async(to, from) => {
  const title = to.meta?.title
  document.title = 'Blog' + (title ? `-${title}` : '')
})


export default router
