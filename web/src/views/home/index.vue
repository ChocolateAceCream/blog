<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/05/20 21:32:11
* @description
!-->
<template>
  <div class="toolbar">
    <div class="icon-wrapper">
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-add`"
        @click="onAddArticle"
      />
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-favorite`"
      />
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-share`"
      />
    </div>
  </div>
  <div
    v-infinite-scroll="onLoad"
    :infinite-scroll-immediate="false"
    class="card-wrapper"
  >
    <template
      v-for="article in articleList"
      :key="article.id"
    >
      <card :article-info="article" />
    </template>

  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getArticleList } from '@/api/article'
import Card from './components/card'
import router from '@/router'
export default defineComponent({
  components: {
    Card,
  },
  setup(props, ctx) {
    const onFetchArticleList = async(payload) => {
      const { data: res } = await getArticleList({ params: state.payload })
      if (res.errorCode === 0) {
        const { list } = res.data
        if (list.length === 0) {
          ElMessage({
            message: 'No more articles!',
            type: 'success',
            duration: 3 * 1000
          })
        } else {
          state.articleList = [...state.articleList, ...list]
        }
      }
    }
    onMounted(onFetchArticleList)
    const state = reactive({
      articleList: [],
      payload: {
        cursorId: 0,
        pageSize: 10,
        desc: true, // list recent posted article first
      }
    })
    const onLoad = () => {
      state.payload.cursorId = state.articleList.slice(-1)[0]?.id
      onFetchArticleList()
    }
    const onAddArticle = () => {
      router.push({
        name: 'draft',
      })
    }
    return {
      onAddArticle,
      onLoad,
      onFetchArticleList,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.toolbar {
  display: flex;
  flex-direction: row-reverse;
  height: 30px;

  .toolbar-icon {
    margin-left: 20px;
    cursor: pointer;
  }
}

.card-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
}

</style>
