<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/04/18 09:01:03
* @description
!-->
<template>
  <el-select
    v-model="method"
    v-bind="rest"
    placeholder="Select"
  >
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
  </el-select>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from 'vue'
export default defineComponent({
  setup(props, ctx) {
    const { options, ...rest } = ctx.attrs
    const state = reactive({
      method: computed({
        get: () => props[`${rest.propName}`],
        set: val => {
          rest.onInput(val)
        }
      })
    })
    return {
      options,
      rest,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
