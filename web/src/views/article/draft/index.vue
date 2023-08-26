<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/20 11:14:22
* @description drafting new article using md-v3 editor
!-->
<template>
  <div class="button-wrapper">
    <el-input
      v-model="articleInfo.title"
      class="title"
      placeholder="Please input title here..."
    />
    <el-button
      type="primary"
      :loading="onSaving"
      :disabled="onSaving"
      @click="onSave"
    > Save</el-button>
    <el-button
      type="primary"
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
import { defineComponent, toRefs, reactive, onMounted, onBeforeUnmount } from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
import PublishModal from './components/publishModal'
import { postAddArticle, putEditArticle, getArticleFile } from '@/api/article'
export default defineComponent({
  components: {
    PublishModal
  },
  setup(props, ctx) {
    onBeforeUnmount(() => {
      const store = useSessionStore()
      window.clearInterval(state.autoSave)
      store.currentEditingArticle = null
    })
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
        await addArticle()
      }
      state.autoSave = window.setInterval(async() => {
        if (!state.onSaving) {
          if (state.savedContent !== state.articleInfo.content) {
            await onSave()
            state.savedContent = state.articleInfo.content
          }
        }
      }, 10000)
    })
    const addArticle = async() => {
      const resp = await postAddArticle()
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
        title: ''
      },
      editorRef: null,
      onSaving: false,
      autoSave: null,
      modalRef: null,
      savedContent: '',
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
  display: flex;
  margin-bottom: 20px;
  align-items: center;
  .title{
    flex:1;
    margin-right: 100px;
    height:60px;
    font-size:25px;
    font-weight:500;
    border: none;
    outline: none;

    :deep(.el-input__wrapper) {
      box-shadow: 0 0 0 0px var(--el-input-border-color, var(--el-border-color)) inset;
      .el-input__inner {
        cursor: default !important;
      }
    }
  }
}
</style>
