<!--
* @fileName icon.vue
* @author Di Sheng
* @date 2023/03/17 10:40:52
* @description
!-->
<template>
  <div>
    <span class="outer-icon">
      <SvgIcon
        :icon-name="`icon-blog-` + icon"
        size="15px"
      />
    </span>
  </div>
  <div class="selector">
    <el-select
      v-model="tempIcon"
    >
      <el-option
        v-for="(item,idx) in iconMapper"
        :key="idx"
        :label="item"
        :value="item"
      >
        <span>
          <SvgIcon
            :icon-name="`icon-blog-` + item"
            size="15px"
            style="margin-right:10px;"
          />
        </span>
        <span style="text-align: left">{{ item }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from 'vue'
import _ from 'lodash'
import iconJson from '@/assets/iconfont/iconfont.json'
export default defineComponent({
  props: {
    icon: {
      type: String,
      default: ''
    }
  },
  emits: ['update:icon'],
  setup(props, ctx) {
    const iconMapper = _.map(iconJson.glyphs, 'font_class')
    const state = reactive({
      tempIcon: computed({
        get: () => props.icon,
        set: val => {
          ctx.emit('update:icon', val)
        }
      })
    })
    return {
      iconMapper,
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.outer-icon{
  position: absolute;
  z-index: 999;
  top: 2px;
  left: 8px;
  font-size: 14px;
  margin-right: 10px;
}

.selector {
  &:deep(.el-input__inner){
    padding-left:  20px;
  }
}
</style>
