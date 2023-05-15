<template>
  <el-row class="header">
    <!--右侧-->
    <el-col
      :span="12"
      class="left-header"
    >
      <breadcrumb />
    </el-col>

    <el-col
      :span="12"
      class="right-header"
    >
      <el-dropdown
        class="avatar-container right-menu-item"
        trigger="click"
      >
        <SvgIcon
          icon-name="icon-blog-setting"
          color="#3498db"
          class="top-bar-setting-icon"
          size="30px"
        />
        <template #dropdown>
          <el-dropdown-menu>
            <router-link :to="{ name: 'resetPassword' }">
              <el-dropdown-item>
                密码重置
              </el-dropdown-item>
            </router-link>
            <el-dropdown-item
              divided
              @click.stop="handleLogout"
            >
              退出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-col>
  </el-row>
</template>
<script>
import { reactive, toRefs, defineComponent } from 'vue'
import { ElMessageBox } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
import { logout } from '@/shared/hooks/index'
import Breadcrumb from './Breadcrumb.vue'
import _ from 'lodash'
export default defineComponent({
  components: {
    Breadcrumb
  },
  setup(props, ctx) {
    const store = useSessionStore()

    const state = reactive({
      username: store.userInfo.username,
      async handleLogout() {
        try {
          await ElMessageBox.confirm(
            '确认退出登录？',
            '提示',
            {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning',
            }
          )
          logout()
        } catch (err) {
          console.log('-------cancel login----', err)
        }
      }
    })
    return { ...toRefs(state) }
  }
})
</script>
<style module lang="scss" src="@/assets/styles/export.scss"></style>
<style lang="scss" scoped>
.header{
  align-items: center;
  height: 100%;
}
  .content {
    display: inline-block;
    @include mobile-device {
      display: none;
    }
  }

  .right-header {
    display: flex;
    flex-direction: row-reverse;
    text-align: right;
    .top-bar-setting-icon {
      cursor: pointer;
    }
  }
  .left-header{
    display: flex;
    align-items: center;
  }
</style>
