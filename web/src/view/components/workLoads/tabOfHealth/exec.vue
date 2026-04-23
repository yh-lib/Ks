<script setup>
  import { ref, watch } from 'vue'

  const props = defineProps(['containerItem', 'probeType'])

  const commandText = ref('')

  const syncTextFromCommand = () => {
    const cmd = props.containerItem?.[props.probeType]?.exec?.command
    commandText.value = Array.isArray(cmd) ? cmd.join('\n') : cmd == null ? '' : String(cmd)
  }

  const syncCommandFromText = () => {
    const item = props.containerItem?.[props.probeType]
    if (!item) return
    if (!item.exec) item.exec = {}

    item.exec.command = String(commandText.value)
      .split(/\r?\n/)
      .map((s) => s.trim())
      .filter(Boolean)
  }

  watch(
    () => props.containerItem?.[props.probeType]?.exec?.command,
    () => {
      syncTextFromCommand()
    },
    { immediate: true, deep: true }
  )
</script>

<template>
  <div>
    <el-form label-width="100px">
      <el-form-item label="探测命令">
        <el-input type="textarea" v-model="commandText" :rows="12" @blur="syncCommandFromText" />
      </el-form-item>
    </el-form>
  </div>
</template>
