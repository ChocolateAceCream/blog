import {computed} from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
import { useRouterStore } from '@/stores/routerStore'
import router from '@/router'
export const navBarCollapsedHook = () => {
  const sessionStore = useSessionStore()
  const isNavBarCollapsed = computed({
    get: () => { return sessionStore.userSetting.isNavBarCollapsed },
    set: () => {
      sessionStore.userSetting.isNavBarCollapsed = !sessionStore.userSetting.isNavBarCollapsed
    }
  })
  const toggleNavBar = () => {
    isNavBarCollapsed.value = false
    console.log('---isNavBarCollapsed--', sessionStore.userSetting.isNavBarCollapsed)
  }
  return {
    isNavBarCollapsed,
    toggleNavBar
  }
}

export const logout = () => {
  router.push({ name: 'login' })
  const routerStore = useRouterStore()
  routerStore.$reset()
  const sessionStore = useSessionStore()
  sessionStore.closeNotificationWebsocket()
  sessionStore.$reset()
  console.log('jump to login', router)
}
