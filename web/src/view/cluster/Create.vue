<script lang="ts" setup>
  import { onBeforeMount, reactive, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { createHandler, updateHandler } from '../../api/generic'

  // ++++++增
  // add_1 定义一个对象，接收向后端传输的集群信息
  const data = reactive({
    itemForm: {
      clusterId: '',
      clusterAlias: '',
      clusterVersion: '',
      clusterLocation: '',
      clusterStatus: '',
      clusterSize: '',
      clusterNodeNum: '',
      clusterPodNum: '',
      kubeconfig: '',
    },
  })

  const itemFormRef = ref()
  const submitForm = (itemForm) => {
    // add_4 表单校验
    itemFormRef.value.validate((valid) => {
      if (valid) {
        // add_5 调用后端接口,添加集群
        if (props.subMethod == 'create') {
          createHandler('cluster', itemForm)
            .then((response) => {
              if (response.data.status === 200) {
                ElMessage({
                  message: response.data.message,
                  type: 'success',
                })
                emit('refresh')
              }
            })
            .catch(() => {
              // 处理网络错误或其他异常
              ElMessage.error('添加集群失败,请核对集群配置后重试')
            })
        } else {
          // update_6 调用后端接口，更新集群
          updateHandler('cluster', itemForm)
            .then((response) => {
              ElMessage({
                message: response.data.message,
                type: 'success',
              })
              // update_7 刷新列表
              emit('refresh')
            })
            .catch(() => {
              // 处理网络错误或其他异常
              ElMessage.error('更新集群失败,请核对集群配置后重试')
            })
        }
      } else {
        ElMessage({
          message: '请完善表单内容',
          type: 'warning',
        })
      }
    })
  }

  const props = defineProps(['subMethod', 'subRow'])

  const rules = reactive({
    clusterId: [{ required: true, message: '请输入集群ID', trigger: 'blur' }],
    clusterAlias: [{ required: true, message: '请输入集群名称', trigger: 'blur' }],
    clusterLocation: [{ required: true, message: '请输入集群地址', trigger: 'blur' }],
    clusterKubeconfig: [{ required: true, message: '请输入集群kubeConfig', trigger: 'blur' }],
  })

  const resetForm = () => {
    itemFormRef.value.resetFields()
  }

  onBeforeMount(() => {
    data.itemForm = JSON.parse(JSON.stringify(props.subRow))
  })

  const emit = defineEmits(['refresh'])
</script>

<template>
  <!-- 表单 -->
  <el-form
    ref="itemFormRef"
    :model="data.itemForm"
    label-width="160px"
    center
    class="el-form"
    :rules="rules"
  >
    <!-- add_2:接受集群信息 -->
    <!-- 左半部分 -->
    <div>
      <el-form-item label="集群ID" prop="clusterId">
        <el-input
          v-model="data.itemForm.clusterId"
          autocomplete="off"
          :readonly="props.subMethod == 'update'"
        />
      </el-form-item>
      <el-form-item label="集群名称" prop="clusterAlias">
        <el-input v-model="data.itemForm.clusterAlias" autocomplete="off" />
      </el-form-item>
      <el-form-item label="集群位置" prop="clusterLocation">
        <el-input v-model="data.itemForm.clusterLocation" autocomplete="off" />
      </el-form-item>
    </div>
    <!-- 右半部分 -->
    <div>
      <el-form-item label="kubeconfig" prop="kubeconfig">
        <el-input
          v-model="data.itemForm.kubeconfig"
          autocomplete="off"
          type="textarea"
          :rows="6"
          style="width: 400px"
        />
      </el-form-item>
      <el-form-item style="margin-left: 150px">
        <!-- add_3 提交表单 -->
        <!-- update_5 提交表单 -->
        <el-button type="primary" @click="submitForm(data.itemForm)">
          {{ props.subMethod == 'create' ? '提交' : '更新' }}
        </el-button>
        <el-button @click="resetForm()">
          {{ props.subMethod == 'create' ? '清空' : '重置' }}
        </el-button>
      </el-form-item>
    </div>
  </el-form>
</template>

<style lang="less" scoped>
  .el-form {
    display: flex;
    flex-direction: row;
  }
</style>
