<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/20 11:14:22
* @description drafting new article using md-v3 editor
!-->
<template>
  <div class="button-wrapper">
    <el-button
      type="primary"
      style="margin-bottom: 15px;"
      :loading="onSaving"
      :disabled="onSaving"
      @click="onSave"
    > Save</el-button>
    <el-button
      type="primary"
      style="margin-bottom: 15px;"
      @click="onPublish"
    > Publish</el-button>
  </div>
  <MyEditor
    ref="editorRef"
    v-model="articleInfo.content"
    class="md-editor"
  />
  <PublishModal
    ref="modalRef"
    :article-info="articleInfo"
  />
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
import PublishModal from './components/publishModal'
import { postAddArticle, putEditArticle, getArticleFile } from '@/api/article'
import { ElMessage } from 'element-plus'
export default defineComponent({
  components: {
    PublishModal
  },
  setup(props, ctx) {
    onMounted(async() => {
      const store = useSessionStore()
      if (store.currentEditingArticle) {
        const resp = await getArticleFile({ params: { id: store.currentEditingArticle } })
        console.log('----res---', resp)
        const { data: res } = resp
        if (res.errorCode === 0) {
          state.articleInfo = res.data
        }
        state.articleId = store.currentEditingArticle
      } else {
        state.articleInfo.authorId = store.userInfo.id
        addArticle()
      }
      let content = ''
      state.autoSave = window.setInterval(() => {
        if (!state.onSaving) {
          if (content !== state.content) {
            onSave()
            content = state.content
          }
        }
      }, 20000)
      console.log('----onMounted----', store.currentEditingArticle)
    })
    const addArticle = async() => {
      const resp = await postAddArticle()
      console.log('----res---', resp)
      const { data: res } = resp
      if (res.errorCode === 0) {
        const store = useSessionStore()
        store.currentEditingArticle = res.data.id
        state.articleInfo = res.data
      }
    }
    const state = reactive({
      articleInfo: {
        content: '',
        articleId: null,
        authorId: null,
      },
      editorRef: null,
      onSaving: false,
      autoSave: null,
      modalRef: null,
    })
    const onSave = async() => {
      console.log('----onSave----')
      state.onSaving = true
      const payload = state.articleInfo
      await putEditArticle(payload)
      state.onSaving = false
    }
    const onPublish = async() => {
      state.modalRef.openModal()
    }
    return {
      onSave,
      onPublish,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.button-wrapper{
  text-align: right;
}
</style>
