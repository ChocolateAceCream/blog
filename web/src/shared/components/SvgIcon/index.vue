<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/03/09 08:59:49
* @description svg icon
!-->
<template>
  <el-icon
    :color="color"
    class="icon"
  >
    <span :class="`iconfont ` + iconName" />
  </el-icon>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  props: {
    iconName: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: ''
    },
    size: {
      type: String,
      default: '20px'
    }
  },
  setup(props, ctx) {
    const state = reactive({
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.icon{
  width: v-bind('size');
  height: v-bind('size');
}
.iconfont{
  font-size: v-bind('size');
}
</style>
