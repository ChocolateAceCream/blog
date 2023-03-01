<!--
* @fileName index.vue
* @author Di Sheng
* @date 2022/03/07 13:47:30
* @description:
usage:
<modal
  ref="modalRef"
  @close="onModalClose"
  @confirm="onModalConfirm"
>
  <div>this is body</div>
  <template #title>
    <div>this is title</div>
  </template>
  <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogFormVisible = false">Cancel</el-button>
      <el-button
        type="primary"
        @click="dialogFormVisible = false"
      >Confirm</el-button>
    </span>
  </template>
</modal>

setup(props, ctx) {
  const handleButtonClick = () => {
    modalState.onModalOpen()
  }

  const modalState = reactive({
    modalRef: null,
    onModalOpen() {
      modalState.modalRef.openModal()
    },
    onModalClose() {
      modalState.modalRef.closeModal()
      console.log('-------onCancel------')
    },
    onModalConfirm() {
      modalState.modalRef.closeModal()
      console.log('-------onModalConfirm------')
    }
  })
  return {
    ...toRefs(modalState)
  }
}
!-->
<template>
  <el-dialog
    v-model="visiable"
    @close="onClose"
  >
    <template
      v-if="!noTitle"
      #title
    >
      <slot name="title">
        <span class="dialog-title">{{ title }}</span>
      </slot>
    </template>
    <component :is="defaultSlot" />
    <template
      v-if="!noFooter"
      #footer
    >
      <slot name="footer">
        <span class="dialog-footer">
          <el-button @click="onClose">取消</el-button>
          <el-button
            type="primary"
            @click="onConfirm"
          >确定</el-button>
        </span>
      </slot>
    </template>
  </el-dialog>
</template>

<script>
import { defineComponent, reactive, toRefs } from 'vue'

export default defineComponent({
  props: {
    noFooter: {
      type: Boolean,
      default: false
    },
    noTitle: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    }
  },
  emits: ['close', 'confirm'],
  setup(props, ctx) {
    const defaultSlot = () => ctx.slots.default()
    const state = reactive({
      visiable: false,
      onClose() {
        ctx.emit('close')
      },
      onConfirm() {
        ctx.emit('confirm')
      },
      closeModal() {
        state.visiable = false
      },
      openModal() {
        state.visiable = true
      }
    })
    return {
      defaultSlot,
      ...toRefs(state)
    }
  }
})
</script>


<style lang='scss' scoped>
.dialog-title{
  font-size: 16px;
  font-weight: bold;
}
</style>
