<template>
  <el-table
    v-bind="$attrs"
    @selection-change="handleSelectionChange"
    @select-all="handleSelectionChange"
    style="width: 100%; margin-bottom: 20px"
  >
    <el-table-column v-if="selectable" type="selection" width="55" />
    <template v-for="(item, idx) in config" :key="idx">
      <el-table-column v-bind="item">
        <template v-if="item.headerSlot" #header>
          <slot :name="item.headerSlot" />
        </template>
        <template v-if="item.bodySlot" #default="scope">
          <slot :name="item.bodySlot" v-bind="scope" />
        </template>
      </el-table-column>
    </template>
  </el-table>
</template>
<script>
import { useSlots, computed, h, resolveDynamicComponent, defineComponent } from "vue";
import { ElTableColumn } from "element-plus";

export default defineComponent({
  props: {
    config: {
      type: Array,
      default: () => [],
    },
    selectable: {
      type: Boolean,
      default: false,
    },
  },
  setup(props, ctx) {
    const handleSelectionChange = (val) => {
      ctx.emit("select", val);
    };
    const handleSelectAll = () => {
      console.log("xxx");
    };
    const slots = computed(() => {
      // return slotsOrigin.default()
    });
    const column = resolveDynamicComponent(ElTableColumn);
    const result = h(column, { label: "Date", prop: "date" });

    const children = () => [result];
    return {
      // useSlots,
      // computed,
      // h,
      // resolveDynamicComponent,
      // ElTableColumn,
      // slotsOrigin,
      // slots,
      // column,
      // result,
      // children,
      handleSelectAll,
      handleSelectionChange,
    };
  },
});
</script>
