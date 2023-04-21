<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/04/11 10:21:49
* @description: customized el-tree component
!-->
<template>
  <!-- TODO: add filter method -->
  <!-- <el-input
    v-model="filterText"
    placeholder="Filter keyword"
  /> -->
  <el-tree
    ref="treeRef"
    class="filter-tree"
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

    // return all selected nodes, including half check ones
    const getSelectKeys = (param = false) => {
      return [...state.treeRef.getHalfCheckedNodes(param), ...state.treeRef.getCheckedNodes(param)]
    }
    const getCheckedNodes = (param = false) => {
      return state.treeRef.getCheckedNodes(param)
    }
    return {
      setKeys,
      getSelectKeys,
      getCheckedNodes,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
