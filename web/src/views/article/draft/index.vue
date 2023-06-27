<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/20 11:14:22
* @description drafting new article using md-v3 editor
!-->
<template>
  <div>asdfasdf</div>
  <MyEditor
    v-model="content"
    class="md-editor"
  />
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  setup(props, ctx) {
    const state = reactive({
      content: '',
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
