<!--
* @fileName comments.vue
* @author Di Sheng
* @date 2023/07/26 08:48:32
* @description display article comments
!-->
<template>
  <CommentInput @submit="onFetchCommentList" />
  <div class="title">All Comments {{ commentsTotal }}</div>
  <div
    v-infinite-scroll="onFetchCommentList"
    class="comment-list-wrapper"
    :infinite-scroll-immediate="false"
  >
    <template
      v-for="comment in commentList"
      :key="comment.id"
    >
      <div class="commment-wrapper">
        <div class="comment-header">
          {{ comment.author }}
          <div class="timestamp">{{ dayjs(comment.updatedAt).fromNow() }}</div>
        </div>
        <div class="comment-content">
          {{ comment.content }}
        </div>
        <div class="action-box">
          <SvgIcon
            class="item"
            :icon-name="`icon-blog-thumb-up`"
          />
          <span> {{ comment.likesCount }}</span>
          <SvgIcon
            class="item"
            :icon-name="`icon-blog-comments`"
          />
          <span> Reply</span>
        </div>
        <div
          v-if="comment.authorId === currentUserId"
          class="delete"
          @click="onDeleteComment(comment.id)"
        >delete</div>
      </div>
    </template>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { defineComponent, toRefs, reactive, onMounted, inject } from 'vue'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import _ from 'lodash'
import { getCommentList, deleteComment } from '@/api/comment'
import { ElMessage, ElMessageBox } from 'element-plus'
import CommentInput from './commentInput.vue'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  name: 'Comments',
  components: {
    CommentInput
  },
  setup(props, ctx) {
    const dayjs = inject('dayjs')
    const route = useRoute()
    const store = useSessionStore()
    onMounted(() => {
      onFetchCommentList()
    })
    const state = reactive({
      cursorId: 0,
      pageSize: 10,
      desc: false, // default comment order by create time
      articleId: parseInt(route.params.id),
      commentsTotal: 0,
      commentList: [],
      currentUserId: store.userInfo.id
    })

    const onDeleteComment = async(id) => {
      await ElMessageBox.confirm(
        'Deletion will permanently remove the comment, continue? ',
        'Warning',
        {
          confirmButtonText: 'YES',
          cancelButtonText: 'NO',
          type: 'warning',
        }
      )
      const { data: res } = await deleteComment({ data: { id: [id] } })
      ElMessage({
        message: res.msg,
        type: res.errorCode === 0 ? 'success' : 'error',
        duration: 3 * 1000
      })
      if (res.errorCode === 0) state.commentList = state.commentList.filter(item => item.id !== id)
    }

    const onFetchCommentList = async() => {
      const payload = {
        pageSize: state.pageSize,
        cursorId: state.commentList.slice(-1)[0]?.id || 0,
        desc: state.desc,
        articleId: state.articleId
      }
      console.log('---payload-----', payload)
      const { data: res } = await getCommentList({ params: payload })
      if (res.errorCode === 0) {
        const { list, total } = res.data
        state.commentsTotal = total
        console.log('---list-------', list)
        if (list.length === 0) {
          ElMessage({
            message: 'No more comments!',
            type: 'success',
            duration: 3 * 1000
          })
        } else {
          state.commentList = [...state.commentList, ...list]
        }
      } else {
        ElMessage({
          message: res.msg,
          type: 'error',
          duration: 3 * 1000
        })
      }
    }


    return {
      dayjs,
      onDeleteComment,
      onFetchCommentList,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.title{
  position: relative;
  line-height: 30px;
  font-weight: 600;
  font-size: 18px;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 0;
}
.comment-list-wrapper{
  display: block;
  padding-bottom: 40px;
}
.commment-wrapper{
  margin-bottom:40px;
  position: relative;
  .comment-header{
      color: $lite-grey;
      display:flex;
      align-items: center;
      margin-bottom:10px;
    }
  .timestamp{
    margin-left: 100px;
    font-size: 14px;
    line-height: 22px;
  }
  .comment-content{
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
    -webkit-box-orient: vertical;
  }

  .delete{
    display: none;
    position: absolute;
    left: 600px;
    bottom: 0;
    font-size: 14px;
    line-height: 22px;
    text-align: right;
    cursor: pointer;
    color: $warning-red;
  }

  &:hover .delete {
    display: block;
  }
}
.action-box {
  display: flex;
  aligh-items: center;
  flex: 0 0 auto;
  margin-top:10px;
  .item {
    margin-right: 8px;
    line-height:22px;
    font-size:14px;
    cursor: pointer;
    color: $lite-grey;
  }
  span {
    margin-right: 16px;
  }
}
</style>
