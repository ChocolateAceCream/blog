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
        @click="onDeleteMenu(scope.row)"
      >Delete</el-button>
    </template>
  </my-table>
  <Modal
    ref="modalRef"
    width="550px"
    :title="modalType + ' Menu'"
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
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { postAddMenu, getMenuList, deleteMenu, putEditMenu } from '@/api/menu'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouterStore } from '@/stores/routerStore'
import _ from 'lodash'
import Icon from './icon.vue'
export default defineComponent({
  components: {
    Icon
  },
  setup(props, ctx) {
    const tableState = reactive({
      tableData: [],
      tableConfig: [
        { label: 'ID', prop: 'id' },
        { label: 'Route', prop: 'name' },
        { label: 'Path', prop: 'path' },
        { label: 'Title', prop: 'meta.title' },
        { label: 'Icon', prop: 'icon', bodySlot: 'iconBody' },
        { label: 'Operation', bodySlot: 'operationBody' },
      ],
      onAddMenu(id) {
        modalState.modalType = 'Add'
        formState.formData = { pid: id }
        modalState.onModalOpen()
      },
      onEditMenu(row) {
        modalState.modalType = 'Edit'
        formState.formData = _.cloneDeep(row)
        modalState.onModalOpen()
      },
      onDeleteMenu(row) {
        const payload = []
        const helper = (node) => {
          payload.push(node.id)
          node.children.forEach(child => {
            helper(child)
          })
        }
        helper(row)
        ElMessageBox.confirm(
          'Delete menu will also delete all child menu and role-menu relations. Continue?',
          'Warning',
          {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
            type: 'warning',
          }
        )
          .then(async() => {
            const { data: res } = await deleteMenu({ data: { id: payload } })
            if (res.errorCode === 0) {
              ElMessage({
                type: 'success',
                message: 'Delete completed',
              })
              fetchMenuList()
              const routerStore = useRouterStore()
              routerStore.asyncRouterFlag = 0
              routerStore.setAsyncRouter()
            }
          })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: 'Delete canceled',
            })
          })
      }
    })
    onMounted(() => {
      fetchMenuList()
    })
    const fetchMenuList = async() => {
      const { data: res } = await getMenuList()
      if (res.errorCode === 0) {
        tableState.tableData = formatMenuTree(res.data)
      } else {
        ElMessage({
          message: res.msg,
          type: 'error',
          duration: 3 * 1000
        })
      }
    }
    const formatMenuTree = (menus) => {
      const mapper = {}
      menus.forEach(menu => {
        if (mapper[menu.pid]) {
          mapper[menu.pid].push(menu)
        } else {
          mapper[menu.pid] = [menu]
        }
      })
      let r = []
      if (mapper[0] && mapper[0].length > 0) {
        r = mapper[0].map(node => {
          return treeHelper(node, mapper)
        })
      }
      return r
    }

    const treeHelper = (root, mapper) => {
      const temp = {
        path: root.path,
        name: root.name,
        component: root.component,
        id: root.id,
        pid: root.pid,
        meta: root.meta,
        icon: root.meta.icon,
        title: root.meta.title,
        children: []
      }
      if (mapper[temp.id]) {
        mapper[temp.id].forEach(node => {
          const child = treeHelper(node, mapper)
          temp.children.push(child)
        })
      }
      return temp
    }

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
        let resp = {}
        if (modalState.modalType === 'Add') {
          resp = await postAddMenu(payload)
        } else {
          resp = await putEditMenu(payload)
        }
        const { data: res } = resp
        if (res.errorCode === 0) {
          ElMessage({
            message: `${modalState.modalType} Menu Success`,
            type: 'success',
            duration: 3 * 1000
          })
          fetchMenuList()
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
