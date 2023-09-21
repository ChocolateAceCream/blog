<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/09/21 15:30:06
* @description notification list
!-->
<template>
  <el-checkbox
    v-model="payload.unreadOnly"
    size="large"
    @change="handleCheckboxChange()"
  >unread only</el-checkbox>
  <div
    v-infinite-scroll="onLoad"
    :infinite-scroll-immediate="false"
    class="card-wrapper"
  >
    <el-space
      direction="vertical"
      fill
      wrap
      style="width: 90%"
    >
      <template
        v-for="notification in notificationList"
        :key="notification.id"
      >
        <card
          :notification-info="notification"
          @delete="onDelete"
        />
      </template>
    </el-space>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted, inject } from 'vue'
import { getNotificationList, deleteNotification} from '@/api/notification'
import { ElMessage } from 'element-plus'
import Card from './components/card'
export default defineComponent({
  components: {
    Card,
  },
  setup(props, ctx) {
    const dayjs = inject('dayjs')

    const state = reactive({
      payload: {
        unreadOnly: true,
        cursorId: 0,
        pageSize: 10,
        desc: true, // default article order by create time
      },
      notificationList: [],
    })

    const fetchList = async() => {
      const { data: res } = await getNotificationList({ params: state.payload })
      const { list } = res.data
      if (list.length === 0) {
        ElMessage({
          message: 'No more notifications!',
          type: 'success',
          duration: 3 * 1000
        })
      } else {
        const formattedList = listFormatter(list)
        state.notificationList = [...state.notificationList, ...formattedList]
      }
    }
    onMounted(fetchList)
    const listFormatter = (list) => {
      return list.map(n => {
        n.timestamp = dayjs(n.updatedAt).fromNow()
        if (n.type === 'likeComment') {
          n.content = JSON.parse(n.content)
          n.verb = 'likes your comment'
        }
        return n
      })
    }

    const onLoad = () => {
      state.payload.cursorId = state.notificationList.slice(-1)[0]?.id
      fetchList()
    }

    const handleCheckboxChange = () => {
      state.payload.cursorId = 0
      state.notificationList = []
      fetchList()
    }
    const onDelete = async(id) => {
      const { data: res } = await deleteNotification({ data: { id: id } })
      if (res.errorCode === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete completed',
        })
        state.notificationList = state.notificationList.filter(n => n.id !== id)
      }
    }
    return {
      handleCheckboxChange,
      onLoad,
      onDelete,
      listFormatter,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
