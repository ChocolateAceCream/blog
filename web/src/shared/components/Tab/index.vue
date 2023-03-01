<!--
* @fileName index.vue
* @author Di Sheng
* @date 2022/12/29 19:16:04
* @description
!-->
<template>
  <el-tabs
    v-model="activeTab"
    class="demo-tabs"
    @tab-click="handleClick"
  >
    <template
      v-for="(tab, idx) in config.tabs"
      :key="idx"
    >
      <el-tab-pane
        :label="tab.label"
        :name="tab.name"
      >
        <component :is="tab.component" />
      </el-tab-pane>
    </template>

  </el-tabs></template>

<script>
import { defineComponent, toRefs, reactive, ref } from 'vue'
export default defineComponent({
  props: {
    config: [Object],
  },
  emits: [ 'change'],
  setup(props, ctx) {
    const handleClick = (tab, event) => {
      ctx.emit('change', tab, event)
    }
    const state = reactive({
      activeTab: props.config?.activeTab || props.config?.tabs[0]?.name,
    })
    return {
      handleClick,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.demo-tabs>.el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
