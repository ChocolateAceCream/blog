<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/02/23 20:16:48
* @description
!-->
<template>
  <div>Home</div>
  <button @click="pushRoute">push route</button>
  <div>Home</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
import { useRouterStore } from '@/stores/routerStore'
import router from '@/router'
export default defineComponent({

  setup(props, ctx) {
    const pushRoute = () => {
      const routerStore = useRouterStore()

      console.log('-------routerStore.routerTree---', routerStore.routerTree)
      // router.addRoute(routerStore.routerTree)
      console.log('-------router---', router)
      router.push({name: 'role'})
    }
    console.log('-------router---', router)

    const state = reactive({
    })
    return {
      pushRoute,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
