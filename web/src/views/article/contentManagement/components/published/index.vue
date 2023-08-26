<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/07/18 16:54:38
* @description: display published article list
!-->
<template>
  <div class="search-bar-wrapper">
    <my-search-bar
      class="search-bar"
      @change="onSearch"
    />
  </div>
  <Collapse
    ref="myCollapse"
    :list="articleList"
    :published="true"
    @delete="onDelete"
  />
  <Pagination
    v-model:currentPage="pagination.pageNumber"
    v-model:pageSize="pagination.pageSize"
    :total="total"
    style="justify-content: end"
    @change="fetchList"
  />
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getArticleSearchList, deleteArticle } from '@/api/article'
import Collapse from '../shared/collapse.vue'
export default defineComponent({
  name: 'Draft',
  components: { Collapse },
  setup(props, ctx) {
    onMounted(() => {
      fetchList()
    })
    const state = reactive({
      keywords: '',
      articleList: [],
      total: 0,
      myCollapse: null,
      pagination: {
        pageNumber: 1,
        pageSize: 10,
      },
    })
    const onSearch = (input) => {
      console.log('----input---', input)
      state.keywords = input
      fetchList()
    }
    const fetchList = async() => {
      state.myCollapse.resetActiveNames()
      const payload = {
        keywords: state.keywords,
        published: 1, // published
        ...state.pagination
      }
      const { data: res } = await getArticleSearchList({ params: payload })
      if (res.errorCode === 0) {
        const { list, total } = res.data
        state.articleList = list
        state.total = total
        console.log('---total---', total)
        console.log('---articleList---', state.articleList)
        // tableState.total = total
      }
    }
    const onDelete = async(id) => {
      console.log('--------id------', id)
      await ElMessageBox.confirm(
        'Deletion will permanently remove all article contents, continue? ',
        'Warning',
        {
          confirmButtonText: 'YES',
          cancelButtonText: 'NO',
          type: 'warning',
        }
      )
      const { data: res } = await deleteArticle({ data: { id: [id] } })
      ElMessage({
        message: res.msg,
        type: res.errorCode === 0 ? 'success' : 'error',
        duration: 3 * 1000
      })
      fetchList()
    }
    return {
      fetchList,
      onSearch,
      onDelete,
      ...toRefs(state),
    }
  }
})
</script>
<style lang='scss' scoped>
.search-bar-wrapper {
  text-align: right;
  margin: 0 20px 10px 0;

  .search-bar {
    width: 450px;
    display: inline-block;
  }
}
</style>

