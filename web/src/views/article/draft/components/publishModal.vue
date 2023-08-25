<!--
* @fileName publishModal.vue
* @author Di Sheng
* @date 2023/07/05 10:35:02
* @description
!-->
<template>
  <Modal
    ref="modalRef"
    width="550px"
    title="Publish Article"
    @close="closeModal"
    @confirm="onModalConfirm"
  >
    <MyForm
      ref="formRef"
      v-model:formData="formData"
      :config="formConfig"
      :form-items="formItems"
    />
  </Modal>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { putEditArticle } from '@/api/article'
import { useSessionStore } from '@/stores/sessionStore'
import { throttle } from 'lodash-es'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'PublishModal',
  props: {
    articleInfo: {
      type: Object,
      default: () => {
        return {
          articleId: null,
          content: null,
          authorId: null
        }
      }
    }
  },
  emits: ['onSubmit'],
  setup(props, ctx) {
    const router = useRouter()

    const onSubmit = async() => {
      console.log('----props---', props)
      console.log('----onSubmit-----')
      try {
        await formState.formRef.validate()
        const { articleInfo } = props
        const payload = {
          ...articleInfo,
          ...formState.formData,
          published: 1,
        }
        console.log('------payload---', payload)
        const { data: res } = await putEditArticle(payload)
        if (res.errorCode === 0) {
          ElMessage({
            message: `Article Published`,
            type: 'success',
            duration: 3 * 1000
          })
          // TODO: jump to article preview page

          console.log('---articleInfo.id---', articleInfo.id)
          await router.push({ path: `preview/${articleInfo.id}` })
          const store = useSessionStore()
          store.currentEditingArticle = null
        }
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }
    const modalState = reactive({
      modalRef: null,
      onModalConfirm: throttle(onSubmit, 2000)
    })
    const formState = reactive({
      formRef: null,
      formData: {
        title: '',
        abstract: '',
      },
      formItems: [
        { prop: 'title', label: 'Title', type: 'input', options: { placeholder: 'Article Title', maxlength: 80 } },
        { prop: 'abstract', label: 'Abstract', type: 'input', options: { placeholder: '', type: 'textarea', maxlength: 150, showWordLimit: true } },
        // { prop: 'method', label: 'Method', type: 'select', options: { placeholder: 'Request Method Type', options: methodOptions, propName: 'method' } },
      ],
      formConfig: {
        rules: {
          title: [{ required: true, message: 'article title required', trigger: 'blur' }],
          abstract: [{ required: true, message: 'article abstract required', trigger: 'blur' }],
        },
        labelPosition: 'right',
        labelWidth: '100px'
      },
    })

    const openModal = () => {
      formState.formData = {
        ...props.articleInfo
      }
      modalState.modalRef.openModal()
    }
    const closeModal = () => {
      modalState.modalRef.closeModal()
    }
    return {
      openModal,
      closeModal,
      onSubmit,
      ...toRefs(modalState),
      ...toRefs(formState)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
