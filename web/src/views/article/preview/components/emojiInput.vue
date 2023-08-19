<!--
* @fileName commentInput.vue
* @author Di Sheng
* @date 2023/08/04 15:23:01
* @description comment input component
!-->
<template>
  <div class="comment-wrapper">
    <el-input
      ref="inputRef"
      v-model="input"
      v-keyboard:[onSubmit]="['ctrlKey', 13]"
      show-word-limit
      type="textarea"
      style="max-width:800px"
    />
    <div class="action-box">
      <el-button
        link
        type="primary"
        @click="toggleEmojiPicker"
      >
        <SvgIcon
          :icon-name="`icon-blog-Emoji`"
          size="20px"
        /> &nbsp;Emoji
      </el-button>
      <div class="submit">
        <span> Ctrl + Enter</span>
        <el-button
          v-debounce
          type="primary"
          @click="onSubmit"
        >Submit</el-button>
      </div>
    </div>

    <Picker
      v-if="showEmojiPicker"
      :set="`twitter`"
      :data="emojiIndex"
      @select="addEmoji"
    />
  </div>

</template>

<script>
import { defineComponent, toRefs, reactive, nextTick } from 'vue'
import { Picker, EmojiIndex } from 'emoji-mart-vue-fast/src'
import data from 'emoji-mart-vue-fast/data/all.json'
import { insertAt } from '@/utils/stringFun.js'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import _ from 'lodash'
export default defineComponent({
  name: 'EmojiInput',
  components: {
    Picker,
  },
  emits: ['submit'],
  setup(props, ctx) {
    const state = reactive({
      inputRef: null,
      input: '',
      showEmojiPicker: false,
    })
    const emojiIndex = new EmojiIndex(data)
    const addEmoji = async(emoji) => {
      const e = state.inputRef.ref
      state.input = insertAt(state.input, emoji.native, e.selectionStart, e.selectionEnd)
      const pos = e.selectionStart + emoji.native.length
      e.focus()
      await nextTick()
      e.setSelectionRange(pos, pos)
    }
    const toggleEmojiPicker = () => {
      state.showEmojiPicker = !state.showEmojiPicker
    }
    const onSubmit = async() => {
      ctx.emit('submit', state.input)
    }

    const reset = () => {
      state.inputRef.blur()
      state.inputRef.clear()
    }
    return {
      addEmoji,
      onSubmit,
      reset,
      toggleEmojiPicker,
      emojiIndex,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.comment-wrapper {
  display: block;
  max-width: 650px;

  .action-box {
    display: flex;
    align-items: center;
    margin-top: 8px;
  }

  .submit {
    span {
      font-size: 14px;
      line-height: 22px;
      letter-spacing: 0.2px;
      margin-right: 16px;
      color: $lite-grey;
    }

    flex: 0 0 auto;
    margin-left:auto;
  }
}</style>
