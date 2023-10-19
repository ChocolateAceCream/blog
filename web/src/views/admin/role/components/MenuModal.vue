<!--
* @fileName MenuModal.vue
* @author Di Sheng
* @date 2023/04/11 08:17:39
* @description
!-->
<template>
  <Modal
    ref="modalRef"
    width="550px"
    title="Assign Role Menus"
    @close="onModalClose"
    @confirm="onModalConfirm"
  >
    <MyTree
      ref="treeRef"
      :data="treeData"
      :props="defaultProps"
      :selected="menuIdList"
      node-key="id"
      show-checkbox
      default-expand-all
    />
  </Modal>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import { getRoleMenuTree, assignRoleMenus } from '@/api/menu'
//import { ElMessage } from 'element-plus'
import { useRouterStore } from '@/stores/routerStore'
import { throttle, filter } from 'lodash-es'

import ListToTree from '@/utils/tree'
export default defineComponent({
  name: 'MenuModal',
  setup(props, ctx) {
    const state = reactive({
      treeData: [],
      treeRef: null,
      menuIdList: [],
      selectedRoleId: null,
      defaultProps: {
        children: 'children',
        label: 'name',
      }
    })
    const onSubmit = async() => {
      const selectedMenus = state.treeRef.getSelectKeys()
      const { data: res } = await assignRoleMenus({ id: state.selectedRoleId, menus: selectedMenus })
      if (res.errorCode === 0) {
        ElMessage({
          message: res.msg,
          type: 'success',
          duration: 3 * 1000
        })
        modalState.onModalClose()
        useRouterStore().updateAsyncRouter()
      }
    }
    const fetchMenuTree = async(roleId) => {
      state.selectedRoleId = roleId
      const { data: res } = await getRoleMenuTree({ id: roleId })
      if (res.errorCode === 0) {
        const { root, mapper } = ListToTree(res.data.menuList)
        state.treeData = root.children
        const leafNodes = filter(res.data.roleMenus, (menu) => { return mapper[menu.id] === undefined })
        const temp = []
        leafNodes.forEach(menu => temp.push(menu.id))
        state.treeRef.setKeys(temp)
        state.menuIdList = temp
      }
    }
    const modalState = reactive({
      modalRef: null,
      async onModalOpen(roleId) {
        fetchMenuTree(roleId)
        modalState.modalRef.openModal()
      },
      onModalClose() {
        modalState.modalRef.closeModal()
      },
      onModalConfirm: throttle(onSubmit, 2000)
    })

    return {
      ...toRefs(state),
      ...toRefs(modalState)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
