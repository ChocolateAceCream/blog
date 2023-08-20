<template>
  <div class="auth-wrapper">
    <router-view :key="path" />
  </div>
</template>

<script>
import { reactive, toRefs, defineComponent } from 'vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  setup(props, ctx) {
    const router = useRoute()
    const state = reactive({
      path: router.path,
    })
    return { ...toRefs(state) }
  },
})
</script>

<style lang="scss" scoped>
.auth-wrapper {
  position: relative;
  height: 100%;
  width: 100%;
  background-image: linear-gradient(315deg,
        hsl(240deg 100% 20%) 0%,
        hsl(230deg 54% 33%) 17%,
        hsl(211deg 49% 42%) 27%,
        hsl(193deg 52% 49%) 34%,
        hsl(181deg 70% 52%) 39%,
        hsl(180deg 86% 58%) 44%,
        hsl(200deg 88% 67%) 48%,
        hsl(222deg 92% 70%) 51%,
        hsl(242deg 96% 70%) 55%,
        hsl(255deg 99% 59%) 59%,
        hsl(254deg 86% 61%) 64%,
        hsl(241deg 57% 67%) 69%,
        hsl(202deg 37% 62%) 76%,
        hsl(156deg 56% 61%) 85%,
        hsl(152deg 100% 50%) 100%);
}
</style>
