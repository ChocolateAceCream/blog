import { createRouter, createWebHistory } from 'vue-router'
import { routeMiddleware } from './middleware'
import { auth } from './auth'
// import Layout from '@/views/layout'
import { ElMessage } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
import { useRouterStore } from '@/stores/routerStore'
export const Layout = () => import('@/views/layout')
const routes = [
  {
    // path: '/', name: 'baseRoute', component: Layout, meta: { title: '主页', requireAuth: true },
    path: '/', name: 'baseRoute', redirect: '/home', component: Layout, meta: { title: 'Home Page', requireAuth: true },
    // children: [
    //   // { path: '', name: 'home', component: () => import('@/views/home'), },
    //   { path: '/home', name: 'home', component: () => import('@/views/home'), meta: { title: '主页', requireAuth: true } },
    // ],

  },
  {
    path: '/:pathMatch(.*)*', component: () => import('@/views/error/index.vue'), meta: { title: '404', requireAuth: false },
    // path: '/:pathMatch(.*)*', name: '404', component: () => import('@/views/error/index.vue'), meta: { title: '404', requireAuth: false },
    // dont include name: '404', otherwise refresh page will end up in 404 page
  },
  auth,
]
const wrappedRouter = routeMiddleware(routes)

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_APP_BASE_URL),
  routes: wrappedRouter
})

let flag = 0
router.beforeEach(async(to, from, next) => {
  async function init() {
    const routerStore = useRouterStore()
    await routerStore.setAsyncRouter()
    flag = 1
  }
  const isAuthenticated = useSessionStore().isAuthenticated
  if (isAuthenticated) {
    if (to.path === '/auth/login') {
      next({path: '/'})
    }
    if (flag) {
      next()
    } else {
      await init()
      next({...to, replace: true})
    }
  } else {
    if (to.meta.requireAuth && to.matched.length > 0 || (to.matched.length === 0)) {
      ElMessage({
        message: 'Please Login',
        type: 'error',
        duration: 5 * 1000
      })
      next({
        path: '/auth/login',
        query: { redirect: to.fullPath },
      })
    }
    next()
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
