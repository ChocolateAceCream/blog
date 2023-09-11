<!--
* @fileName Notification.vue
* @author Di Sheng
* @date 2023/09/09 08:51:33
* @description notification dropdown menu
!-->
<template>
  <el-dropdown
    class="avatar-container right-menu-item"
    trigger="click"
  >
    <div>
      <el-badge
        :value="200"
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
    <template #dropdown>
      <el-dropdown-menu />
    </template>
  </el-dropdown>
</template>

<script>
import { defineComponent, toRefs, reactive, watch } from 'vue'
import useWebsocket from '@/utils/websocket.js'
export default defineComponent({
  setup(props, ctx) {
    const onMessage = (ws, e) => {
      console.log('----------ws-------', ws)
      console.log(e)
    }
    const url = `${import.meta.env.VITE_WEBSOCKET_PROXY}/api/v1/notification/ws`
    console.log('--------url-----', url)
    const { status, data, close, open } = useWebsocket(url, {
      autoReconnect: true,
      onMessage: onMessage
    })
    open()
    watch(status, val => {
      console.log('----status-----', status)
    })
    const state = reactive({
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
