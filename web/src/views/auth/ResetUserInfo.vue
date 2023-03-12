/<!--
* @fileName RegisterUser.vue
* @author Di Sheng
* @date 2022/03/14 10:11:23
* @description: register user
!-->
<template>
  <el-col class="reset-user-info-wrapper">
    <el-row class="title">用户信息修改</el-row>
    <el-row
      class="form-wrapper"
      :span="24"
    >
      <el-form
        ref="resetUserInfoFormRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item
          prop="account"
          label="账号"
        >
          <el-input
            v-model="form.account"
            disabled
            placeholder="账号"
          />
        </el-form-item>
        <el-form-item
          prop="username"
          label="姓名"
        >
          <el-input
            v-model="form.username"
            placeholder="请输入姓名"
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
          prop="depart"
          label="部门"
        >
          <el-autocomplete
            v-model="form.depart"
            :fetch-suggestions="handleAsyncDepartmentSearch"
            placeholder="请输入部门"
            @select="handleSelectDepartment"
          />
        </el-form-item>
        <el-form-item
          prop="code"
          label="验证码"
        >
          <el-space>
            <el-input
              v-model="form.code"
              placeholder="验证码"
            />
            <VerificationCode ref="verificationCodeRef" />
          </el-space>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            @click="onSubmit"
          >完成</el-button>
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
import {validateEmail} from '@/utils/validate'
import VerificationCode from '@/shared/components/VerificationCode'
import { postAddUser, getDepartList } from '@/api/auth'
import {putUpdateUserInfo} from '@/api/auth'
import { useSessionStore } from '@/stores/sessionStore'
import useLoading from '@/shared/useLoading'
import { ElMessage } from 'element-plus'
export default defineComponent({
  components: {
    VerificationCode
  },
  setup(props, ctx) {
    const router = useRouter()
    const resetUserInfoFormRef = ref()
    const store = useSessionStore()
    const {account, depart, email, username, id} = store.userInfo
    const form = ref({
      account: account,
      depart: depart,
      email: email,
      username: username,
    })
    const state = reactive({
      handleSelectDepartment: (item) => {
        form.value.depart = item.value
      },
      handleAsyncDepartmentSearch: async(query, cb) => {
        console.log('---query--', query)
        const payload = {params: {depart: query}}
        const {data: res} = await getDepartList({params: payload})
        console.log('--------department list------', res.data)
        const list = res.data.map(item => {
          return {value: item}
        })
        cb(list)
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        code: [
          { required: true, message: '请输入验证码', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入电子邮箱', trigger: 'blur'},
          { validator: validateEmail, trigger: 'blur' }
        ],
        depart: [
          { required: true, message: '请输入部门', trigger: 'blur'}
        ],
      },
    })


    const verificationCodeRef = ref()
    const {loading, wrapLoading} = useLoading()
    const onSubmit = async() => {
      const resetUserInfoForm = unref(resetUserInfoFormRef)
      try {
        await resetUserInfoForm.validate()
        wrapLoading(async() => {
          const {username, email, depart} = form.value
          const payload = {
            id: id,
            username: username,
            email: email,
            depart: depart,
          }
          putUpdateUserInfo(payload).then(response => {
            console.log('------response---', response)
            const {data: res} = response
            if (res.errorCode === 0) {
              console.log('------login success---', res.data)
              Object.keys(payload).map(key => {
                store.userInfo[key] = payload[key]
              })
              ElMessage({
                type: 'success',
                message: '修改成功',
              })
              router.back()
            }
          }).catch((err) => {
            ElMessage({
              type: 'error',
              message: '修改失败' + err,
            })
            verificationCodeRef.value.getUrl()
          })
        })
      } catch (err) {
        console.log('-----form validation err-', err)
      }
      console.log('-----onSubmit----')
    }
    const handleStepBack = () => {
      router.back()
    }

    return {
      resetUserInfoFormRef,
      verificationCodeRef,
      onSubmit,
      handleStepBack,
      form,
      loading,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.reset-user-info-wrapper {
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
