<script setup>
import { reactive } from 'vue';
import TableOfKeyValue from '../TableOfKeyValue.vue';
import TableOfTolerations from '../TableOfTolerations.vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)

const data = reactive({
  nodeLabelsList:[],
  // 容忍操作选项
  tolerationsOperatorSelectOptions : [
    {value: 'Equal', label: '等值'},
    {value: 'Exists',label: '存在'}
  ],
  // 容忍影响选项
  tolerationsEffectSelectOptions : [
    {label: "禁止调度", value: "NoSchedule"},
    {label: "驱逐", value: "NoExecute"},
    {label: "尽量不调度", value: "PreferNoSchedule"}
  ]
})

defineExpose({
    data
})

// 节点标签
const addNodeLabelItem = () => {data.nodeLabelsList.unshift({key:"",value:""})}
const deleteNodeLabelItem = (index) => {data.nodeLabelsList.splice(index,1)}
// 容忍标签
const addTolerationItem = () => {workLoadItem.value.item.spec.template.spec.tolerations.unshift({key:"",value:""})}
const deleteTolerationItem = (index) => {workLoadItem.value.item.spec.template.spec.tolerations.splice(index,1)}
</script>

<template>
  <el-tabs tab-position="left" style="height: 560px;" class="no-border-input">
    <el-tab-pane label="节点选择">
      <TableOfKeyValue
      :table-list="data.nodeLabelsList"
      @add-table-row="addNodeLabelItem"
      @delete-table-row="deleteNodeLabelItem"                  
      />
    </el-tab-pane>
    <el-tab-pane label="污点容忍">
      <table-of-tolerations 
      :tolerations-list="workLoadItem.item.spec.template.spec.tolerations"
      :tolerations-operator-select-options="data.tolerationsOperatorSelectOptions"
      :tolerations-effect-select-options="data.tolerationsEffectSelectOptions"
      @add-toleration="addTolerationItem"
      @delete-toleration="deleteTolerationItem" 
      />
    </el-tab-pane>
  </el-tabs>
</template>

<style>
</style>