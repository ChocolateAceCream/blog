<template>
  <el-row>
    <!--右侧-->
    <el-col
      :span="24"
      class="right-menu"
    >
      <el-dropdown
        class="avatar-container right-menu-item"
        trigger="click"
      >
        <i-svg-setting
          class="top-bar-setting-icon"
        />
        <template #dropdown>
          <el-dropdown-menu>
            <!-- <router-link :to="{ name: 'resetUserInfo' }">
              <el-dropdown-item>
                信息修改
              </el-dropdown-item>
            </router-link>
            <router-link :to="{ name: 'resetPhone' }">
              <el-dropdown-item>
                手机号修改
              </el-dropdown-item>
            </router-link>-->
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
  <el-row style="height: 40px">
    <!-- <el-col :span="20">
      <breadcrumb class="breadcrumb" />
    </el-col> -->
  </el-row>
</template>
<script>
import { useRoute, useRouter } from 'vue-router'
import { reactive, toRefs, defineComponent, createVNode, computed } from 'vue'
import { ElMessageBox } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
import { userRouterStore } from '@/stores/routerStore'
import Breadcrumb from './Breadcrumb.vue'
import _ from 'lodash'
export default defineComponent({
  components: {
    Breadcrumb
  },

  setup(props, ctx) {
    const router = useRouter()
    const store = useSessionStore()

    const state = reactive({
      username: store.userInfo.username,
      async handleLogout() {
        console.log('-------handleLogout-----')
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
          router.push({name: 'login'})
          store.logout()
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
  .content {
    display: inline-block;
    .breadcrumb {
      padding-left: 10px;
      line-height: 30px;
    }
    @include mobile-device {
      display: none;
    }
  }

  .right-menu {
    //line-height: 28px;
    overflow: hidden;
    flex: 1;
    text-align: right;
    .top-bar-setting-icon {
      cursor: pointer;
      height: 30px;
      width: 30px;
      margin-top: 15px;
      fill: $blue;
    }
  }
</style>
