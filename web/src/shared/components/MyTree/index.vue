<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/04/11 10:21:49
* @description: customized el-tree component
!-->
<template>
  <el-input
    v-model="filterText"
    placeholder="Filter keyword"
  />
  <el-tree
    ref="treeRef"
    class="filter-tree"
    node-key="id"
    :props="defaultProps"
    v-bind="$attrs"
    default-expand-all
  />
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  props: {
    selected: {
      type: Array,
      default: () => []
    }
  },
  setup(props, ctx) {
    const state = reactive({
      treeRef: null,
      filterText: '',
      selectedMenus: [],
      defaultProps: {
        children: 'children',
        label: 'label',
      }
    })
    const setKeys = (keys) => {
      state.treeRef.setCheckedKeys(keys, false)
    }
    const getSelectKeys = () => {
      return [...state.treeRef.getHalfCheckedNodes(), ...state.treeRef.getCheckedNodes()]
    }
    return {
      setKeys,
      getSelectKeys,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
