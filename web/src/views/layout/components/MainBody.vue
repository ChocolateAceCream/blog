<template>
  <div
    class="main-body"
    :class="{'is-full': isFullContainer}"
  >
    <el-card
      style="height:100%"
      :body-style="{height: '100%'}"
    >
      <el-scrollbar height="100%">
        <!-- 避免了组件复用，子路由间来回切换时，页面都会重新加载. -->
        <!-- <router-view :key="path" /> -->
        <router-view />
      </el-scrollbar>
    </el-card>
  </div>
</template>

<script>
import { reactive, toRefs, defineComponent, computed } from 'vue'
import { useRoute } from 'vue-router'
export default defineComponent({
  setup(props, ctx) {
    const router = useRoute()
    const state = reactive({
      isFullContainer: computed(() => {
        return !!router.meta.fullContainer
      }),
      path: router.path,
    })
    return { ...toRefs(state) }
  }

})

</script>

<style lang="scss" scoped>
  .main-body {
    flex: 1;
    height: 100%;
    overflow: hidden;

    &.is-full {
      padding: 0;
      margin: 0;
    }
  }
</style>
