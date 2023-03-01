import {computed} from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
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
