<script setup>
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)
</script>

<template>
    <el-table :data="workLoadItem.item.spec.template.spec.volumes" style="width: 100%; height:100%">
        <!-- 名称 -->
        <el-table-column prop="name" label="名称">
            <template #default="scope">
                <el-input v-model="scope.row.name" placeholder="请输入存储卷的名称"></el-input>
            </template>                        
        </el-table-column>
        <!-- 类型 -->
        <el-table-column prop="type" label="类型">
            <!-- <template #default="scope">
                <el-input v-model="scope.row.value" placeholder="请输入标签的 Value"></el-input>
            </template>  -->
        </el-table-column>
         <!-- Volume 配置 -->
        <el-table-column prop="config" label="Volume 配置">
            <!-- <template #default="scope">
                <el-input v-model="scope.row.value" placeholder="请输入标签的 Value"></el-input>
            </template>  -->
        </el-table-column>
        <!-- Operation -->
        <el-table-column width="200px" align="center">
            <template #header>
                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('addTableRow')">添加</el-button>
            </template>
            <template #default="scope">
                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('deleteTableRow',scope.$index)">删除</el-button>
            </template>                      
        </el-table-column>
    </el-table> 
</template>    