import { defineStore, acceptHMRUpdate } from 'pinia'
import { userRouterStore } from './routerStore'
export const useSessionStore = defineStore({
  id: 'useSessionStore',
  state: () => ({
    counter: 0,
    name: 'di',
    userInfo: {
      account: '',
      email: '',
      id: null,
      telno: '',
      status: null,
      isAdmin: false,
      username: '',
      locale: 'cn',
      isAuthenticated: false,
    }
  }),
  getters: {
    isAuthenticated: (state) => state.userInfo.isAuthenticated
  },
  actions: {
    increment() {
      this.counter++
    },
    setUserInfo(info) {
      Object.keys(info).map(key => {
        this.userInfo[key] = info[key]
      })
      this.userInfo.isAuthenticated = true
    },
    setPermissionList(permissions) {

    },
    logout() {
      this.userInfo = {
        account: '',
        email: '',
        id: null,
        telno: '',
        status: null,
        isAdmin: false,
        username: '',
        locale: 'cn',
        isAuthenticated: false,
      }
      const routerStore = userRouterStore()
      routerStore.$reset()
      console.log('---routerStore.asyncRouterFlag---', routerStore.asyncRouterFlag)
    },
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
  import.meta.hot.accept(acceptHMRUpdate(useSessionStore, import.meta.hot))
}
