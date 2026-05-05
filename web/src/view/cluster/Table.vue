<script lang="ts" setup>
  import { CircleCheck, CircleClose } from '@element-plus/icons-vue'
  import { getListHandler } from '../../api/generic'
  import { computed, onBeforeMount, reactive } from 'vue'

  const props = defineProps(['tableData'])
  const emit = defineEmits(['getItem', 'deleteItem'])

  const data = reactive({
    nodeNum: '3',
    podNum: '5',
  })
  // const getNodeNum = (clusterId) => {
  //   getListHandler(clusterId, '', 'node', '').then((res) => {
  //     if (res.data.status == 200) {
  //       data.nodeNum = res.data.data.items.length
  //     }
  //   })
  // }
</script>

<template>
  <el-table :data="props.tableData" style="width: 100%" height="1010px">
    <el-table-column label="ID" prop="clusterId">
      <template #default="scope">
        <el-button type="primary" link @click="emit('getItem', scope.row)">
          {{ scope.row.clusterId }}</el-button
        >
      </template>
    </el-table-column>
    <el-table-column label="名称" prop="clusterAlias" />
    <el-table-column label="版本" prop="clusterVersion" />
    <el-table-column label="位置" prop="clusterLocation" />
    <el-table-column label="状态" prop="clusterStatus">
      <template #default="scope">
        <el-icon v-if="scope.row.clusterStatus == 'Active'" color="green"><CircleCheck /></el-icon>
        <el-icon v-else color="red"><CircleClose /></el-icon>
      </template>
    </el-table-column>
    <el-table-column label="节点数量" prop="clusterNodeNum" />
    <el-table-column label="Pod数量" prop="clusterPodNum" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="warning" @click="emit('getItem', scope.row)" link> 编辑 </el-button>
        <el-button type="danger" @click="emit('deleteItem', scope.row)" link> 删除 </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>
