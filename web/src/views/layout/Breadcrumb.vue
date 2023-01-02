<template>
  <el-breadcrumb separator="/" class="breadcrumb fl">
    <el-breadcrumb-item v-for="(item, index) in levelList" :key="index" :to="{ name: item.name, params: $route.params }">
      <span class="item">{{ item.meta.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
import _ from 'lodash'
import { reactive, toRefs, defineComponent, computed } from 'vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  setup(props, ctx) {
    const route = useRoute()
    const levels = route.matched.filter(item => item.name)
    console.log('-------levels--------', levels)
    const state = reactive({
      levelList: computed(() => {
        const levels = route.matched.filter(item => item.name)
        const indexForLevelPos = _.findLastIndex(levels, (item) => item.meta?.indexFor)
        if (indexForLevelPos >= 0) {
          const indexForLevel = levels[indexForLevelPos]
          const parentLevelPos = _.findLastIndex(levels, (item) => item.name === indexForLevel.meta.indexFor)
          if (parentLevelPos >= 0) {
            return _.slice(levels, parentLevelPos, indexForLevelPos - parentLevelPos)
          }
        }
        return levels
      })
    })
    return { ...toRefs(state) }
  }
})
</script>

<style lang="scss" scoped>
.item {
  line-height: 31px;
}
</style>
