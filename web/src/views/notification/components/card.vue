<!--
* @fileName card.vue
* @author Di Sheng
* @date 2023/07/10 21:31:37
* @description
!-->
<template>
  <el-card
    class="abstract-card"
  >
    <span
      v-if="notificationInfo.status === 'unread'"
      class="unread-dot"
    />
    <div class="content">
      <el-button
        type="primary"
        style="margin-bottom:4px"
        link
      > {{ notificationInfo.initiator }}</el-button>
      <span>{{ notificationInfo.verb }}</span>
    </div>
    <div class="timestamp">{{ notificationInfo.timestamp }} </div>
    <div class="card-footer">
      <el-button
        type="primary"
        link
        @click="onViewNotification(notificationInfo)"
      > View</el-button>
      <el-button
        type="danger"
        link
        @click.stop="onDeleteNotification(notificationInfo.id)"
      > Delete</el-button>

    </div>
  </el-card>
</template>

<script>
import { defineComponent, toRefs, reactive, inject, computed } from 'vue'
import { readNotification } from '@/api/notification'
import { useSessionStore } from '@/stores/sessionStore'
import router from '@/router'
export default defineComponent({
  name: 'Card',
  props: {
    notificationInfo: {
      type: Object,
      default: () => {
        return {
          notificationId: null,
          content: null,
          authorId: null,
          viewedTimes: 0,
        }
      }
    }
  },
  emits: ['delete'],
  setup(props, ctx) {
    const dayjs = inject('dayjs')
    const state = reactive({
      notificationInfo: { ...props.notificationInfo },
      timestamp: computed(() => {
        return dayjs(state.notificationInfo.updatedAt).fromNow()
      })
    })
    const onViewNotification = async(notification) => {
      const { data: res } = await readNotification({ id: notification.id })
      if (res.errorCode === 0) {
        notification.status = 'read'
        useSessionStore().updateNotificationCount()
        if (notification.type === 'likeComment') {
          router.push({ path: '/preview/' + notification.content.articleId })
        }
      }
    }
    const onDeleteNotification = (id) => {
      ctx.emit('delete', id)
    }
    return {
      onViewNotification,
      onDeleteNotification,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.card-header {
  display: flex;
  // justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
  cursor: pointer;

  .dividing {
    width: 1px;
    height: 14px;
    background: $liter-background;
    margin: 0 8px;
  }

  .tag {
    position: relative;
    flex-shrink: 0;
    font-size: 13px;
    line-height: 22px;
    padding: 0 8px;

    &:not(:last-child):after {
      position: absolute;
      right: -1px;
      display: block;
      content: " ";
      width: 2px;
      height: 2px;
      border-radius: 50%;
      background-color: $dark-brown;
    }
  }
}

.content {
  // font-weight: 600;
  // font-size: 18px;
  // line-height: 24px;
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-bottom: 8px;
}

.timestamp {
  font-size: 14px;
  line-height: 22px;
  overflow: hidden;
  margin-bottom: 8px;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.card-footer {
  display: flex;
  align-items: center;
  padding: 0;
  margin: 0;
  font-size: 12px;

  li {
    list-style: none;
    margin-right: 20px;
  }

  .toolbar-icon {
    margin-right: 5px;
  }
}

.abstract-card {
  margin: 2px;
  padding-left: 20px;
  position:relative;

  @include mobile-device {
    width: 100%
  }

  @include desktop-device {
    width: calc(50% - 4px);
  }
}
.unread-dot{
  height: 10px;
  width: 10px;
  background-color: $dark-blue;
  border-radius: 50%;
  left: 17px;
  top: calc(50% - 5px);
  position:absolute;
}
</style>
