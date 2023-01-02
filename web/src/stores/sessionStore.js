import { defineStore, acceptHMRUpdate } from 'pinia'

export const sessionStore = defineStore({
  id: 'sessionStore',
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
    doubleCount: (state) => state.counter * 2,
    isAuthenticated: (state) => state.userInfo.isAuthenticated
  },
  actions: {
    increment() {
      this.counter++
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
  import.meta.hot.accept(acceptHMRUpdate(sessionStore, import.meta.hot))
}
