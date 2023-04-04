<script>
import { ElConfigProvider } from 'element-plus'
import { defineComponent, reactive, toRefs } from 'vue'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import en from 'element-plus/lib/locale/lang/en'
import { useSessionStore } from '@/stores/sessionStore'
export default defineComponent({
  components: {
    ElConfigProvider,
  },
  setup() {
    const elementPlusLocaleMapper = {
      'cn': zhCn,
      'en': en
    }
    const store = useSessionStore()
    store.$subscribe((_, s) => {
      state.locale = elementPlusLocaleMapper[s.userInfo.locale]
    })

    const state = reactive({
      locale: zhCn,
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>

<template>
  <el-config-provider
    :locale="locale"
    :size="`default`"
    :z-index="3000"
  >
    <router-view />
  </el-config-provider>

</template>
