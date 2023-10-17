/*
* @fileName mouse.js
* @author Di Sheng
* @date 2023/10/17 09:45:04
* @description composition API for getting mouse x,y
*/

/* @usage

<template>
  <div>{{ t.x }}, {{ t.y }}</div>
</template>

import {useMouse} from '@/utils/mouse.js'
export default defineComponent({
  setup(props, ctx) {
    const t = useMouse()
    return {
      t
    }
  }
})

*/

import { toRefs, reactive, onMounted, onUnmounted } from 'vue'
export const useMouse = () => {
  const state = reactive({ x: 0, y: 0 })
  const update = (e) => {
    state.x = e.pageX
    state.y = e.pageY
  }
  onMounted(() => {
    window.addEventListener('mousemove', update)
  })
  onUnmounted(() => {
    window.removeEventListener('mousemove', update)
  })
  return toRefs(state)
}

