<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/02/23 10:18:13
* @description
!-->
<template>
  <MyForm
    ref="searchFormRef"
    v-model:formData="searchFormData"
    :config="searchFormConfig"
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
  >
    <SvgIcon
      :icon-name="`icon-blog-search`"
      size="8px"
    /> &nbsp;&nbsp;Search
  </el-button>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onReset"
  >
    <SvgIcon
      :icon-name="`icon-blog-reset`"
      size="8px"
    /> &nbsp;&nbsp;Reset</el-button>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onBatchDelete"
  >
    <SvgIcon
      :icon-name="`icon-blog-delete`"
      size="8px"
    /> &nbsp;&nbsp;Batch Delete
  </el-button>
  <my-table
    :data="tableData"
    :config="tableConfig"
    row-key="id"
    :selectable="true"
    @sort-change="sortChange"
    @selection-change="handleSelectionChange"
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
  <Modal
    ref="modalRef"
    width="550px"
    :title="modalType + ' Endpoint'"
    @close="onModalClose"
    @confirm="onModalConfirm"
  >
    <MyForm
      ref="formRef"
      v-model:formData="formData"
      :config="formConfig"
      :form-items="formItems"
    />
  </Modal>

</template>

<script>
import { defineComponent, toRefs, reactive, onMounted, computed } from 'vue'
import { getEndpointList, putEditEndpoint, postAddEndpoint, deleteEndpoint } from '@/api/endpoint'
import { ElMessage, ElMessageBox } from 'element-plus'
import _ from 'lodash'
import { useI18n } from 'vue-i18n'
export default defineComponent({
  setup(props, ctx) {
    const { t } = useI18n()
    const methodOptions = [
      {value: 'GET', label: 'GET' },
      {value: 'POST', label: 'POST'},
      {value: 'PUT', label: 'PUT'},
      {value: 'DELETE', label: 'DELETE'}
    ]
    const formState = reactive({
      searchFormData: {},
      searchFormRef: null,
      searchFormConfig: {
        labelPosition: 'left',
        // labelWidth: '10px',
        inline: true
      },
      formItems: [
        { prop: 'path', label: 'Path', type: 'input', options: { placeholder: 'e.g. /api/v1/xxx/xxxx' } },
        { prop: 'name', label: 'Name', type: 'input', options: { placeholder: 'Endpoint name' } },
        { prop: 'method', label: 'Method', type: 'select', options: { placeholder: 'Request Method Type', options: methodOptions, propName: 'method' } },
        { prop: 'groupName', label: 'Group', type: 'input', options: { placeholder: 'Endpoint group' } },
      ],

      // add & edit form
      formData: {},
      formRef: null,
      formConfig: {
        rules: {
          path: [{ required: true, message: 'endpoint path required', trigger: 'blur' }],
          method: [{ required: true, message: 'endpoint method required', trigger: 'blur' }],
          name: [{ required: true, message: 'endpoint name required', trigger: 'blur' }],
          groupName: [{ required: true, message: 'endpoint group required', trigger: 'blur' }],
        },
        labelPosition: 'right',
        labelWidth: '100px'
      },

      deleteIds: null
    })

    const handleSelectionChange = (val) => {
      console.log('-------val-------', _.map(val, 'id'))
      formState.deleteIds = val
    }
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
      tableConfig: computed(() => {
        return [
          { label: 'ID', prop: 'id', sortable: 'custom' },
          { label: t('message.apiTable.method'), prop: 'method', sortable: 'custom' },
          { label: t('message.apiTable.group'), prop: 'groupName', sortable: 'custom' },
          { label: t('message.apiTable.path'), prop: 'path', sortable: 'custom' },
          { label: t('message.apiTable.name'), prop: 'name', sortable: 'custom' },
          { label: t('message.apiTable.operation'), bodySlot: 'operationBody' },
        ]
      }),

      onAdd() {
        modalState.modalType = 'Add'
        formState.formData = {}
        modalState.onModalOpen()
      },
      onEdit(row) {
        modalState.modalType = 'Edit'
        formState.formData = _.cloneDeep(row)
        modalState.onModalOpen()
      },

      async onDelete(row) {
        await ElMessageBox.confirm(
          'Deletion will permanently remove all roles\' api, continue? ',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
        )
        const { data: res } = await deleteEndpoint({ data: { id: [row.id] } })
        if (res.errorCode === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
            duration: 3 * 1000
          })
          fetchList()
        }
      },
      async onBatchDelete() {
        await ElMessageBox.confirm(
          'Deletion will permanently remove all roles\' api, continue? ',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
        )
        const { data: res } = await deleteEndpoint({ data: { id: _.map(formState.deleteIds, 'id') } })
        if (res.errorCode === 0) {
          ElMessage({
            message: res.msg,
            type: 'success',
            duration: 3 * 1000
          })
          fetchList()
        }
      },
      onReset() {
        formState.searchFormRef.reset()
      },
    })

    const sortChange = ({ prop, order }) => {
      tableState.sorting.desc = order === 'descending'
      tableState.sorting.orderBy = _.snakeCase(prop)
      fetchList()
    }

    const fetchList = async() => {
      const payload = {...formState.searchFormData, ...tableState.pagination, ...tableState.sorting }
      const { data: res } = await getEndpointList({params: payload})
      if (res.errorCode === 0) {
        const {list, total} = res.data
        tableState.tableData = list
        tableState.total = total
      }
    }

    onMounted(fetchList)

    const onSubmit = async() => {
      try {
        await formState.formRef.validate()
        let resp = {}
        if (modalState.modalType === 'Add') {
          resp = await postAddEndpoint(formState.formData)
        } else {
          resp = await putEditEndpoint(formState.formData)
        }
        const { data: res } = resp
        if (res.errorCode === 0) {
          ElMessage({
            message: `${modalState.modalType} Endpoint Success`,
            type: 'success',
            duration: 3 * 1000
          })
          fetchList()
          modalState.modalRef.closeModal()
        }
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }

    const modalState = reactive({
      modalRef: null,
      modalType: 'Add',
      onModalOpen() {
        modalState.modalRef.openModal()
        formState.formRef?.clearAllValidate()
      },
      onModalClose() {
        modalState.modalRef.closeModal()
      },
      onModalConfirm: _.throttle(onSubmit, 2000)
    })

    return {
      onSubmit,
      fetchList,
      sortChange,
      handleSelectionChange,
      ...toRefs(tableState),
      ...toRefs(formState),
      ...toRefs(modalState),
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
