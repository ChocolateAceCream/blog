<!--
* @fileName menuItem.vue
* @author Di Sheng
* @date 2023/03/10 09:20:09
* @description
!-->
<script>
import { defineComponent, toRefs, reactive, computed, h } from 'vue'
import SvgIcon from '@/shared/components/SvgIcon'
import { ElMenuItem, ElSubMenu } from 'element-plus'
export default defineComponent({
  name: 'MenuItem',
  props: {
    routeInfo: {
      type: Object,
      default: () => {}
    }
  },
  setup(props, ctx) {
    const state = reactive({
    })

    const menuComponent = computed(() => {
      return true
      // if (props.routeInfo.children?.length) {
      // }
    })
    return {
      menuComponent,
      ...toRefs(state)
    }
  },
  render() {
    const svg = (node) => {
      return h(SvgIcon, {
        iconName: 'icon-blog-' + node.icon,
        color: '#3498db',
        size: '15px',
      })
    }
    const renderHelper = (root) => {
      if (root.children?.length) {
        const children = []
        root.children.forEach(node => {
          children.push(renderHelper(node))
        })
        return h(ElSubMenu, { index: root.routeName, class: 'menu-item' }, {
          // slot content
          title: () => {
            return [svg(root), h('span', null, root.title)]
          },
          default: () => {
            return children
          }
        })
      } else {
        return h(ElMenuItem, { index: root.routeName, class: 'menu-item' }, {
          // slot content
          title: () => {
            return h('span', null, root.title)
          },
          default: () => {
            return svg(root)
          }
        })
      }
    }
    return renderHelper(this.routeInfo)
  }
})
</script>
<style lang='scss' scoped>
.menu-item {
  background-color: $liter-background;
}
</style>
