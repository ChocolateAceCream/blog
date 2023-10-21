<!--
* @fileName accountLogin.vue
* @author Di Sheng
* @date 2022/12/29 18:41:05
* @description
!-->
<template>
  <div ref="loadingTarget">
    <div class="login-title">欢迎登录</div>
    <el-form
      ref="loginFormRef"
      :model="form"
      :rules="rules"
      label-width="0px"
      @keyup.enter="onSubmit"
    >
      <el-form-item prop="username">
        <el-input
          v-model="form.username"
          placeholder="用户名"
        />
      </el-form-item>
      <el-form-item prop="password">
        <el-space>
          <el-input
            v-model="form.password"
            type="password"
            placeholder="密码"
          />
          <el-button
            link
            @click="handleForgetPassword"
          >忘记密码</el-button>
        </el-space>

      </el-form-item>
      <el-form-item prop="code">
        <el-space>
          <el-input
            v-model="form.code"
            placeholder="验证码"
          />
          <VerificationCode ref="verificationCodeRef" />
        </el-space>
      </el-form-item>
      <div class="login-btn">
        <el-button
          type="primary"
          @click="onSubmit"
        >登录</el-button>
      </div>
      <div class="login-btn-wrapper">
        <el-button
          link
          @click="handleRegisterUser"
        > 新用户注册</el-button>
      </div>

    </el-form>
  </div>
</template>

<script>
import { reactive, toRefs, defineComponent, unref, ref } from 'vue'
import VerificationCode from '@/shared/components/VerificationCode'
import { useRouter } from 'vue-router'
import { validatePassword } from '@/utils/validate'
import { useSessionStore } from '@/stores/sessionStore'
import { useRouterStore } from '@/stores/routerStore'
import { throttle } from 'lodash-es'
import useLoading from '@/shared/useLoading'
import { postLogin } from '@/api/auth'

export default defineComponent({
  components: {
    VerificationCode
  },
  setup(props, ctx) {
    const router = useRouter()
    const loginFormRef = ref()
    const verificationCodeRef = ref()
    const login = async() => {
      try {
        const form = unref(loginFormRef)
        await form.validate()
        wrapLoading(async() => {
          const { username, password, code } = state.form
          const payload = {
            username: username,
            password: password,
            // password: md5(password),
            code: code,
          }
          postLogin(payload).then(async response => {
            const { data: res } = response
            if (res.errorCode === 0) {
              console.log('------login success---', res.data)
              const sStore = useSessionStore()
              sStore.setUserInfo(res.data.user)
              sStore.setNotificationWebsocket()
              const routerStore = useRouterStore()
              await routerStore.setAsyncRouter()
              router.push({ name: 'home' })
            }
          }).catch(() => {
            verificationCodeRef.value.getUrl()
          })
        })
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }
    const state = reactive({
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        code: [
          { required: true, message: '请输入验证码', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入验密码', trigger: 'blur' },
          { validator: validatePassword, trigger: 'blur' }
        ]
      },
      onSubmit: throttle(login, 1000),
      form: {},
    })
    const { loadingTarget, wrapLoading } = useLoading()


    const handleForgetPassword = () => {
      router.push({ name: 'resetPassword'})
    }

    const handleRegisterUser = () => {
      router.push({ name: 'register' })
    }
    return {
      verificationCodeRef,
      handleRegisterUser,
      handleForgetPassword,
      loginFormRef,
      loadingTarget,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.login-title {
  width: 100%;
  text-align: center;
  font-size: 20px;
  color: $dark-brown;
  ;
  margin-bottom: 20px;
}

button {
  @include when-inside('.login-btn') {
    background-image: linear-gradient(270deg, #87B5FE 0%, #4F7EFC 100%);
    width: 100%;
    border-radius: 20px;
    height: 40px;
  }
}

.login-btn-wrapper {
  text-align: end;
  width: 100%;
}
</style>
