<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/07/18 10:15:27
* @description
!-->
<template>
  <div style="position: relative;">
    <el-tabs
      v-model="activeTab"
      class="demo-tabs"
      @tab-click="onTabSwitch"
    >
      <el-tab-pane
        label="Drafts"
        name="drafts"
      >
        <draft />
      </el-tab-pane>
      <el-tab-pane
        label="Published"
        name="published"
      >
        <published />
      </el-tab-pane>
    </el-tabs>

  </div></template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import Draft from './components/draft/index.vue'
import Published from './components/published/index.vue'
export default defineComponent({
  components: { Draft, Published },
  setup(props, ctx) {
    const state = reactive({
      activeTab: 'drafts',
    })
    const onTabSwitch = (tab, e) => {
      console.log(tab, e)
    }

    return {
      onTabSwitch,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
