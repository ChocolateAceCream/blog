/*
* @fileName debounce.js
* @author Di Sheng
* @date 2023/08/03 09:07:34
* @description: used for debounce button click event
* @usage: <el-button type="primary" @click="onSubmit" v-debounce.once="{ delay:1000 }">
*  @param delay: debounce interval, default is 1000ms
*  @param once: only applied debounce once until element reloaded
*/

const VDebounce = {
  mounted: (el, binding) => {
    el.handler = () => {
      const { delay = 1000 } = binding.value || {}
      el.disabled = true
      const { once } = binding.modifiers
      if (once) return
      setTimeout(() => {
        el.disabled = false
      }, delay)
    }
    el.addEventListener('click', el.handler)
  },
  unmounted: (el) => {
    el.removeEventListener('click', el.handler)
  }
}
export default VDebounce
