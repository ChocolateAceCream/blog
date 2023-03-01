/<!--
* @fileName index.vue
* @author Di Sheng
* @date 2022/03/17 14:07:47
* @description
!-->
<template>
  <slot
    v-if="permission"
    v-bind="{permission}"
  />
</template>

<script>
import { computed, reactive, toRefs, defineComponent } from 'vue'
import { useSessionStore } from '@/stores/sessionStore'

export default defineComponent({
  props: {
    value: [String],
  },
  setup(props, ctx) {
    const state = reactive({
      permission: computed(() => {
        return useSessionStore().hasPermission(props.value)
      })
    })
    return { ...toRefs(state) }
  },
})
</script>

<style lang="scss" scoped>
</style>
