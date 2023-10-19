<!--
* @fileName EndpointModal.vue
* @author Di Sheng
* @date 2023/04/11 08:17:39
* @description
!-->
<template>
  <Modal
    ref="modalRef"
    width="550px"
    title="Assign Role Endpoints"
    @close="onModalClose"
    @confirm="onModalConfirm"
  >
    <MyTree
      ref="treeRef"
      :data="treeData"
      :props="defaultProps"
      :default-checked-keys="selectedIdList"
      show-checkbox
      node-key="id"
      default-expand-all
    />
  </Modal>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import { getCasbinByRoleId, postUpdateCasbin } from '@/api/casbin'
import { getEndpointList } from '@/api/endpoint'
//import { ElMessage } from 'element-plus'
import { throttle, map, groupBy } from 'lodash-es'

export default defineComponent({
  name: 'EndpointModal',
  setup(props, ctx) {
    const state = reactive({
      treeData: [],
      treeRef: null,
      selectedIdList: [],
      selectedRoleId: null,
      defaultProps: {
        children: 'children',
        label: 'name',
      }
    })
    const onSubmit = async() => {
      const payload = state.treeRef.getCheckedNodes(true).map(i => { return { path: i.path, method: i.method } })
      const { data: res } = await postUpdateCasbin({ roleId: state.selectedRoleId, endpoints: payload })
      if (res.errorCode === 0) {
        ElMessage({
          message: res.msg,
          type: 'success',
          duration: 3 * 1000
        })
        modalState.onModalClose()
      }
    }
    const fetchTree = async(roleId) => {
      const { data: casbinRes } = await getCasbinByRoleId({ params: { id: roleId } })
      if (casbinRes.errorCode !== 0) {
        return
      }
      const payload = { pageNumber: 1, pageSize: 1000, orderBy: 'group_name'}
      const { data: endpointRes } = await getEndpointList({ params: payload })
      if (casbinRes.errorCode !== 0) {
        return
      }
      const groups = groupBy(endpointRes.data.list, 'groupName')
      state.treeData = map(groups, (arr, key) => {
        const node = {
          name: key,
          children: []
        }
        arr.forEach(item => {
          node.children.push({name: item.name, id: item.path + ':' + item.method, path: item.path, method: item.method})
        })
        return node
      })
      state.selectedIdList = map(casbinRes.data, (val) => {
        return val.path + ':' + val.method
      })
    }
    const modalState = reactive({
      modalRef: null,
      onModalOpen(roleId) {
        modalState.modalRef.openModal()
        state.selectedRoleId = roleId
        fetchTree(roleId)
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
<style lang='scss' scoped></style>
