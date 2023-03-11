<!--
* @fileName ResetPassword.vue
* @author Di Sheng
* @date 2022/03/11 09:11:26
* @description
!-->
<template>
  <el-col class="reset-password-wrapper">
    <el-row class="title">密码重置</el-row>
    <el-row
      class="form-wrapper"
      :span="24"
    >
      <el-form
        ref="resetPasswordFormRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item
          prop="newPassword"
          label="新密码"
        >
          <el-input
            v-model="form.newPassword"
            type="password"
            placeholder="8~20位, 至少包含数字、字母、特殊字符中的两种"
          />
        </el-form-item>
        <el-form-item
          prop="passwordConfirm"
          label="确认新密码"
        >
          <el-input
            v-model="form.passwordConfirm"
            type="password"
            placeholder="8~20位, 至少包含数字、字母、特殊字符中的两种"
          />
        </el-form-item>
        <el-form-item
          prop="emailCode"
          label="邮箱验证码"
        >
          <el-space>
            <el-input
              v-model="form.emailCode"
              placeholder="验证码"
            />
            <el-button
              type="warning"
              :disabled="isEmailButtonDisabled"
              @click="onSendCode"
            >{{ isEmailButtonDisabled ? "倒计时" +
              emailCountdown + "秒" : "获取验证码" }}</el-button>
          </el-space>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            @click="onResetPassword"
          >重置密码</el-button>
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
import {validatePassword} from '@/utils/validate'
import { putResetPassword, postSendEmailCode } from '@/api/auth'
import useLoading from '@/shared/useLoading'
import { ElMessage } from 'element-plus'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  setup(props, ctx) {
    const router = useRouter()
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== state.form.newPassword) {
        callback(new Error('两次输入的密码不一致'))
      }
      callback()
    }
    const state = reactive({
      rules: {
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { validator: validatePassword, trigger: 'blur' },

        ],
        passwordConfirm: [
          { required: true, message: '请二次输入新密码', trigger: 'blur' },
          { validator: validateConfirmPassword, trigger: 'blur' }
        ]
      },
      form: {},
    })


    const {loading, wrapLoading} = useLoading()
    const resetPasswordFormRef = ref()
    const onResetPassword = () => {
      console.log('-----onResetPassword----')
      try {
        wrapLoading(async() => {
          const form = unref(resetPasswordFormRef)
          await form.validate()
          const { newPassword, emailCode } = state.form
          const payload = {
            newPassword: newPassword,
            code: emailCode
          }
          putResetPassword(payload).then(response => {
            const {data: res} = response
            if (res.errorCode === 0) {
              ElMessage({
                type: 'success',
                message: '修改成功',
              })
              router.back()
            }
          }).catch((err) => {
            ElMessage({
              type: 'error',
              message: err,
            })
          })
        })
      } catch (err) {
        console.log(err)
      }
    }
    const handleStepBack = () => {
      router.back()
    }
    const emailStates = reactive({
      isEmailButtonDisabled: false,
      emailCountdown: 60,
      onSendCode() {
        const fieldToValidate = ['password', 'passwordConfirm']
        Promise.all(
          fieldToValidate.map(item => {
            const p = new Promise((resolve, reject) => {
              resetPasswordFormRef.value.validateField(item, valid => {
                resolve(valid)
              })
            })
            return p
          })
        ).then(results => {
          const filteredResult = results.filter(item => item !== true)
          console.log('---filteredResult=--- ', filteredResult)
          if (filteredResult.length === 0) {
            // send request here to fetch email code
            const store = useSessionStore()
            const payload = {
              email: store.userInfo.email,
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
    const setTimeCallback = () => {
      emailStates.emailCountdown = 60
      emailStates.isEmailButtonDisabled = true
      const timer = setInterval(() => {
        if (emailStates.emailCountdown === 0) {
          clearInterval(timer)
          emailStates.isEmailButtonDisabled = false
        } else {
          emailStates.emailCountdown--
        }
      }, 1000)
    }


    return {
      loading,
      resetPasswordFormRef,
      onResetPassword,
      setTimeCallback,
      handleStepBack,
      ...toRefs(state),
      ...toRefs(emailStates)
    }
  }
})
</script>
<style lang='scss' scoped>
.reset-password-wrapper {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 600px;
  height: 370px;
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
