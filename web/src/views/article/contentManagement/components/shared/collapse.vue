<!--
* @fileName collapse.vue
* @author Di Sheng
* @date 2023/07/21 16:28:38
* @description
!-->
<template>
  <el-collapse v-model="activeNames">
    <el-collapse-item
      v-for="(item, index) in list"
      :key="index"
      class="deploy-setting"
      :name="index"
    >
      <template #title>
        <div class="title-wrapper">
          <span class="title">{{ item.title || 'Untitled' }}</span>
          <span class="date">{{ formatTimeToStr(item.updatedAt, 'yyyy-MM-dd hh:mm:ss') }}</span>
          <template v-if="published">
            <span class="split-line" />
            <span class="info">Viewed · {{ item.viewedTimes }}</span>
          </template>
        </div>
      </template>
      <div>{{ item.abstract }}</div>
      <div>
        <el-button
          type="primary"
          style="margin-top: 15px;"
          @click="onEdit(item.id)"
        > Edit</el-button>
        <el-button
          type="danger"
          style="margin-top: 15px;"
          @click="onDelete(item.id)"
        >
          <SvgIcon
            :icon-name="`icon-blog-delete`"
            size="8px"
          /> &nbsp;&nbsp;Delete
        </el-button>
      </div>
    </el-collapse-item>
  </el-collapse>
</template>

<script>
import { defineComponent, toRefs, reactive} from 'vue'
import { formatTimeToStr } from '@/utils/date'
import { useRouter } from 'vue-router'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  name: 'Collapse',
  props: {
    list: {
      type: Array,
      default: () => []
    },
    published: {
      type: Boolean,
    }
  },
  emits: ['delete'],
  setup(props, ctx) {
    const state = reactive({
      activeNames: [],
    })
    const resetActiveNames = () => {
      state.activeNames = []
    }
    const router = useRouter()
    const onEdit = (id) => {
      console.log('-------router-------', router)
      const store = useSessionStore()
      store.currentEditingArticle = id
      router.push({ name: 'draft' })
    }

    const onDelete = (id) => {
      ctx.emit('delete', id)
    }
    return {
      resetActiveNames,
      formatTimeToStr,
      onDelete,
      onEdit,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.title-wrapper{
  position:absolute;
  span{
    margin-right:20px;
  }

  .title {
    font-weight: bold;
    font-size: 16px;
    color: $blue;
    // style="float:left;font-weight:bold;font-size:14px;color:#2C8DF4;"
  }
  .date {
    font-style: italic;
  }
  .split-line{
    height: 14px;
    width: 0;
    border-left: 1px solid #e5e6eb;
  }
  .info{
    font-size: 14px;
    color: $lite-grey;
  }
}
</style>
