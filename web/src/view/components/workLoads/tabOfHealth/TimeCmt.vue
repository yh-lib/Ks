<script setup>
  const props = defineProps(['containerItem', 'probeType'])
</script>

<template>
  <el-form :model="props.containerItem" label-width="100px">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-form-item label="延迟间隔">
          <el-input-number
            v-model="props.containerItem[props.probeType].initialDelaySeconds"
            style="width: 300px"
          >
            <template #suffix>
              <span>秒</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="检查间隔">
          <el-input-number
            v-model="props.containerItem[props.probeType].periodSeconds"
            style="width: 300px"
          >
            <template #suffix>
              <span>秒</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="超时时间">
          <el-input-number
            v-model="props.containerItem[props.probeType].timeoutSeconds"
            style="width: 300px"
          >
            <template #suffix>
              <span>秒</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="失败次数">
          <el-input-number
            v-model="props.containerItem[props.probeType].failureThreshold"
            style="width: 300px"
          >
            <template #suffix>
              <span>次</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="成功次数">
          <el-input-number
            v-model="props.containerItem[props.probeType].successThreshold"
            style="width: 300px"
          >
            <template #suffix>
              <span>次</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
      <el-col :span="8">
        <el-form-item label="终止限期">
          <el-input-number
            v-model="props.containerItem[props.probeType].terminationGracePeriodSeconds"
            style="width: 300px"
          >
            <template #suffix>
              <span>秒</span>
            </template>
          </el-input-number>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>
