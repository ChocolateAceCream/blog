<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/06/20 21:01:23
* @description
!-->
<template>
  <div class="container">
    <MdEditor
      ref="editorRef"
      v-model="text"
      v-bind="$attrs"
      :editor-id="editorId"
      :language="locale"
      :theme="theme"
      :preview-theme="previewTheme"
      :code-theme="codeTheme"
      :toolbars="toolbar"
      :footers="['markdownTotal', '=', 0, 'scrollSwitch']"
      show-code-row-number
      auto-detect-code
      @onUploadImg="uploadImg"
    >
      <!-- <template #defToolbars>
        <MarkExtension :on-insert="insert" />
        <EmojiExtension :on-insert="insert" />
        <ReadExtension :md-text="state.text" />
      </template>
      <template #defFooters>
        <TimeNow />
      </template> -->
    </MdEditor>
    <!-- <span class="tips-text">
      {{ tips
      }}<a
        href="https://github.com/imzbf/md-editor-v3/tree/docs/src/components"
        target="_blank"
      >components</a>
    </span> -->
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from 'vue'
import { MdEditor } from 'md-editor-v3'
import { toolbarConfig } from './toolbarConfig.js'
import 'md-editor-v3/lib/style.css'
import { useSessionStore } from '@/stores/sessionStore'
import { postUploadFile } from '@/api/oss'
export default defineComponent({
  components: { MdEditor },
  emits: ['update:modelValue'],
  setup(props, ctx) {
    const state = reactive({
      text: computed({
        get: () => ctx.attrs.modelValue,
        set: val => {
          ctx.emit('update:modelValue', val)
        }
      }),
      editorId: 'editor',
      locale: 'en-US',
      theme: 'light', // or dark
      previewTheme: 'github', // type: 'default' | 'github' | 'vuepress' | 'mk-cute' | 'smart-blue' | 'cyanosis'
      codeTheme: 'atom', // type: 'atom'|'a11y'|'github'|'gradient'|'kimbie'|'paraiso'|'qtcreator'|'stackoverflow'
      toolbar: toolbarConfig,
    })

    const store = useSessionStore()
    const MdEditorLocaleMapper = {
      'cn': 'zh-CN',
      'en': 'en-US'
    }
    store.$subscribe(async(_, s) => {
      state.locale = MdEditorLocaleMapper[s.userInfo.locale]
    })

    const uploadImg = async(files, callback) => {
      const forms = new FormData()
      forms.append('file', files[0])
      const { data: res } = await postUploadFile(
        forms,
        {
          'Content-Type': 'multipart/form-data'
        }
      )
      callback([res.data.url])
      // callback(res.map(item => item.data.url))
    }
    return {
      uploadImg,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
</style>
