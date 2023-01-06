import { createRouter, createWebHistory } from 'vue-router'
import { routeMiddleware } from './middleware'
import { auth } from './auth'
import Layout from '@/views/layout'
import { ElMessage } from 'element-plus'
import { sessionStore } from '@/stores/sessionStore'
const routes = [
  {
    path: '/', component: Layout, redirect: '/home', meta: { title: '主页', requireAuth: true },
    children: [
      // { path: '', name: 'home', component: () => import('@/views/home'), },
      { path: '/home', name: 'home', component: () => import('@/views/home'), meta: { title: '主页', requireAuth: true } },
    ]
  },
  auth,
]
const wrappedRouter = routeMiddleware(routes)

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_APP_BASE_URL),
  routes: wrappedRouter
})


router.beforeEach(async(to, from) => {
  const isAuthenticated = sessionStore().isAuthenticated
  console.log('-----isAuthenticated-------', isAuthenticated)
  console.log('-----!isAuthenticated && to.meta.requireAuth-------', to)
  if (isAuthenticated) {
    if (to.name.includes('login')) {
      return { path: to.query?.redirect || '/' }
    }
  }
  if (!isAuthenticated && to.meta.requireAuth) {
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
})

router.beforeResolve(async to => {
  return true
})

router.afterEach(async(to, from) => {
  const title = to.meta?.title
  document.title = 'Blog' + (title ? `-${title}` : '')
})


export default router
