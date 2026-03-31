<script lang="ts" setup>
import { computed, reactive, ref, onBeforeMount  } from 'vue'
import { getnodeListHandler as getListHandler,getnodeHandler as getHandler } from '../../api/node.js'
import { getClusterListHandler} from '../../api/cluster.js'
import { ElSelect } from 'element-plus'

// 需要的数据变量
const data = reactive({
    loading:false,  // 默认关闭加载图标
    items: [],  // 后端返回的元数据
    search: "", // 接收 header 搜索框中数据
    // 集群选择器
    clusterOptions:[],  // 集群选择 options
    curClusterId: "",   // 当前选择的集群id
    defaultClusterId: "in-cluster", // 默认选择的集群id
})



// 子组件加载前自动获取数据
onBeforeMount(async () => {
    await getclusterOptions()   // 获取集群列表
    data.curClusterId=data.defaultClusterId // 获取基础设施集群的Node列表
    getList();
})

// 关闭dialog时刷新用户列表
const closeDialog = () =>{
    getList()
}

// 更新node配置后，刷新node列表
const refreshList = () =>{
    opDialog.value = false
    getList()
}

// 内存地址单位转换
const memoryKi2MB = (memory) => {
    if (memory === null || memory === undefined) return ''
    const match = String(memory).trim().match(/^(\d+(?:\.\d+)?)Ki$/i)
    return match ? Math.floor(Number(match[1]) / 1000) : ''
}

// 搜索后的 node 列表数据
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !data.search ||
      item.metadata.name.toLowerCase().includes(data.search.toLowerCase()) ||
      item.status.addresses[0].address.toLowerCase().includes(data.search.toLowerCase())
  )
)

// 获取当前集群Node列表
const getList = () =>{
    data.loading = true
    getListHandler(data.curClusterId).then((res)=>{
        data.items = res.data.data.items;
        data.loading = false
    })
}

// 获取集群列表
const getclusterOptions = async ()=>{
    await getClusterListHandler().then((res)=>{
        if (res.data.status === 200) {
            data.clusterOptions = res.data.data.items;
            console.log('获取集群列表:::', data.clusterOptions);
        } else {
            console.error('获取集群列表失败:', res.data.message);
        }
    })
}
</script>

<template>
    <!-- card -->
    <el-card>
        <!-- header -->
        <template #header>
            <div class="card-header">
                <div>
                    <span style="font-size: 24px;">节点列表</span>
                </div>
                <div>
                    <el-select v-model="data.curClusterId" placeholder="选择集群" style="width: 240px;margin-right: 10px;" @change="getList(data.curClusterId)">
                        <el-option
                            v-for="item in data.clusterOptions"
                            :key="item.clusterId"
                            :label="item.clusterId"
                            :value="item.clusterId"
                            :disabled="item.disabled"
                        />     
                    </el-select>               
                    <el-input v-model="data.search" style="width: 240px" placeholder="搜索节点"/>
                </div>
            </div>
        </template>
        <!-- middle -->
        <el-table :data="filterTableData" style="width: 100%;"  height="70vh" v-loading="data.loading">
            <el-table-column label="主机名">
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>
            </el-table-column>
            <el-table-column label="IP地址" prop="ipAddress" >
                <template #default="scope">
                    {{ scope.row.status.addresses[0].address}} 
                </template>
            </el-table-column>
            <el-table-column label="规格" prop="capacity" >
                <template #default="scope">
                    {{ scope.row.status.capacity.cpu}}C
                    |
                    {{ memoryKi2MB(scope.row.status.capacity.memory) }}MB
                </template>
            </el-table-column>
            <el-table-column label="角色" prop="role" >
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">
                        {{ scope.row.status.addresses[1].address}} 
                    </el-button>                        
                </template>                  
            </el-table-column>
            <el-table-column label="状态" prop="status" >
                <template #default="scope">
                    <span v-if="scope.row.status.conditions[scope.row.status.conditions.length-1].status == 'True'" style="color: green;">Ready</span>
                    <span v-else  style="color: red;">NotReady</span>
                </template>
            </el-table-column>
            <el-table-column label="禁止调度" prop="scheduling" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>
            </el-table-column>
            <el-table-column label="驱逐" prop="eviction" >
                <template #default="scope">
                    <el-switch v-model="value1" />
                </template>     
            </el-table-column>
            <el-table-column label="操作" prop="edit">
                <template #default="scope">
                    <el-button link type="primary" @click="edit(scope.row)">配置</el-button>
                </template>
            </el-table-column>
        </el-table>  
    </el-card>
</template>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>





        <!-- <el-dialog 
            v-model="opDialog"
            :title="method == 'create' ? '添加节点' : '更新节点'"
            width="50vw"
            @close="closeDialog"
            destroy-on-close
        > -->
            <!-- add_6:添加节点表单 -->
            <!-- update3 将method.value、data.itemForm赋值给子组件 -->
            <!-- <Add :subMethod='method' :subRow="data.itemForm" @refresh="refreshList"/>
        </el-dialog>             -->

<!-- // +++++++++++++++++


// 样式
// const props = {
//   value: 'id',
//   label: 'label',
//   options: 'options',
//   disabled: 'disabled',
// }

//   editItem: {},
//   editNodeName: "",
//   detailItem: {},
//   detailNodeName: "",
//   createName: "",





// update_2 更新
// const updateItem = (row) => {
//     // 获取节点详情
//     getHandler(row.nodeId).then((response)=>{
//         data.itemForm=response.data.data.item  
//         // 注意下面这两步：如果放在axios外面，则赋值可能失败，因为是异步运行
//         // 传递给操作参数给子组件
//         method.value='update'
//         // 打开表单弹窗
//         opDialog.value = true
//     }) 
// } -->