import { ElMessage } from 'element-plus'
import { defineStore, acceptHMRUpdate } from 'pinia'
import { getCurrentUserMenu } from '@/api/menu'
import router from '@/router'

const modules = import.meta.glob('../views/**/*.vue')

const routerListArr = []
const treeHelper = (root, mapper) => {
  const temp = {
    path: root.path,
    name: root.name,
    component: modules[`../${root.component}`],
    id: root.id,
    meta: root.meta,
    children: []
  }
  if (mapper[temp.id]) {
    mapper[temp.id].forEach(node => {
      const child = treeHelper(node, mapper)
      temp.children.push(child)
    })
  }
  routerListArr.push(temp)
  return temp
}
const formatRouterTree = (data) => {
  const mapper = {}
  data.forEach(node => {
    if (mapper[node.pid]) {
      mapper[node.pid].push(node)
    } else {
      mapper[node.pid] = [node]
    }
  })
  let r = []
  if (mapper[0] && mapper[0].length > 0) {
    r = mapper[0].map(node => {
      return treeHelper(node, mapper)
    })
  }
  return r
}

export const userRouterStore = defineStore({
  id: 'userRouterStore',
  state: () => ({
    asyncRouterFlag: 0, // prevent duplicated get router action
    routerList: [],
    routerTree: [],
    asyncRouters: [],
  }),

  actions: {
    async getRouters() {
      const result = await getCurrentUserMenu()
      const asyncRouter = result.data
      if (asyncRouter.errorCode === 0) {
        this.asyncRouters = asyncRouter.data
        this.asyncRouterFlag++
      } else {
        ElMessage({
          message: result.msg,
          type: 'error',
          duration: 5 * 1000
        })
      }
    },
    async setAsyncRouter() {
      if (this.asyncRouterFlag === 0) {
        await this.getRouters()
      }
      this.routerTree = formatRouterTree(this.asyncRouters)
      this.routerList = routerListArr
      console.log('----routerTree---', this.routerTree.value)
      this.routerTree.forEach(route => router.addRoute('baseRoute', route))
      console.log('----router---', router)
      return true
    }
  },
  // 开启数据缓存
  persist: {
    enabled: true,
    strategies: [
      {
        storage: localStorage,
        // paths: ['name', 'counter']
      }
    ]
  }
})

// enable module hot reload
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(userRouterStore, import.meta.hot))
}
