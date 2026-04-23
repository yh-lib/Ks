<script setup>
  import { ref } from 'vue'
  import ProbeMethod from './ProbeMethod.vue'
  import TimeCmt from './TimeCmt.vue'

  const props = defineProps(['containerItem'])
  const activeCollapse = ref(['1'])
  const activeTab = ref('startupProbe')
</script>

<template>
  <el-tabs v-model="activeTab" type="card">
    <el-tab-pane label="启动探针 (startupProbe)" name="startupProbe">
      <el-collapse v-model="activeCollapse" accordion>
        <el-collapse-item title="探测参数" name="1">
          <time-cmt
            :container-item="props.containerItem"
            probe-type="startupProbe"
            style="padding: 10px 50px"
          />
        </el-collapse-item>
        <el-collapse-item title="探测方式" name="2">
          <probe-method :containerItem="props.containerItem" probe-type="startupProbe" />
        </el-collapse-item>
      </el-collapse>
    </el-tab-pane>
    <el-tab-pane label="存活探针 (livenessProbe)" name="livenessProbe">
      <el-collapse v-model="activeCollapse" accordion>
        <el-collapse-item title="通用参数" name="1">
          <time-cmt
            :container-item="props.containerItem"
            probe-type="livenessProbe"
            style="padding: 10px 50px"
          />
        </el-collapse-item>
        <el-collapse-item title="探测方式" name="2">
          <probe-method :containerItem="props.containerItem" probe-type="livenessProbe" />
        </el-collapse-item>
      </el-collapse>
    </el-tab-pane>
    <el-tab-pane label="就绪探针 (readinessProbe)" name="readinessProbe">
      <el-collapse v-model="activeCollapse" accordion>
        <el-collapse-item title="通用参数" name="1">
          <time-cmt
            :container-item="props.containerItem"
            probe-type="readinessProbe"
            style="padding: 10px 50px"
          />
        </el-collapse-item>
        <el-collapse-item title="探测方式" name="2">
          <probe-method :containerItem="props.containerItem" probe-type="readinessProbe" />
        </el-collapse-item>
      </el-collapse>
    </el-tab-pane>
  </el-tabs>
</template>
