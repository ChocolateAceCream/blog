<template>
  <el-aside :width="isNavBarCollapsed ? 'auto' : '220px'">
    <el-menu
      class="left-bar-menu"
      active-text-color="#3498db"
      :default-active="currentMenuIndex"
      :collapse="isNavBarCollapsed"
      unique-opened
      @select="onSelectMenuItem"
    >
      <div class="title">
        <img
          class="logo"
          src="/img/logo.jpg"
        >
        <div
          v-if="!isNavBarCollapsed"
          class="title-text"
        >BLOG</div>
      </div>
      <template
        v-for="item in routerTree"
        :key="item.name"
      >
        <menu-item
          :route-info="item"
        />
      </template>
    </el-menu>
  </el-aside>
</template>

<script>
import { defineComponent, toRefs, reactive, computed, getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'
import { navBarCollapsedHook } from '@/shared/hooks/index.js'
import _ from 'lodash'
import MenuItem from './menuItem.vue'
import { useRouterStore } from '@/stores/routerStore'
export default defineComponent({
  components: {
    MenuItem
  },
  setup(props, ctx) {
    const router = useRouter()
    const routerStore = useRouterStore()

    const { isNavBarCollapsed } = navBarCollapsedHook()
    console.log('----getCurrentInstance().appContext.components----', getCurrentInstance().appContext.components)

    const state = reactive({
      currentMenuIndex: computed(() => {
        const routeInMenu = _.findLast(
          router.currentRoute.value.matched,
          (item) => item.meta.isMenu
        )
        if (routeInMenu) {
          return routeInMenu.name
        } else {
          return router.currentRoute.value.name
        }
      }),
      routerTree: computed(() => {
        return routerStore.menuTree
      }),
      form: {},
    })

    const onSelectMenuItem = (routeName) => {
      router.push({name: routeName })
    }
    return {
      isNavBarCollapsed,
      onSelectMenuItem,
      ...toRefs(state),
    }
  },
})
</script>

<style lang="scss" scoped>
.left-bar-menu {
  background-color: $liter-background;
  height: 100%;
}
.title{
  min-height: 60px;
  line-height: 60px;
  text-align: center;
  transition: all .3s;
  display: flex;
  align-items: center;
  // justify-content: center;
  .title-text{
    display: inline-block;
    font-weight: 600;
    font-size: 20px;
    padding-left: 20px;
  }

  .logo {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    padding: 3px;
    margin-left:20px;
  }
}


</style>
