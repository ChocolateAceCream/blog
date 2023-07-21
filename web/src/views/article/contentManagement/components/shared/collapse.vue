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
          <span class="title">{{ item.title }}</span>
          <span class="date">{{ formatTimeToStr(item.updatedAt, 'yyyy-MM-dd hh:mm:ss') }}</span>
        </div>
      </template>
      <div>{{ item.abstract }}</div>
      <div>
        <el-button
          type="primary"
          style="margin-top: 15px;"
          @click="onEdit(item.id)"
        > &#43; Edit</el-button>
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
import router from '@/router'
export default defineComponent({
  name: 'Collapse',
  props: {
    list: {
      type: Array,
      default: () => []
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
    const onEdit = (id) => {
      router.push({ path: '/article/' + id })
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

  .title {
    font-weight: bold;
    font-size: 16px;
    color: $blue;
    margin-right:20px;
    // style="float:left;font-weight:bold;font-size:14px;color:#2C8DF4;"
  }
  .date {
    font-style: italic;
  }
}
</style>
