<script setup>
  const props = defineProps(['containerItem', 'probeType'])
</script>

<template>
  <div style="margin-left: 50px; margin-top: 20px">
    <el-form label-width="100px">
      <el-row :gutter="200">
        <el-col :span="12">
          <el-form-item label="探测服务">
            <el-input v-model="props.containerItem[props.probeType].grpc.service" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="探测端口">
            <el-input v-model="props.containerItem[props.probeType].grpc.port" />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>
