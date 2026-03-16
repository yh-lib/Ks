<script lang="ts" setup>
import { onBeforeMount, onMounted, reactive, ref } from 'vue'
import { addClusterHandler as addHandler,updateClusterHandler as updateHandler } from '../../api/cluster.js'
import { ElMessage } from 'element-plus'

// add_1 定义一个对象，接收向后端传输的集群信息
const itemForm = reactive({
    clusterId:"",
    clusterAlias:"",
    clusterVersion:"",
    clusterLocation:"",
    clusterStatus:"",
    clusterSize:"",
    clusterKubeconfig:""     
})
const loading = ref(false)
const itemFormRef = ref()
const submitForm = (itemForm) => {
    // add_4 表单校验
    itemFormRef.value.validate((valid)=>{
      if (valid) {
        loading.value = true
        // add_5 调用后端接口,添加集群
        if (props.subMethod == 'create') {
          addHandler(itemForm).then((Response)=>{
            ElMessage({
              message: Response.data.msg,
              type: 'success',
            })
            loading.value = false
          })          
        }else{
          updateHandler(itemForm).then((Response)=>{
            ElMessage({
              message: Response.data.msg,
              type: 'success',
            })
            loading.value = false
            emit('refresh')
          })        
        }
      }else{
          ElMessage({
            message: "请完善表单内容",
            type: 'warning',
          })        
      }
    })
}
// add_3 接受父组件的参数
const subMethod = ref('')
const subRow = reactive({})
const props = defineProps(['subMethod','subRow'])
// add_4 表单校验
const rules = reactive({
    clusterId: [{ required: true, message: '请输入集群ID', trigger: 'blur' },],
    clusterAlias: [{ required: true, message: '请输入集群名称', trigger: 'blur' }],
    clusterLocation: [{ required: true, message: '请输入集群地址', trigger: 'blur' }],
    clusterKubeconfig: [{ required: true, message: '请输入集群kubeConfig', trigger: 'blur' }],
})
// add_6 重置表单
const resetForm = () => {
    itemFormRef.value.resetFields()
}
// add_7 关闭dialog时刷新列表
// 在父组件中通过 @close="closeDialog" 实现




// 解决单向数据流的问题，否则子组件将可以修改父组件的数据：原理是将浅拷贝修改为深度拷贝
onBeforeMount(()=>{
  // step1: 将父组件传递的对象转换成一个字符串
  //    JSON.stringify(props.subRow)
  // step2: 再将字符串转换成JSON
//   data.userForm=JSON.parse(JSON.stringify(props.subRow))
})

// 更新用户后刷新用户列表
const emit = defineEmits(['refresh'])
</script>

<template>
    <!-- 表单 -->
    <el-form
        ref="itemFormRef"
        :model="itemForm"
        label-width="160px"
        center
        class="el-form"
        v-loading="loading"
        :rules="rules"
    >
        <!-- add_2:接受集群信息 -->
        <!-- 左半部分 -->
        <div>
            <el-form-item label="集群ID"  prop="clusterId">
                <el-input v-model="itemForm.clusterId" autocomplete="off" />
            </el-form-item>
            <el-form-item label="集群名称"  prop="clusterAlias">
                <el-input v-model="itemForm.clusterAlias" autocomplete="off" />
            </el-form-item> 
            <el-form-item label="集群位置"  prop="clusterLocation">
                <el-input v-model="itemForm.clusterLocation" autocomplete="off" />
            </el-form-item>                              
        </div>
        <!-- 右半部分 -->
        <div>
            <el-form-item label="kubeconfig"  prop="clusterKubeconfig">
                <el-input v-model="itemForm.clusterKubeconfig" autocomplete="off" type="textarea" :rows="6" style="width: 400px;" />
            </el-form-item>
            <el-form-item style="margin-left: 150px;">
                <!-- add_3: 提交表单 -->
                <el-button type="primary" @click="submitForm(itemForm)" > 
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
.el-form{
    display: flex;
    flex-direction: row;
}
</style>