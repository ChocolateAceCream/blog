<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/20 11:14:22
* @description drafting new article using md-v3 editor
!-->
<template>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    :loading="onLoading"
    @click="onSave"
  > Save</el-button>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onPublish"
  > Publish</el-button>

  <MyEditor
    ref="editorRef"
    v-model="content"
    class="md-editor"
  />
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { useSessionStore } from '@/stores/sessionStore'
import { postAddArticle, putEditArticle } from '@/api/article'
import { ElMessage, ElMessageBox } from 'element-plus'
export default defineComponent({
  setup(props, ctx) {
    onMounted(() => {
      const store = useSessionStore()
      if (store.currentEditingArticle) {
        state.articleId = store.currentEditingArticle
      } else {
        addArticle()
      }
      console.log('----onMounted----', store.currentEditingArticle)
    })
    const addArticle = async() => {
      const resp = await postAddArticle()
      console.log('----res---', resp)
      const { data: res } = resp
      if (res.errorCode === 0) {
        const store = useSessionStore()
        store.currentEditingArticle = res.data.id
        state.articleId = res.data.id
      }
    }
    const state = reactive({
      content: '',
      articleId: null,
      editorRef: null,
      onLoading: false,
    })
    const onSave = async() => {
      console.log('----onSave----')
      const payload = {
        id: state.articleId,
        content: state.editorRef.text
      }
      const resp = await putEditArticle(payload)
      const { data: res } = resp
      if (res.errorCode === 0) {
        ElMessage({
          message: `Article Saved`,
          type: 'success',
          duration: 3 * 1000
        })
      }
    }
    const onPublish = async() => {
      state.onLoading = !state.onLoading
      console.log('----onPublish----')
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
</style>
