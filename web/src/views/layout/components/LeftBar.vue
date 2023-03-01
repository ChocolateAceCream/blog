<template>
  <el-aside :width="isNavBarCollapsed ? 'auto' : '220px'">
    <el-menu
      class="left-bar-menu"
      active-text-color="#fff"
      :default-active="currentMenuIndex"
      :collapse="isNavBarCollapsed"
    >
      <router-link :to="{ name: 'home' }">
        <el-menu-item index="home">
          <template #title>
            <i-svg-home
              class="left-menu-icon"
            />
            <span>首页</span>
          </template>
        </el-menu-item>
      </router-link>
      <el-sub-menu index="device">
        <template #title>
          <svg-icon
            class="left-menu-icon"
            iconname="icon-shebeiguanli"
          />
          <span>设备管理</span>
        </template>
      <!-- <router-link :to="{ name: 'importDevice' }">
        <el-menu-item index="importDevice">AP设备入库</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'manufacturers' }">
        <el-menu-item index="manufacturers">厂家型号管理</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'deviceSetting' }">
        <el-menu-item index="deviceSetting">AP设备配置</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'groupList' }">
        <el-menu-item index="groupList">AP组列表</el-menu-item>
      </router-link> -->
      </el-sub-menu>
      <el-sub-menu index="3">
        <template #title>
          <svg-icon
            class="left-menu-icon"
            iconname="icon-xitongpeizhi"
          />
          <span>系统管理</span>
        </template>
      <!-- <router-link :to="{ name: 'menu' }">
        <el-menu-item index="menu">菜单管理</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'role' }">
        <el-menu-item index="role">角色管理</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'user' }">
        <el-menu-item index="user">账号管理</el-menu-item>
      </router-link>
      <router-link :to="{ name: 'logs' }">
        <el-menu-item index="logs">操作日志</el-menu-item>
      </router-link> -->
      </el-sub-menu>
      <el-menu-item
        index="4"
        disabled
      />
    </el-menu>
  </el-aside></template>

<script>
import { defineComponent, toRefs, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { navBarCollapsedHook } from '@/shared/hooks/navBarCollapsed'
import _ from 'lodash'
export default defineComponent({
  setup(props, ctx) {
    const router = useRouter()
    const { isNavBarCollapsed } = navBarCollapsedHook()
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
      form: {},
    })
    return {
      isNavBarCollapsed,
      ...toRefs(state),
    }
  },
})
</script>

<style lang="scss" scoped>
.left-menu-icon {
  margin-right: 8px;
  fill: $shadow-color;
  cursor: pointer;
  height: 20px;
  width: 20px;
}
.left-bar-menu {
  // color: $dark-brown;
  background-color: $lite-background;
  height: 100%;
  padding: 5px;
}
.el-menu-item {
  background-color: $lite-background;
}
.el-menu-item.is-active {
  background-color: $blue !important; //你要修改的颜色
  border-radius: 4px;
}
</style>
