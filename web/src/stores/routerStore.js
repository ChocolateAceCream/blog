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
    pid: root.pid,
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

const formatMenuTree = (nodes) => {
  const r = []
  nodes.forEach(node => {
    const temp = {
      routeName: node.name,
      title: node.meta.title,
      icon: node.meta.icon,
      id: node.id,
      children: []
    }
    if (node.children.length > 0) {
      temp.children = formatMenuTree(node.children)
    }
    r.push(temp)
  })
  return r
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



export const useRouterStore = defineStore({
  id: 'useRouterStore',
  state: () => ({
    asyncRouterFlag: 0, // prevent duplicated get router action
    routerList: [],
    routerTree: [],
    asyncRouters: [],
    menuTree: [],
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
          message: asyncRouter.msg,
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
      this.menuTree = formatMenuTree(this.routerTree)
      console.log('----routerTree---', this.routerTree)
      console.log('----menuTree---', this.menuTree)
      const pushRoutes = (nodes) => {
        nodes.forEach(node => {
          if (node.children.length > 0) {
            pushRoutes(node.children)
          } else {
            router.addRoute('baseRoute', node)
          }
        })
      }
      pushRoutes(this.routerTree)
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
  import.meta.hot.accept(acceptHMRUpdate(useRouterStore, import.meta.hot))
}
