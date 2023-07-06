<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/06 14:19:27
* @description blog preview page
!-->
<template>
  <h1>{{ title }}</h1>
  <div style="font-size:12px;">
    <el-link type="primary">
      <span>Author: {{ author }}</span>
    </el-link> &nbsp;&nbsp;
    <el-text
      type="info"
      style="vertical-align: top;"
    >
      <time datetime="publishDate">Last Update: {{ publishDate }}</time>  &nbsp;&nbsp;
    </el-text>

    <el-button
      type="primary"
      link
      @click="onEdit()"
    >Edit</el-button>
  </div>
  <MdPreview
    :editor-id="id"
    :model-value="text"
  />
</template>

<script>
import { useRoute } from 'vue-router'
import router from '@/router'
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { getArticleFile } from '@/api/article'
import { ElMessage } from 'element-plus'
import { MdPreview } from 'md-editor-v3'
import { formatTimeToStr } from '@/utils/date'
import { useSessionStore } from '@/stores/sessionStore'
import 'md-editor-v3/lib/preview.css'
export default defineComponent({
  components: { MdPreview },
  setup(props, ctx) {
    onMounted(() => {
      onFetchArticle()
    })
    const route = useRoute()
    const state = reactive({
      text: '',
      title: '',
      author: '',
      id: 'preview-only',
      publishDate: '',
      scrollElement: document.documentElement,
      onClose: () => {
        onFetchArticle()
      }
    })

    const onFetchArticle = async() => {
      try {
        const resp = await getArticleFile({ params: { id: parseInt(route.params.id) } })
        console.log('----res---', resp)
        const { data: res } = resp
        if (res.errorCode === 0) {
          state.text = res.data.content
          state.title = res.data.title
          state.author = res.data.author.username
          const date = new Date(res.data.updatedAt)
          state.publishDate = formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
        }
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }

    const onEdit = () => {
      console.log('---onEdit---')
      const store = useSessionStore()
      store.currentEditingArticle = parseInt(route.params.id)
      router.push({ path: '/article/draft' })
    }
    return {
      onEdit,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped></style>
