<script setup>
  import { computed } from 'vue'

  const props = defineProps(['tableData'])
  const emit = defineEmits(['deleteItem'])

  // 搜索后的 namespace 列表数据
  const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
      (item) =>
        !props.tableData.search ||
        item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
  )
</script>
<template>
  <el-table :data="filterTableData" style="width: 100%" height="70vh">
    <el-table-column label="名称" prop="metadata.name" />
    <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
    <el-table-column label="状态" prop="status.phase" />
    <el-table-column label="操作" prop="operation">
      <template #default="scope">
        <el-button link type="danger" @click="emit('deleteItem', scope.row.metadata.name)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>
