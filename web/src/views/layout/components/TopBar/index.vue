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
          class="setting-icon"
          icon-name="icon-blog-setting"
          color="#3498db"
          size="30px"
        />
        <template #dropdown>
          <el-dropdown-menu>
            <router-link :to="{ name: 'resetPassword' }">
              <el-dropdown-item>
                {{ $t('message.setting.resetPassword') }}
              </el-dropdown-item>
            </router-link>
            <el-dropdown-item
              divided
              @click.stop="handleLogout"
            >
              {{ $t('message.setting.logout') }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <el-dropdown
        class="avatar-container right-menu-item"
        trigger="click"
      >
        <div>
          <SvgIcon
            icon-name="icon-blog-global"
            color="#3498db"
            size="30px"
          />
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item
              v-for="(item,idx) in Object.keys(locales)"
              :key="idx"
              :divided="idx > 0"
              @click.stop="locale=item"
            >
              {{ locales[item] }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <notification />
    </el-col>
  </el-row>
</template>
<script>
import { reactive, toRefs, defineComponent, computed } from 'vue'
import { ElMessageBox } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
import { logout } from '@/shared/hooks/index'
import Breadcrumb from './components/Breadcrumb.vue'
import Notification from './components/Notification.vue'
export default defineComponent({
  components: {
    Breadcrumb, Notification
  },
  setup(props, ctx) {
    const store = useSessionStore()
    console.log('------topbar reload----')
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
    const localeState = reactive({
      locales: {
        'en': 'English',
        'cn': '中文'
      },
      locale: computed({
        get: () => store.getLocale,
        set: val => {
          store.setLocale(val)
        }
      })
    })
    return { ...toRefs(state), ...toRefs(localeState) }
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
  i{
    cursor: pointer;
    margin-left:25px;
  }

  .setting-icon:hover {
      animation: logo-spin infinite 20s linear;
      transition: transform 0.5s;
    }
  }
.left-header{
  display: flex;
  align-items: center;
}

@keyframes logo-spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
