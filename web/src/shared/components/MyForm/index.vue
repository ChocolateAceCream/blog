<!--
* @fileName index.vue
* @author Di Sheng
* @date 2023/03/15 19:03:00
* @description
!-->
<script>
import { defineComponent, toRefs, reactive, h, render, ref, unref } from 'vue'
//import { ElForm, ElFormItem } from 'element-plus'
import {mapper } from './mapper'
export default defineComponent({
  props: {
    config: {
      type: Object,
      default: () => {
        return {
          labelWidth: '85px',
          inline: true,
          labelPosition: 'top'
        }
      }
    },
    formData: {
      type: Object,
      default: () => { }
    },
    formItems: {
      type: Array,
      default: () => []
    }
  },
  emits: ['update:formData'],
  setup(props, ctx) {
    const formRef = ref()
    const validate = () => {
      const form = unref(formRef)
      return form.validate()
    }
    const clearAllValidate = () => {
      const form = unref(formRef)
      return form.clearValidate()
    }
    const reset = () => {
      formRef.value.resetFields()
    }
    const state = reactive({
    })
    return {
      validate,
      clearAllValidate,
      formRef,
      reset,
      ...toRefs(state)
    }
  },
  render() {
    const formAttrs = {
      ref: 'formRef',
      model: this.formData,
      ...this.config
    }
    const attrHelper = (form, field, options) => {
      return {
        modelValue: form[field],
        ...options,
        onInput: (target) => {
          Object.assign(form, { [field]: target })
          console.log('-----form----', form)
          this.$emit('update:formData', form)
        }
      }
    }
    const childrenList = []
    this.formItems.forEach(item => {
      const { prop, slot, type } = item
      let slotContent = {}
      if (slot !== undefined) {
        // pass this.formData as a scope to the slot, so it can be accessible from parent component where slots implemented
        slotContent = { default: () => this.$slots[slot](this.formData) }
      } else {
        slotContent = { default: () => h(mapper[type], attrHelper(this.formData, prop, item.options)) }
      }
      const child = h(ElFormItem, item, slotContent)
      childrenList.push(child)
    })
    return h(ElForm, formAttrs, { default: () => childrenList })
  }
})
</script>
<style lang='scss' scoped>

</style>
