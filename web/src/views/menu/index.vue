<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/02/23 10:18:13
* @description
!-->
<template>
  <div>menu</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
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
