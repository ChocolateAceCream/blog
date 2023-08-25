<script>
import { ElConfigProvider } from 'element-plus'
import { defineComponent, reactive, toRefs, inject, provide, nextTick, watch } from 'vue'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import en from 'element-plus/lib/locale/lang/en'
import { useSessionStore } from '@/stores/sessionStore'
import { useI18n } from 'vue-i18n'
import 'dayjs/locale/zh-cn'
export default defineComponent({
  components: {
    ElConfigProvider,
  },
  setup() {
    const elementPlusLocaleMapper = {
      'cn': zhCn,
      'en': en
    }
    const dayjsLocaleMapper = {
      'cn': 'zh-cn',
      'en': 'en'
    }
    const { locale } = useI18n()

    const store = useSessionStore()

    const dayjs = inject('dayjs')
    dayjs.locale(dayjsLocaleMapper[store.userInfo.locale])

    watch(
      () => store.userInfo.locale,
      (newValue, _) => {
        state.locale = elementPlusLocaleMapper[newValue]
        locale.value = newValue
        dayjs.locale(dayjsLocaleMapper[newValue])
        reload()
      }
    )

    const state = reactive({
      locale: zhCn,
      isRouterAlive: true,
    })
    const reload = () => {
      state.isRouterAlive = false
      nextTick(() => {
        state.isRouterAlive = true
      })
    }
    provide('reload', reload)
    return {
      reload,
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
    <router-view
      v-if="isRouterAlive"
      :key="$route.fullPath"
    />
  </el-config-provider>

</template>
