<!--
* @fileName EndpointModal.vue
* @author Di Sheng
* @date 2023/04/11 08:17:39
* @description
!-->
<template>
  <Modal
    ref="modalRef"
    width="550px"
    title="Assign Role Endpoints"
    @close="onModalClose"
    @confirm="onModalConfirm"
  >
    <div>asdf</div>
    <!-- <MyForm
        ref="formRef"
        v-model:formData="formData"
        :config="formConfig"
        :form-items="formItems"
      /> -->
  </Modal>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import { getRoleMenuTree } from '@/api/menu'
import { ElMessage, ElMessageBox } from 'element-plus'
import _ from 'lodash'
export default defineComponent({
  name: 'EndpointModal',
  setup(props, ctx) {
    const state = reactive({
    })
    const onSubmit = async() => {
      console.log('-----onsubmit----')
    }
    const fetchMenuTree = async(roleId) => {
      console.log('-----fetchMenuTree----')
      const { data: res } = await getRoleMenuTree({ id: roleId })
      if (res.errorCode === 0) {
        console.log('-----menu list----', res.data)
      } else {
        ElMessage({
          message: res.msg,
          type: 'error',
          duration: 3 * 1000
        })
      }
    }
    const modalState = reactive({
      modalRef: null,
      onModalOpen(roleId) {
        modalState.modalRef.openModal()
        fetchMenuTree(roleId)
      },
      onModalClose() {
        modalState.modalRef.closeModal()
      },
      onModalConfirm: _.throttle(onSubmit, 2000)
    })

    return {
      ...toRefs(state),
      ...toRefs(modalState)
    }
  }
})
</script>
<style lang='scss' scoped></style>
