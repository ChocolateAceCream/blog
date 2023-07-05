<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/06 14:19:27
* @description blog preview page
!-->
<template>
  <!-- <MdEditor v-model="text" /> -->
  <MdCatalog
    :editor-id="id"
    :scroll-element="scrollElement"
  />
  <MdPreview
    :editor-id="id"
    :model-value="text"
  />
</template>

<script>
import { useRoute } from 'vue-router'
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { getArticleFile } from '@/api/article'
import { ElMessage } from 'element-plus'
import { MdPreview, MdCatalog } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'
export default defineComponent({
  components: { MdPreview, MdCatalog },
  setup(props, ctx) {
    onMounted(() => {
      onFetchArticle()
    })
    const route = useRoute()
    const state = reactive({
      text: '',
      id: 'preview-only',
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
          ElMessage({
            message: res.msg,
            type: 'success',
            duration: 3 * 1000
          })
          state.text = res.data.content
          console.log('----state.text---', state.text)
        }
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped></style>
