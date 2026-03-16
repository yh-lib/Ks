<script lang="ts" setup>
import { computed, reactive, ref, onMounted, onBeforeMount  } from 'vue'
import { getClusterListHandler as getListHandler,deleteClusterHandler as deleteHandler } from '../../api/cluster.js'
import { ElMessage, ElMessageBox} from 'element-plus'
import { CircleCheck, CircleClose } from '@element-plus/icons-vue'
// import Add from './Add.vue'

// 独立配置
// Main 标题
const titleName = "集群列表"
// Table 表头
const tableTtile = {
    f1: { prop: "clusterId", label: "ID" },
    f2: { prop: "clusterAlias", label: "别名" },
    f3: { prop: "clusterVersion", label: "版本" },
    f4: { prop: "clusterLocation", label: "位置" },
    f5: { prop: "clusterStatus", label: "状态" },
    f6: { prop: "clusterSize", label: "规模" }
}
// 集群信息接口
interface Cluster {
  clusterId: string
  clusterAlias: string
  clusterVersion: string
  clusterLocation: string
  clusterStatus: string
  clusterSize: number
}



// getList 入口
const getList = () =>{
    loading.value = true
    getListHandler().then((response)=>{
        if (response.status === 200) {
            data.tableData = response.data.data; // 更新 tableData
            loading.value = false
        } else {
            console.error('获取用户列表失败:', response.msg);
        }
    })
}

const search = ref('')

const filterTableData = computed(() =>
  data.tableData.filter(
    (data) =>
      !search.value ||
      data.clusterId.toLowerCase().includes(search.value.toLowerCase())
  )
)

const handleEdit = (index: number, row: User) => {
  console.log(index, row)
}
const handleDelete = (index: number, row: User) => {
  console.log(index, row)
}

const data = reactive({
    tableData:[] as Cluster[],
    userForm:{
        username:"",
        qq:"",
        address:"",
        id:""        
    }    
})

// 加载图标
const loading = ref(false)

// 删除用户
const deleteRow = (row) => {
    console.log("获取到的数据:",filterTableData)
    // 删除提醒
    ElMessageBox.confirm(
        '确认删除集群：' + row.clusterId,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deleteHandler(row.id).then((response)=>{
            ElMessage({
                type: 'success',
                message: response.data.msg,
            })
            // 刷新列表
            getList()
        })
    })
    .catch(() => {
        return
    }) 
}

// 添加用户
const userDialog = ref(false)

const addUser = () => {
    method.value='create'
    data.userForm={}
    userDialog.value = true
}
// 关闭弹窗时是否刷新用户列表
const closeDialog = () =>{
    method.value == 'create' && getList()
}
// 更新用户
const method = ref('')
const updateUser = (row) => {
    // 传递给子组件的操作参数
    method.value='update'
    // 传递当前用户数据给子组件
    data.userForm = row
    // 打开表单弹窗
    userDialog.value = true
}
// 更新用户时刷新用户列表
const updateUserOperation = () =>{
    userDialog.value = false
    getList()
}
// 组件加载时自动调用 getList 方法
onBeforeMount(() => {
    getList();
})
</script>

<template>
    <el-card>
        <!-- Main 标题 -->
        <template #header>
            <div class="card-header">
                <span style="font-size: 24px;">{{ titleName }}</span>
                <el-button type="primary" @click="addUser">添加</el-button>
            </div>
        </template>
        <el-table :data="filterTableData" style="width: 100%"  height="70vh" v-loading="loading">
            <el-table-column :label="tableTtile.f1.label"  :prop="tableTtile.f1.prop" />
            <el-table-column :label="tableTtile.f2.label"   :prop="tableTtile.f2.prop" />
            <el-table-column :label="tableTtile.f3.label" :prop="tableTtile.f3.prop" />
            <el-table-column :label="tableTtile.f4.label" :prop="tableTtile.f4.prop" />
            <el-table-column :label="tableTtile.f5.label" prop="" >
                <template #default="scope">
                    <el-icon v-if="scope.row.clusterStatus == 'true'" color="green"><CircleCheck /></el-icon>  
                    <el-icon v-else color="red"><CircleClose /></el-icon>  
                </template>
            </el-table-column>
            <el-table-column :label="tableTtile.f6.label" :prop="tableTtile.f6.prop" />
            <el-table-column align="right">
            <template #header>
                <el-input v-model="search" size="small" placeholder="Type to search" />
            </template>             
            <template #default="scope">
                <el-button :disabled="scope.row.clusterStatus != 'true'" size="small" @click="updateUser(scope.row)">
                    编辑
                </el-button>
                <el-button
                    size="small"    
                    type="danger"
                    @click="deleteRow(scope.row)"
                >
                    删除
                </el-button>
            </template>
            </el-table-column>
        </el-table>  
        <!-- 添加用户的表单 -->
        <el-dialog 
            v-model="userDialog"
            :title="method == 'create' ? '添加集群' : '更新集群'"
            width="500px"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- 添加用户的表单组件 -->
            <!-- <Add :subMethod='method' :subRow="data.userForm" @refresh="updateUserOperation"/> -->
        </el-dialog>            
    </el-card>
</template>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;

}
</style>