<template>
  <el-col
    :span="1"
    @click="toggleNavBar"
  >
    <i-svg-right-double-arrow
      v-if="isNavBarCollapsed"
      class="double-arrow"
    />
    <i-svg-left-double-arrow
      v-if="!isNavBarCollapsed"
      class="double-arrow"
    />
  </el-col>
  <el-col :span="10">
    <el-breadcrumb
      separator="/"
      class="breadcrumb"
    >
      <el-breadcrumb-item
        v-for="(item, index) in currentRouter"
        :key="index"
      >
        <span class="item">{{ item.meta.title }}</span>
      </el-breadcrumb-item>
    </el-breadcrumb>
  </el-col>
</template>

<script>
import _ from 'lodash'
import { reactive, toRefs, defineComponent, computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useRouterStore } from '@/stores/routerStore'
import { navBarCollapsedHook } from '@/shared/hooks/navBarCollapsed'
export default defineComponent({
  setup(props, ctx) {
    const router = useRoute()
    const routerStore = useRouterStore()
    const { isNavBarCollapsed, toggleNavBar } = navBarCollapsedHook()
    console.log('-----router-----', router.name)
    const state = reactive({
      currentRouter: computed(() => {
        const arr = []
        let currentRoute = routerStore.routerList.find(routerItem => { return routerItem.name === router.name })
        console.log('-----currentRoute-----', currentRoute)
        while (currentRoute.pid !== 0) {
          arr.unshift(currentRoute)
          currentRoute = routerStore.routerList.find(routerItem => { return routerItem.id === currentRoute.pid })
        }
        arr.unshift(currentRoute)
        console.log('----ar-----', arr)
        return arr
      })
    })

    return {
      toggleNavBar,
      isNavBarCollapsed,
      ...toRefs(state)
    }
  }
})
</script>

<style lang="scss" scoped>
.item {
  line-height: 31px;
}
.double-arrow{
  height:100%;
  width:20px;
  margin:auto;
}
</style>
