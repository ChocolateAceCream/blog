<template>
  <div class="main-body" :class="{'is-full': isFullContainer}">
    <el-card>
    <router-view></router-view>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'AppMain',
  computed: {
    cachedViews() {
      return []
    },
    key() {
      return this.$route.fullPath
    },
    isFullContainer() {
      return !!this.$route.meta.fullContainer
    }
  }
}
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
