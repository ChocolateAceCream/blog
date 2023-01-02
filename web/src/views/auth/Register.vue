/<!--
* @fileName Register.vue
* @author Di Sheng
* @date 2022/03/14 10:11:23
* @description: register user
!-->
<template>
  <el-col class="register-user-wrapper">
    <el-row class="title">新用户注册</el-row>
    <el-row
      class="form-wrapper"
      :span="24"
    >
      <el-form
        ref="registerUserFormRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item
          prop="username"
          label="用户名"
        >
          <el-input
            v-model="form.username"
            placeholder="用户名"
          />
        </el-form-item>
        <el-form-item
          prop="password"
          label="密码"
        >
          <el-input
            v-model="form.password"
            type="password"
            placeholder="8~20位, 至少包含数字、字母、特殊字符中的两种"
          />
        </el-form-item>
        <el-form-item
          prop="passwordConfirm"
          label="确认密码"
        >
          <el-input
            v-model="form.passwordConfirm"
            type="password"
            placeholder="8~20位, 至少包含数字、字母、特殊字符中的两种"
          />
        </el-form-item>
        <el-form-item
          label="邮箱"
          prop="email"
        >
          <el-input
            v-model="form.email"
            placeholder="请输入邮箱"
          />
        </el-form-item>
        <el-form-item
          prop="code"
          label="邮箱验证码"
        >
          <el-space>
            <el-input
              v-model="form.code"
              placeholder="验证码"
            />
            <el-button
              type="warning"
              :disabled="isEmailCodeButtonDisabled"
              @click="onSendCode"
            >{{ isEmailCodeButtonDisabled ? "倒计时" + emailCodeCountdown + "秒" : "获取邮箱验证码" }}</el-button>
          </el-space>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            @click="onSubmit"
          >注册</el-button>
          <el-button
            @click="handleStepBack"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </el-row>
  </el-col>

</template>

<script>
import { reactive, toRefs, defineComponent, unref, ref } from 'vue'
import { useRouter } from 'vue-router'
import {validatePassword, validateEmail} from '@/utils/validate'
import { postRegister, postSendEmailCode } from '@/api/auth'
import useLoading from '@/components/shared/useLoading'
import { ElMessage } from 'element-plus'
import { sessionStore } from '@/stores/sessionStore'
export default defineComponent({
  setup(props, ctx) {
    const router = useRouter()
    const registerUserFormRef = ref()
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== state.form.password) {
        callback(new Error('两次输入的密码不一致'))
      }
      callback()
    }
    const state = reactive({
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur', max: 30 }
        ],
        code: [
          { required: true, message: '请输入邮箱验证码', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { validator: validatePassword, trigger: 'blur' }
        ],
        passwordConfirm: [
          { required: true, message: '请二次输入密码', trigger: 'blur' },
          { validator: validateConfirmPassword, trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入电子邮箱', trigger: 'blur'},
          { validator: validateEmail, trigger: 'blur' }
        ],
      },
      form: {},
    })

    const {loading, wrapLoading} = useLoading()
    const onSubmit = async() => {
      console.log('-----onSubmit----', state.form)
      const form = unref(registerUserFormRef)
      try {
        await form.validate()
        wrapLoading(async() => {
          const payload = {
            ...state.form,
          }
          console.log('-----payload----', payload)
          postRegister(payload).then(response => {
            const {data: res} = response
            if (res.errorCode === 0) {
              console.log('------login success---', res.data)
              const store = sessionStore()
              Object.keys(res.data).map(key => {
                store.userInfo[key] = res.data[key]
              })
              store.userInfo.isAuthenticated = true
              router.push({name: 'home'})
            }
          }).catch((err) => {
            console.log('-----postRegister err------', err)
            ElMessage({
              type: 'error',
              message: '注册失败，请重试',
            })
          })
        })
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }

    const handleStepBack = () => {
      router.back()
    }

    const setTimeCallback = () => {
      emailCodeStates.emailCodeCountdown = 60
      emailCodeStates.isEmailCodeButtonDisabled = true
      const timer = setInterval(() => {
        if (emailCodeStates.emailCodeCountdown === 0) {
          clearInterval(timer)
          emailCodeStates.isEmailCodeButtonDisabled = false
        } else {
          emailCodeStates.emailCodeCountdown--
        }
      }, 1000)
    }
    const emailCodeStates = reactive({
      isEmailCodeButtonDisabled: false,
      emailCodeCountdown: 60,
      onSendCode() {
        const fieldToValidate = ['username', 'password', 'email', 'passwordConfirm', 'account']
        Promise.all(
          fieldToValidate.map(item => {
            const p = new Promise((resolve, reject) => {
              registerUserFormRef.value.validateField(item, valid => {
                resolve(valid)
              })
            })
            return p
          })
        ).then(results => {
          const filteredResult = results.filter(item => item !== true)
          if (filteredResult.length === 0) {
            // send request here to fetch sms code
            const payload = {
              email: state.form.email
            }
            postSendEmailCode(payload).then(response => {
              const { data: res } = response
              if (res.errorCode === 0) {
                ElMessage({
                  type: 'success',
                  message: '验证码已发送至邮箱，请查收',
                })
              }
            }).catch((err) => {
              console.log('-----sendEmailCode err------', err)
              ElMessage({
                type: 'error',
                message: '验证码发送失败，请重试',
              })
            })
            setTimeCallback()
          }
        })
      }
    })
    return {
      registerUserFormRef,
      onSubmit,
      loading,
      handleStepBack,
      ...toRefs(emailCodeStates),
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.register-user-wrapper {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 600px;
  height: 450px;
  border-radius: 5px;
  background: #fff;
  transform: translate(-50%, -50%);
  box-shadow: 0 0 30px 15px rgba(0, 0, 0, .4);
  padding: 15px;
  .title{
    font-size: 14px;
    font-weight: bold;
  }
  .form-wrapper{
    padding: 65px 100px;
  }
}
</style>
