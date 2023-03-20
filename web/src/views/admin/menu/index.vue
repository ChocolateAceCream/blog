<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/02/23 10:18:13
* @description
!-->
<template>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onAddMenu(0)"
  > &#43; Add Root Menu</el-button>
  <my-table
    :data="tableData"
    :config="tableConfig"
    row-key="id"
  >
    <template #iconBody="scope">
      <div v-if="scope.row.icon">
        <SvgIcon
          :icon-name="`icon-blog-` + scope.row.icon"
          size="15px"
        /><span>&nbsp;{{ scope.row.icon }}</span>
      </div>
    </template>
    <template #operationBody="scope">
      <el-button
        type="primary"
        link
        @click="onAddMenu(scope.row.id)"
      >Add</el-button>
      <el-button
        type="primary"
        link
        @click="onEditMenu(scope.row)"
      >Edit</el-button>
      <el-button
        type="primary"
        link
        @click="onDeleteMenu(scope.row.id)"
      >Delete</el-button>
    </template>
  </my-table>
  <Modal
    ref="modalRef"
    width="550px"
    :title="modalTitle"
    @close="onModalClose"
    @confirm="onModalConfirm"
  >
    <MyForm
      ref="formRef"
      v-model:formData="formData"
      :config="formConfig"
      :form-items="formItems"
    >
      <template #icon="scope">
        <Icon v-model:icon="scope.icon" />
      </template>
    </MyForm>
  </Modal>
</template>

<script>
import { defineComponent, toRefs, reactive, unref, ref } from 'vue'
import { useRouterStore } from '@/stores/routerStore'
import { postAddMenu } from '@/api/menu'
import { ElMessage } from 'element-plus'
import _ from 'lodash'
import Icon from './icon.vue'
export default defineComponent({
  components: {
    Icon
  },
  setup(props, ctx) {
    const routerStore = useRouterStore()
    routerStore.$subscribe((mutation, state) => {
      tableState.tableData = state.menuTree
    })
    const tableState = reactive({
      tableData: routerStore.menuTree,
      tableConfig: [
        { label: 'ID', prop: 'id' },
        { label: 'Route', prop: 'routeName' },
        { label: 'path', prop: 'path' },
        { label: 'title', prop: 'title' },
        { label: 'icon', prop: 'icon', bodySlot: 'iconBody' },
        { label: 'operation', bodySlot: 'operationBody' },
      ],
      onAddMenu(id) {
        modalState.modalTitle = 'Add Menu'
        formState.formData = { pid: id }
        modalState.onModalOpen()
      },
      onEditMenu(row) {
        modalState.modalTitle = 'Edit Menu'
        formState.formData = _.cloneDeep(row)
        modalState.onModalOpen()
      },
      onDeleteMenu(id) {
        console.log('----onDeleteMenu--', id)
      }
    })
    const onSubmit = async() => {
      try {
        await formState.formRef.validate()
        const payload = { meta: {} }
        _.forIn(formState.formData, (val, key) => {
          if (['icon', 'title'].includes(key)) {
            payload.meta[key] = val
          } else {
            payload[key] = val
          }
        })
        await postAddMenu(payload).then(async response => {
          const { data: res } = response
          if (res.errorCode === 0) {
            ElMessage({
              message: 'Add Menu Success',
              type: 'success',
              duration: 3 * 1000
            })
            const routerStore = useRouterStore()
            routerStore.asyncRouterFlag = 0
            await routerStore.setAsyncRouter()
            tableState.tableData = routerStore.menuTree
            modalState.modalRef.closeModal()
          }
        })
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }
    const modalState = reactive({
      modalRef: null,
      modalTitle: '',
      onModalOpen() {
        modalState.modalRef.openModal()
        formState.formRef?.clearAllValidate()
      },
      onModalClose() {
        modalState.modalRef.closeModal()
        console.log('-------onCancel------')
      },
      onModalConfirm: _.throttle(onSubmit, 2000)
    })

    const validatePath = (rule, value, callback) => {
      if (formState.formData.pid !== 0 && value[0] === '/') {
        callback(new Error('Relative path should not start with /'))
      } else if (formState.formData.pid === 0 && value[0] !== '/') {
        callback(new Error('Absolute path should start with /'))
      }
      callback()
    }
    const formState = reactive({
      formData: {},
      formRef: null,
      formConfig: {
        rules: {
          route: [{required: true, message: 'route name required', trigger: 'blur' }],
          title: [{ required: true, message: 'title required', trigger: 'blur' }],
          icon: [{ required: true, message: 'icon required', trigger: 'blur' }],
          component: [{ required: true, message: 'component required', trigger: 'blur' }],
          path: [{ required: true, message: 'path required', trigger: 'blur'}, { validator: validatePath, trigger: 'blur' }]
        },
        labelPosition: 'right',
        labelWidth: '100px'
      },
      formItems: [
        { prop: 'name', label: 'Router', type: 'input', style: 'width:80%', options: { placeholder: 'Please input route name' } },
        { prop: 'title', label: 'Title', type: 'input', style: 'width:80%', options: { placeholder: 'Please input title' } },
        { prop: 'path', label: 'path', type: 'input', style: 'width:80%', options: { placeholder: 'Please input path' } },
        { prop: 'component', label: 'component', type: 'input', style: 'width:80%', options: { placeholder: 'view/xxx.vue or view/xxx/index.vue ' } },
        { prop: 'icon', label: 'icon', slot: 'icon', style: 'width:80%' },
      ],
    })

    return {
      onSubmit,
      ...toRefs(tableState),
      ...toRefs(modalState),
      ...toRefs(formState)
    }
  }
})
</script>
<style lang='scss' scoped>

</style>
