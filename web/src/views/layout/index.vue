<template>
  <el-container class="layout-wrapper">
    <el-header style="border-bottom: 1px solid #DCDCDC;">
      <top-bar />
    </el-header>
    <el-container>
      <left-bar />
      <!-- <el-header style="--el-header-height: 30px"><top-bar></top-bar></el-header> -->
      <el-main>
        <main-body />
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import LeftBar from './components/LeftBar.vue'
import TopBar from './components/TopBar.vue'
// import TagsView from './TagsView.vue'
import MainBody from './components/MainBody.vue'
// import { getProfile } from '@/api/session'
import { reactive, toRefs, defineComponent } from 'vue'

export default defineComponent({
  components: {
    LeftBar, TopBar, MainBody
  },
  setup(props, ctx) {
    const state = reactive({})
    return { ...toRefs(state) }
  }
})
</script>

<style lang="scss" scoped>
.layout-wrapper{
  .el-header{
    --el-header-height: 60px;
    --el-header-width: 100%;
  }
  height:100%;
}
</style>
