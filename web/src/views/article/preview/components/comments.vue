<!--
* @fileName comments.vue
* @author Di Sheng
* @date 2023/07/26 08:48:32
* @description display article comments
!-->
<template>
  <EmojiInput
    ref="commentInputRef"
    @submit="onCommentSubmit"
  />
  <div class="title">All Comments {{ commentsTotal }}</div>
  <div
    v-infinite-scroll="onFetchCommentList"
    class="comment-list-wrapper"
    :infinite-scroll-immediate="false"
  >
    <template
      v-for="(comment,idx) in commentList"
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
        <div
          class="action-box"
        >
          <SvgIcon
            class="item"
            :icon-name="comment.isLiked ? `icon-blog-thumb-up-fill` : `icon-blog-thumb-up-line`"
            @click="onLikeComment(comment)"
          />

          <span> {{ comment.likesCount }}</span>
          <div
            @click="onReplying(comment)"
          >
            <SvgIcon
              class="item"
              :icon-name="`icon-blog-comments`"
            />
            <span>{{ comment.isReplying ? 'Cancel': 'Reply' }} </span>
          </div>

          <div
            v-if="comment.repliesCount>0"
            @click="onShowReplies(comment)"
          >
            <SvgIcon
              class="item"
              :icon-name="`icon-blog-check-all-replies`"
            />
            <span>All Replies ({{ comment.repliesCount }})</span>
          </div>
        </div>
        <div
          v-if="comment.authorId === currentUserId"
          class="delete"
          @click="onDeleteComment(comment.id)"
        >delete</div>
      </div>
      <EmojiInput
        v-if="comment.isReplying"
        ref="replyInputRef"
        @submit="onReplySubmit(comment, $event)"
      />
      <Replies
        v-if="comment.showReplies && comment.replyList?.length > 0"
        :current-user-id="currentUserId"
        :comment="comment"
        @update-comment="updateComment(idx,$event)"
      />
    </template>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { defineComponent, toRefs, reactive, onMounted, inject } from 'vue'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import _ from 'lodash'
import { getCommentList, deleteComment, likeComment, postAddComment } from '@/api/comment'
import { getReplyList, postAddReply } from '@/api/reply'
import { ElMessage, ElMessageBox } from 'element-plus'
import EmojiInput from './emojiInput.vue'
import Replies from './replies.vue'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  name: 'Comments',
  components: {
    EmojiInput, Replies
  },
  setup(props, ctx) {
    const dayjs = inject('dayjs')
    const route = useRoute()
    const store = useSessionStore()
    onMounted(() => {
      onFetchCommentList()
    })
    const onLikeComment = _.throttle(async(comment) => {
      console.log('---onLikeComment-------')
      const payload = {
        commentId: comment.id,
        like: !comment.isLiked,
      }
      const { data: res } = await likeComment(payload)
      if (res.errorCode === 0) {
        console.log('---success-------', res)
        comment.isLiked = !comment.isLiked
        comment.likesCount = comment.isLiked ? comment.likesCount + 1 : comment.likesCount - 1
      }
    }, 1000)
    const state = reactive({
      commentInputRef: null,
      currentReplyingComment: {},
      cursorId: 0,
      pageSize: 10,
      desc: false, // default comment order by create time
      articleId: parseInt(route.params.id),
      commentsTotal: 0,
      commentList: [],
      currentUserId: store.userInfo.id
    })

    const replyState = reactive({
      replyInputRef: null,
      pageSize: 10,
      parentReply: {},
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
      const { data: res } = await deleteComment({ data: { id: id } })
      ElMessage({
        message: res.msg,
        type: res.errorCode === 0 ? 'success' : 'error',
        duration: 3 * 1000
      })
      if (res.errorCode === 0) state.commentList = state.commentList.filter(item => item.id !== id)
    }

    const onFetchCommentList = _.throttle(async() => {
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
      }
    }, 1000)

    const onShowReplies = (comment) => {
      if (!comment.showReplies) {
        // currently replies not shown, so fetch the reply list
        fetchReplyList(comment)
        comment.showReplies = true
      } else {
        // current showing replies, so hide it
        comment.showReplies = false
      }
    }

    const fetchReplyList = async(comment) => {
      const payload = {
        pageSize: replyState.pageSize,
        cursorId: comment.replyList?.slice(-1)[0]?.id || 0, // 0 ?
        desc: state.desc,
        commentId: comment.id
      }
      console.log('---payload-----', payload)
      const { data: res } = await getReplyList({ params: payload })
      if (res.errorCode === 0) {
        const { list, total } = res.data
        comment.repliesCount = total
        console.log('---reply list-------', list)
        if (list.length === 0) {
          ElMessage({
            message: 'No more replies!',
            type: 'success',
            duration: 3 * 1000
          })
        } else {
          comment.replyList = comment.replyList ? [...comment.replyList, ...list] : list
        }
      }
    }

    const onCommentSubmit = async(content) => {
      const payload = {
        articleId: state.articleId,
        commentContent: content
      }
      const { data: res } = await postAddComment(payload)
      if (res.errorCode === 0) {
        onFetchCommentList()
        state.commentInputRef.reset()
        ElMessage({
          message: 'comment posted',
          type: 'success',
          duration: 3 * 1000
        })
      }
    }

    const onReplySubmit = async(comment, content) => {
      console.log('----comment, content-------', comment, content)
      const payload = {
        replyContent: content,
        commentID: state.currentReplyingComment.id,
      }
      if (replyState.parentReply.id) {
        payload.parentReply = replyState.parentReply.id
      }
      const { data: res } = await postAddReply(payload)
      if (res.errorCode === 0) {
        state.currentReplyingComment.isReplying = false // reset emoji input
        ElMessage({
          message: 'reply posted',
          type: 'success',
          duration: 3 * 1000
        })

        // display replies
        comment.repliesCount += 1
        comment.showReplies = true
        fetchReplyList(comment)
      }
    }
    const onReplying = (comment) => {
      comment.isReplying = !comment.isReplying
      if (comment !== state.currentReplyingComment) {
        state.currentReplyingComment.isReplying = false
      }
      state.currentReplyingComment = comment
    }

    const updateComment = (idx, newComment) => {
      state.commentList[idx] = newComment
      // close
    }
    return {
      updateComment,
      onReplySubmit,
      onReplying,
      onCommentSubmit,
      onLikeComment,
      onShowReplies,
      dayjs,
      onDeleteComment,
      onFetchCommentList,
      ...toRefs(state),
      ...toRefs(replyState)
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
  margin-bottom:20px;
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
  align-items: center;
  flex: 0 0 auto;
  margin-top:10px;
  .item {
    margin-right: 8px;
    line-height:22px;
    font-size:14px;
    cursor: pointer;
    color: $blue;
  }
  span {
    margin-right: 16px;
    cursor: pointer;
  }
}
</style>
