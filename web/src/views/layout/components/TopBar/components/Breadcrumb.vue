<template>
  <div
    class="icon-wrapper"
    @click="toggleNavBar"
  >
    <SvgIcon
      v-if="isNavBarCollapsed"
      icon-name="icon-blog-double-arrow-right"
      size="15px"
    />
    <SvgIcon
      v-if="!isNavBarCollapsed"
      icon-name="icon-blog-double-arrow-left"
      size="15px"
    />
  </div>
  <el-breadcrumb
    separator="/"
    class="breadcrumb"
  >
    <el-breadcrumb-item
      v-for="(item, index) in currentRouter"
      :key="index"
    >
      <span class="item">{{ item?.meta.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
import { reactive, toRefs, defineComponent, computed} from 'vue'
import { useRoute } from 'vue-router'
import { useRouterStore } from '@/stores/routerStore'
import { navBarCollapsedHook } from '@/shared/hooks/index.js'
export default defineComponent({
  setup(props, ctx) {
    const router = useRoute()
    const routerStore = useRouterStore()
    const { isNavBarCollapsed, toggleNavBar } = navBarCollapsedHook()
    const state = reactive({
      currentRouter: computed(() => {
        const arr = []
        let currentRoute = routerStore.routerList.find(routerItem => { return routerItem.name === router.name })
        if (currentRoute !== undefined) {
          while (currentRoute?.pid !== 0) {
            arr.unshift(currentRoute)
            currentRoute = routerStore.routerList.find(routerItem => { return routerItem.id === currentRoute.pid })
          }
          arr.unshift(currentRoute)
        }
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
  line-height: 30px;
}
.icon-wrapper{
  position:absolute;
  line-height: 30px;
}
.breadcrumb{
  margin-left: 25px;
}
</style>
