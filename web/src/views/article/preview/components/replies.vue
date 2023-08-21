<!--
* @fileName reply.vue
* @author Di Sheng
* @date 2023/08/16 08:55:32
* @description dispaly each replies
!-->
<template>
  <div
    class="reply-wrapper"
  >
    <template
      v-for="reply in comment.replyList"
      :key="reply.id"
    >
      <div class="reply">
        <div class="reply-header">
          <span> {{ reply.author }} </span>
          <template v-if="reply.parentReply">
            <span> Reply to  </span>
            <span> {{ reply.parentReply.author.username }}</span>
          </template>
          <div class="timestamp">{{ dayjs(reply.updatedAt).fromNow() }}</div>
        </div>
        <div class="reply-content">
          {{ reply.content }}
        </div>
        <div
          v-if="reply.parentReplyId"
          class="parent-reply-content"
        >
          {{ reply.parentReply ? reply.parentReply.replyContent : '------The Reply Has Been Deleted-------' }}
        </div>
        <div
          class="action-box"
        >
          <SvgIcon
            class="item"
            :icon-name="reply.isLiked ? `icon-blog-thumb-up-fill` : `icon-blog-thumb-up-line`"
            @click="onLikeReply(reply)"
          />

          <span> {{ reply.likesCount }}</span>
          <div
            @click="onReplying(reply)"
          >
            <SvgIcon
              class="item"
              :icon-name="`icon-blog-comments`"
            />
            <span>{{ reply.isReplying ? 'Cancel' : 'Reply' }} </span>
          </div>

          <div
            v-if="reply.authorId === currentUserId"
            class="delete"
            @click="onDeleteReply(reply)"
          >delete</div>
        </div>
        <EmojiInput
          v-if="reply.isReplying"
          ref="replyInputRef"
          @submit="onReplySubmit(reply, $event)"
        />
      </div>
    </template>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, inject } from 'vue'
import { getReplyList, deleteReply, likeReply, postAddReply } from '@/api/reply'
import { ElMessage, ElMessageBox } from 'element-plus'
import {cloneDeep} from 'lodash-es'
import EmojiInput from './emojiInput.vue'
export default defineComponent({
  name: 'Replies',
  components: {
    EmojiInput
  },
  props: {
    comment: {
      type: Object,
      default: () => {
        return {}
      }
    },
    currentUserId: {
      type: Number,
      default: null,
    }
  },
  emits: ['updateComment'],
  setup(props, ctx) {
    console.log('----props.comment---', props.comment)
    const dayjs = inject('dayjs')

    const onDeleteReply = async(reply) => {
      await ElMessageBox.confirm(
        'Deletion will permanently remove the reply, continue? ',
        'Warning',
        {
          confirmButtonText: 'YES',
          cancelButtonText: 'NO',
          type: 'warning',
        }
      )
      console.log('---reply.id---', reply.id)
      const { data: res } = await deleteReply({ data: { id: reply.id } })
      if (res.errorCode === 0) {
        // delete success, comment reply count -1
        const deep = cloneDeep(props.comment)
        deep.repliesCount -= 1
        deep.replyList = deep.replyList.filter(r => r.id !== reply.id)
        deep.replyList = deep.replyList.reduce((accumulator, r) => {
          if (r.id !== reply.id) {
            if (r.parentReplyId === reply.id) {
              r.parentReply = null
            }
            accumulator.push(r) // Multiply each remaining number by 2 and add it to the accumulator
          }
          return accumulator
        }, [])
        ctx.emit('updateComment', deep)
      }
    }

    const state = reactive({
      currentReplyingReply: {},
      pageSize: 10,
    })

    const onReplying = (reply) => {
      reply.isReplying = !reply.isReplying
      if (reply !== state.currentReplyingReply) {
        state.currentReplyingReply.isReplying = false
      }
      state.currentReplyingReply = reply
    }

    const onReplySubmit = async(reply, content) => {
      const payload = {
        replyContent: content,
        commentID: props.comment.id,
        parentReplyId: reply.id
      }
      const { data: res } = await postAddReply(payload)
      if (res.errorCode === 0) {
        state.currentReplyingReply.isReplying = false // reset emoji input
        ElMessage({
          message: 'reply posted',
          type: 'success',
          duration: 3 * 1000
        })

        // update comment
        const deep = cloneDeep(props.comment)
        deep.repliesCount += 1
        await fetchReplyList(deep)
        ctx.emit('updateComment', deep)
        reply.showReplies = true
      }
    }

    const fetchReplyList = async(comment) => {
      const payload = {
        pageSize: state.pageSize,
        cursorId: comment.replyList?.slice(-1)[0]?.id || 0, // 0 ?
        desc: state.desc,
        commentId: comment.id
      }
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


    return {
      onReplying,
      dayjs,
      onReplySubmit,
      onDeleteReply,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.reply-wrapper{
  background: $liter-background;
  margin-bottom:20px;
  max-width:700px;
  position: relative;
  padding: 16px 8px 16px 20px;
}
.reply{
  margin-bottom:20px;
  .reply-header{
    position: relative;
    color: $lite-grey;
    display:flex;
    align-items: center;
    margin-bottom:10px;
    span{
      margin-right:8px;
    }
    .timestamp{
      position: absolute;
      right: 15px;
      font-size: 14px;
      text-align: right;
    }
  }
  .reply-content{
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
    -webkit-box-orient: vertical;
  }
  .parent-reply-content{
    background: $lite-background;
    border: 1px solid $lite-background;
    box-sizing: border-box;
    border-radius: 4px;
    padding: 0 12px;
    line-height: 36px;
    height: 36px;
    font-size: 14px;
    color: $lite-grey;
    margin-top: 8px;
  }
  .action-box {
    display: flex;
    align-items: center;
    flex: 0 0 auto;
    position:relative;
    margin:10px 0 20px 0;
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
  .delete{
    display: none;
    position: absolute;
    right: 15px;
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
</style>
