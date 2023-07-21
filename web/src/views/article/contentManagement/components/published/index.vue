<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/07/18 16:54:38
* @description: display published article list
!-->
<template>
  <div>published</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  name: 'Published',
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
</style>
