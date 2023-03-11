<template>
  <el-row class="login">
    <el-col
      class="login-left-panel"
      :span="10"
    >
      <img
        class="qr-code"
        :src="qrCodeUrl"
        alt=""
      >
      <div style="margin-top: 10px;">微信扫码登录</div>
    </el-col>
    <el-col
      class="login-body"
      :span="14"
    >
      <Tab
        :config="tabConfig"
        @change="handleTabChange"
      />
    </el-col>

  </el-row>
</template>

<script>
import { reactive, toRefs, defineComponent, h, markRaw } from 'vue'
import AccountLogin from './components/accountLogin'
import Tab from '@/shared/components/Tab/index.vue'

export default defineComponent({
  components: {
    AccountLogin, Tab
  },
  setup(props, ctx) {
    const state = reactive({
      qrCodeUrl: '/src/assets/images/qr.png',
      tabConfig: {
        defaultTabIdx: 0,
        tabs: [
          {
            label: '账户登录',
            name: 'accountLogin',
            component: markRaw(AccountLogin)
          },
          {
            label: '验证码登录',
            name: 'phoneLogin',
            component: h('div', '开发中')
          }
        ],

      },
      handleTabChange: (tab, event) => {
        console.log('--------handleTabChange-------', tab, event)
      }
    })
    return {
      ...toRefs(state),
    }
  }
})
</script>

<style lang="scss" scoped>
.login {
    background-color:#f4f4f4d1;
    position: absolute;
    left: 50%;
    top: 50%;
    width: 800px;
    height: 500px;
    border-radius: 5px;
    transform: translate(-50%, -50%);
    box-shadow: 0 0 30px 15px rgba(0, 0, 0, .4);
  }
  .login-body{
    padding: 65px 100px;
  }
  .login-left-panel{
    margin:auto;
    text-align: center;
  }

  .qr-code{
    height:120px;
    width:120px;
  }
</style>
