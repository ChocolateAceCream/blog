<template>
  <div
    class="main-body"
    :class="{'is-full': isFullContainer}"
  >
    <breadcrumb />
    <el-row>
      <el-card>
        <router-view />
      </el-card>
    </el-row>
  </div>
</template>

<script>
import Breadcrumb from './Breadcrumb.vue'
import { reactive, toRefs, defineComponent, computed } from 'vue'
import { useRoute } from 'vue-router'
export default defineComponent({
  components: {
    Breadcrumb
  },
  setup(props, ctx) {
    const router = useRoute()
    const state = reactive({
      isFullContainer: computed(() => {
        return !!router.meta.fullContainer
      })
    })
    return { ...toRefs(state) }
  }

})

</script>

<style lang="scss" scoped>
  .main-body {
    flex: 1;
    overflow: hidden;

    &.is-full {
      padding: 0;
      margin: 0;
    }
  }
</style>
