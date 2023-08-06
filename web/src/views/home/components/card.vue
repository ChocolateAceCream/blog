<!--
* @fileName card.vue
* @author Di Sheng
* @date 2023/07/10 21:31:37
* @description
!-->
<template>
  <el-card
    class="abstract-card"
    @click="onCardClick(articleInfo.id)"
  >
    <div class="card-header">
      <el-link :underline="false">{{ articleInfo.author }}</el-link>
      <div class="dividing" />
      <el-link :underline="false">{{ timestamp }}</el-link>
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
    <div class="card-title">{{ articleInfo.title || 'N/A' }}</div>
    <div class="card-abstract">{{ articleInfo.abstract || 'N/A' }}     </div>
    <ul class="card-footer">
      <li>
        <SvgIcon
          class="toolbar-icon"
          :icon-name="`icon-blog-watching`"
        />
        <span> {{ articleInfo.viewedTimes }}</span>
      </li>
      <li>
        <SvgIcon
          class="toolbar-icon"
          :icon-name="`icon-blog-thumb-up-line`"
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
</template>

<script>
import { defineComponent, toRefs, reactive, inject, computed } from 'vue'
import router from '@/router'
export default defineComponent({
  name: 'Card',
  props: {
    articleInfo: {
      type: Object,
      default: () => {
        return {
          articleId: null,
          content: null,
          authorId: null,
          viewedTimes: 0,
        }
      }
    }
  },
  setup(props, ctx) {
    const dayjs = inject('dayjs')
    const state = reactive({
      articleInfo: { ...props.articleInfo },
      onCardClick: (id) => {
        router.push({ path: 'article/' + id })
      },
      timestamp: computed(() => {
        return dayjs(state.articleInfo.updatedAt).fromNow()
      })
    })
    return {
      ...toRefs(state)
    }
  }
})
</script>
<style lang='scss' scoped>

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
    .toolbar-icon{
      margin-right: 5px;
    }
  }

.abstract-card {
  margin:2px;
  @include mobile-device {
    width: 100%
  }

  @include desktop-device {
    width: calc(50% - 4px);
  }
}
</style>
