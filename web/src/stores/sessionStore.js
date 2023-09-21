import { defineStore, acceptHMRUpdate } from 'pinia'
import { ref } from 'vue'
import useWebsocket from '@/utils/websocket.js'
import { getUnreadNotificationCount } from '@/api/notification'
export const useSessionStore = defineStore({
  id: 'useSessionStore',
  state: () => ({
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
      permissions: [],
    },
    userSetting: {
      isNavBarCollapsed: false,
    },
    currentEditingArticle: null,
    closeNotificationWebsocket: null,
    newNotificationCount: 0,
  }),
  getters: {
    isAuthenticated: (state) => state.userInfo.isAuthenticated,
    getLocale: (state) => state.userInfo.locale,
  },
  actions: {
    setUserInfo(info) {
      Object.keys(info).map(key => {
        this.userInfo[key] = info[key]
      })
      this.userInfo.isAuthenticated = true
    },
    setNotificationWebsocket() {
      const url = `${import.meta.env.VITE_WEBSOCKET_PROXY}/websocket/api/v1/notification/ws`
      const onMessage = (ws, e) => {
        if (e.data === 'New Notification') {
          this.updateNotificationCount()
        }
      }
      // const { status, data, close, open } = useWebsocket(url, {
      const { status, data, close, open } = useWebsocket(url, {
        autoReconnect: true,
        onMessage: onMessage
      })
      open()
      this.closeNotificationWebsocket = close
      this.updateNotificationCount()
    },

    async updateNotificationCount() {
      const { data: res } = await getUnreadNotificationCount()
      if (res.errorCode === 0) {
        this.newNotificationCount = res.data
      }
    },

    hasPermission(name) {
      return !!this.userInfo.permissions[name]
    },
    setLocale(locale) {
      this.userInfo.locale = locale
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
  import.meta.hot.accept(acceptHMRUpdate(useSessionStore, import.meta.hot))
}
