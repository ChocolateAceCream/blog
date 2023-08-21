<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/04/02 08:31:38
* @description role page
!-->
<template>
  <el-button
    type="primary"
    style="margin-bottom: 15px;"
    @click="onAddRole(0)"
  > &#43; Add Root Role</el-button>
  <my-table
    :data="tableData"
    :config="tableConfig"
    row-key="id"
  >
    <template #operationBody="scope">
      <el-button
        type="primary"
        link
        @click="onAddRole(scope.row.id)"
      >Add</el-button>
      <el-button
        type="primary"
        link
        @click="onEditRole(scope.row)"
      >Edit</el-button>
      <el-button
        type="primary"
        link
        @click="onMenuModalOpen(scope.row.id)"
      >Menu</el-button>
      <el-button
        type="primary"
        link
        @click="onEndpointModalOpen(scope.row.id)"
      >Endpoint</el-button>
      <el-button
        type="primary"
        link
        @click="onDeleteRole(scope.row)"
      >Delete</el-button>
    </template>
  </my-table>
  <Modal
    ref="roleModalRef"
    width="550px"
    :title="roleModalType + ' Role'"
    @close="onRoleModalClose"
    @confirm="onRoleModalConfirm"
  >
    <MyForm
      ref="formRef"
      v-model:formData="formData"
      :config="formConfig"
      :form-items="formItems"
    />
  </Modal>
  <MenuModal
    ref="menuModalRef"
  />
  <EndpointModal
    ref="endpointModalRef"
  />
</template>

<script>
import { defineComponent, toRefs, reactive, onMounted } from 'vue'
import { postAddRole, getRoleList, deleteRole, putEditRole } from '@/api/role'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouterStore } from '@/stores/routerStore'
import MenuModal from './components/MenuModal'
import EndpointModal from './components/EndpointModal'
import _ from 'lodash'
export default defineComponent({
  components: {
    MenuModal,
    EndpointModal
  },
  setup(props, ctx) {
    const tableState = reactive({
      tableData: [],
      tableConfig: [
        { label: 'ID', prop: 'id' },
        { label: 'Name', prop: 'name' },
        { label: 'Operation', bodySlot: 'operationBody' },
      ],
      onAddRole(id) {
        modalState.roleModalType = 'Add'
        formState.formData = { pid: id }
        modalState.onRoleModalOpen()
      },
      onEditRole(row) {
        modalState.roleModalType = 'Edit'
        formState.formData = _.cloneDeep(row)
        modalState.onRoleModalOpen()
      },
      onDeleteRole(row) {
        const payload = []
        const helper = (node) => {
          payload.push(node.id)
          node.children.forEach(child => {
            helper(child)
          })
        }
        helper(row)
        ElMessageBox.confirm(
          'Delete role will also delete all child role and role-menu, casbin rules and role-user relations. Continue?',
          'Warning',
          {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
            type: 'warning',
          }
        )
          .then(async() => {
            const { data: res } = await deleteRole({ data: { id: payload } })
            if (res.errorCode === 0) {
              ElMessage({
                type: 'success',
                message: 'Delete completed',
              })
              fetchRoleList()
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
    const fetchRoleList = async() => {
      const { data: res } = await getRoleList()
      if (res.errorCode === 0) {
        tableState.tableData = formatRoleTree(res.data)
      }
    }

    onMounted(fetchRoleList)

    const formatRoleTree = (roles) => {
      const mapper = {}
      roles.forEach(role => {
        if (mapper[role.pid]) {
          mapper[role.pid].push(role)
        } else {
          mapper[role.pid] = [role]
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
        name: root.name,
        id: root.roleId,
        pid: root.pid,
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
        let resp = {}
        if (modalState.roleModalType === 'Add') {
          resp = await postAddRole(formState.formData)
        } else {
          resp = await putEditRole(formState.formData)
        }
        const { data: res } = resp
        if (res.errorCode === 0) {
          ElMessage({
            message: `${modalState.roleModalType} Role Success`,
            type: 'success',
            duration: 3 * 1000
          })
          fetchRoleList()
          modalState.roleModalRef.closeModal()
        }
      } catch (err) {
        console.log('-----form validation err-', err)
      }
    }
    const modalState = reactive({
      roleModalRef: null,
      roleModalType: 'Add',
      onRoleModalOpen() {
        modalState.roleModalRef.openModal()
        formState.formRef?.clearAllValidate()
      },
      onRoleModalClose() {
        modalState.roleModalRef.closeModal()
      },
      onRoleModalConfirm: _.throttle(onSubmit, 2000),
      menuModalRef: null,
      onMenuModalOpen(roleId) {
        modalState.menuModalRef.onModalOpen(roleId)
      },
      endpointModalRef: null,
      onEndpointModalOpen(roleId) {
        modalState.endpointModalRef.onModalOpen(roleId)
      },
    })

    const formState = reactive({
      formData: {},
      formRef: null,
      formConfig: {
        rules: {
          name: [{ required: true, message: 'role name required', trigger: 'blur' }],
        },
        labelPosition: 'right',
        labelWidth: '100px'
      },
      formItems: [
        { prop: 'name', label: 'Name', type: 'input', style: 'width:80%', options: { placeholder: 'Please input role name' } },
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
<style lang='scss' scoped></style>
