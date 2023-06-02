<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/05/20 21:32:11
* @description
!-->
<template>
  <div class="toolbar">
    <div class="icon-wrapper">
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-add`"
      />
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-favorite`"
      />
      <SvgIcon
        class="toolbar-icon"
        :icon-name="`icon-blog-share`"
      />
    </div>
  </div>
  <div class="card-wrapper">
    <el-card class="abstract-card">
      <div class="card-header">
        <el-link :underline="false">Author</el-link>
        <div class="dividing" />
        <el-link :underline="false">6 days ago</el-link>
        <div class="dividing" />
        <el-link
          class="tag"
          :underline="false"
        >Tag 1</el-link>
        <el-link
          class="tag"
          :underline="false"
        >Tag 2</el-link>
      </div>
      <div class="card-title">title: this is a long long long title.....</div>
      <div class="card-abstract">matlab设置画布大小代码-pyplot-tutorial:代码示例了解
      </div>
      <ul class="card-footer">
        <li>
          <SvgIcon
            class="toolbar-icon"
            :icon-name="`icon-blog-watching`"
          />
          <span> 198</span>
        </li>
        <li>
          <SvgIcon
            class="toolbar-icon"
            :icon-name="`icon-blog-thumb-up`"
          />
          <span> 198</span>
        </li>
        <li>
          <SvgIcon
            class="toolbar-icon"
            :icon-name="`icon-blog-comments`"
          />
          <span> 198</span>
        </li>
      </ul>
    </el-card>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from 'vue'
export default defineComponent({
  setup(props, ctx) {
    const state = reactive({
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>
.toolbar {
  display: flex;
  flex-direction: row-reverse;
  height: 30px;

  .toolbar-icon {
    margin-left: 20px;
    cursor: pointer;
  }
}

.card-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .card-header {
    display: flex;
    // justify-content: space-between;
    align-items: center;
    margin-bottom: 8px;
    cursor: pointer;

    .dividing {
      width: 1px;
      height: 14px;
      background: $lite-background;
      margin: 0 8px;
    }

    .tag {
      position: relative;
      flex-shrink: 0;
      font-size: 13px;
      line-height: 22px;
      padding: 0 8px;

      &:not(:last-child):after {
        position: absolute;
        right: -1px;
        display: block;
        content: " ";
        width: 2px;
        height: 2px;
        border-radius: 50%;
        background-color: $dark-brown;
      }
    }
  }

  .card-title {
    font-weight: 600;
    font-size: 18px;
    line-height: 24px;
    width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-bottom: 8px;
  }

  .card-abstract {
    font-size: 14px;
    line-height: 22px;
    overflow: hidden;
    margin-bottom: 8px;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .card-footer {
    display: flex;
    align-items: center;
    padding: 0;
    margin: 0;
    font-size: 12px;

    li {
      list-style: none;
      margin-right: 20px;
    }
  }
}

.abstract-card {
  @include mobile-device {
    width: 100%
  }

  @include desktop-device {
    width: 50%
  }
}
</style>
