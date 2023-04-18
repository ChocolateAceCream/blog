<template>
  <el-pagination
    :current-page="currentPage"
    :page-size="pageSize"
    :page-sizes="[10, 15, 20, 25]"
    :small="false"
    :disabled="false"
    :background="false"
    layout="total, prev, pager, next, sizes, jumper"
    :total="total"
    style="padding-right:20px"
    @current-change="handleCurrentChange"
    @size-change="handleSizeChange"
  />
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    currentPage: {
      type: Number,
      default: 1,
    },
    pageSize: {
      type: Number,
      default: 10,
    },
    total: {
      type: Number,
      default: 0,
    },
  },
  emits: ['update:currentPage', 'update:pageSize', 'change'],
  setup(prop, ctx) {
    const handleCurrentChange = (val) => {
      // console.log(`current page: ${val} `)
      ctx.emit('update:currentPage', val)
      ctx.emit('change')
    }

    const handleSizeChange = (size) => {
      // console.log(`${size} items per page`)
      ctx.emit('update:pageSize', size)
      ctx.emit('change')
    }

    return {
      handleCurrentChange,
      handleSizeChange,
    }
  },
})
</script>
