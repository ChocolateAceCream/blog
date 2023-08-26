<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/07/18 10:38:05
* @description
sample usage
<my-search-bar @change="onSearch" />
!-->
<template>
  <el-input
    v-model="input"
    class="w-50 m-2"
    placeholder="Search"
    @change="onSearch"
  >
    <template #suffix>
      <SvgIcon
        :icon-name="'icon-blog-search'"
        size="15px"
        style="margin-right:10px;cursor:pointer;"
        @click="onSearch"
      />
    </template>

  </el-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  name: 'MySearchBar',
  emits: [ 'change'],
  setup(props, ctx) {
    const state = reactive({
      input: '',
    })
    const onSearch = () => {
      ctx.emit('change', state.input)
    }
    return {
      onSearch,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
