<!--
* @fileName Notification.vue
* @author Di Sheng
* @date 2023/09/09 08:51:33
* @description notification dropdown menu
!-->
<template>
  <router-link :to="{ name: 'notification' }">
    <div
      class="avatar-container right-menu-item"
      trigger="click"
    >
      <el-badge
        :hidden="notificationCount === 0"
        :value="notificationCount"
        :max="99"
        class="item"
      >
        <SvgIcon
          icon-name="icon-blog-notification"
          color="#3498db"
          class="top-bar-setting-icon"
          size="30px"
        />
      </el-badge>
    </div>
  </router-link>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  setup(props, ctx) {
    const sessionStore = useSessionStore()
    const state = reactive({
      notificationCount: computed(() => sessionStore.newNotificationCount)
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.top-bar-setting-icon{
  cursor: pointer;
  margin-left:25px;
}
</style>
