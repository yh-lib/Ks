<script setup>
  import { ref } from 'vue'
  import exec from './exec.vue'
  import tcpSocket from './TcpSocket.vue'
  import httpGet from './HttpGet.vue'
  import grpc from './Grpc.vue'

  const props = defineProps(['containerItem', 'probeType'])

  const probeMethod = ref('exec')

  const options = [
    {
      value: 'exec',
      label: 'Exec 探测',
    },
    {
      value: 'tcpSocket',
      label: 'TCP 端口探测',
    },
    {
      value: 'httpGet',
      label: 'HTTP GET 探测',
    },
    {
      value: 'grpc',
      label: 'gRPC 健康探测',
    },
  ]
  // 组件列表
  const volumeTypeCommpents = {
    exec,
    tcpSocket,
    httpGet,
    grpc,
  }
</script>

<template>
  <div style="display: flex; justify-content: right; margin-right: 90px">
    <el-select v-model="probeMethod" placeholder="请选择探测方法" style="width: 300px">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
  </div>
  <!-- 探测方法 动态组件 -->
  <component
    :is="volumeTypeCommpents[probeMethod]"
    :probeType="props.probeType"
    :containerItem="props.containerItem"
    style="padding: 20px 90px 0px 0px; height: 270px"
  />
</template>
