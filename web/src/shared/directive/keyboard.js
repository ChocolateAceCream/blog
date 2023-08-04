/*
* @fileName keyboard.js
* @author Di Sheng
* @date 2023/08/03 09:07:34
* @description: used for detect keyboard input
* @usage:
    <el-input
      ref="inputRef"
      v-model="input"
      v-keyboard:[hotKeyHandler]="['ctrlKey',13]"
      show-word-limit
      type="textarea"
      style="max-width:800px"
    />
*  @param ['ctrlKey',13]: first one is hotkey, alternatives are null (if no key combination needed), altKey, shiftKey, metaKey etc, second param is the actual key code of certain keyboard keys
*  @param [hotKeyHandler]: the method binding.arg(event) will trigger
*/

const VKeyboard = {
  mounted: (el, binding) => {
    el.handler = (event) => {
      const elInput = el.querySelector('input,textarea') // to select the actual input element this way. 'el' is only a wrapper and cannot be focused on
      // elInput === document.activeElement check if the input is focused
      let flag
      if (binding.value[0] != null) {
        flag = event[binding.value[0]] && event.keyCode === binding.value[1] && elInput === document.activeElement
      } else {
        flag = event.keyCode === binding.value[1] && elInput === document.activeElement
      }
      if (flag) {
        binding.arg(event)
      }
    }
    document.addEventListener('keydown', el.handler)
  },
  unmounted: (el) => {
    document.removeEventListener('keydown', el.handler)
  },
}

export default VKeyboard
