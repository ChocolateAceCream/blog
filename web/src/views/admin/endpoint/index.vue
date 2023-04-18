<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/02/23 10:18:13
* @description
!-->
<template>
  <MyForm
    ref="formRef"
    v-model:formData="formData"
    :config="formConfig"

    :form-items="formItems"
  />
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onAdd"
  > &#43; Add</el-button>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="fetchList"
  >     <SvgIcon
    :icon-name="`icon-blog-search`"
    size="8px"
  /> &nbsp;&nbsp;Search</el-button>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onReset"
  >     <SvgIcon
    :icon-name="`icon-blog-reset`"
    size="8px"
  /> &nbsp;&nbsp;Reset</el-button>
  <my-table
    :data="tableData"
    :config="tableConfig"
    row-key="id"
    @sort-change="sortChange"
  >
    <template #operationBody="scope">
      <el-button
        type="primary"
        link
        @click="onEdit(scope.row)"
      >Edit</el-button>
      <el-button
        type="primary"
        link
        @click="onDelete(scope.row)"
      >Delete</el-button>
    </template>
  </my-table>
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
import { getEndpointList } from '@/api/endpoint'
import { ElMessage } from 'element-plus'
export default defineComponent({
  setup(props, ctx) {
    const methodOptions = [
      {value: 'GET', label: 'GET' },
      {value: 'POST', label: 'POST'},
      {value: 'PUT', label: 'PUT'},
      {value: 'DELETE', label: 'DELETE'}
    ]
    const formState = reactive({
      formData: {},
      formRef: null,
      formConfig: {
        labelPosition: 'left',
        // labelWidth: '10px',
        inline: true
      },
      formItems: [
        { prop: 'group', label: 'Group', type: 'input', options: { placeholder: 'Endpoint group' } },
        { prop: 'path', label: 'Path', type: 'input', options: { placeholder: 'Endpoint path' } },
        { prop: 'name', label: 'Name', type: 'input', options: { placeholder: 'Endpoint name' } },
        { prop: 'method', label: 'Method', type: 'select', options: { placeholder: 'Endpoint name', options: methodOptions, propName: 'method' } },
      ],
    })
    const tableState = reactive({
      tableData: [],
      pagination: {
        pageNumber: 1,
        pageSize: 20,
      },
      sorting: {
        orderBy: 'id',
        desc: false,
      },
      total: 0,
      tableConfig: [
        { label: 'ID', prop: 'id', sortable: 'custom' },
        { label: 'Method', prop: 'method', sortable: 'custom' },
        { label: 'Group', prop: 'group', sortable: 'custom' },
        { label: 'Path', prop: 'path', sortable: 'custom' },
        { label: 'Name', prop: 'name', sortable: 'custom' },
        { label: 'Operation', bodySlot: 'operationBody' },
      ],
      onAdd() {
        console.log('-----formdata----', formState.formData)
      },
      onEdit(row) {

      },
      onDelete(row) {

      },
      onReset() {
        formState.formRef.reset()
      },
    })

    const sortChange = ({ prop, order }) => {
      tableState.sorting.desc = order === 'descending'
      tableState.sorting.orderBy = prop
      fetchList()
    }

    onMounted(() => {
      fetchList()
    })

    const fetchList = async() => {
      const payload = {...formState.formData, ...tableState.pagination, ...tableState.sorting }
      const { data: res } = await getEndpointList({params: payload})
      if (res.errorCode === 0) {
        const {list, total} = res.data
        tableState.tableData = list
        tableState.total = total
      } else {
        ElMessage({
          message: res.msg,
          type: 'error',
          duration: 3 * 1000
        })
      }
    }
    return {
      fetchList,
      sortChange,
      ...toRefs(tableState),
      ...toRefs(formState)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
