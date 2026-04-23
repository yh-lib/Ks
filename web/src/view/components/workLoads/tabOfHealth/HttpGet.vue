<script setup>
  import TableOfKeyValue from '../TableOfKeyValue.vue'

  const props = defineProps(['containerItem', 'probeType'])
  const addTableRow = () => {
    props.containerItem[props.probeType].httpGet.httpHeaders.push({ name: '', value: '' })
  }
  const deleteTableRow = (index) => {
    props.containerItem[props.probeType].httpGet.httpHeaders.splice(index, 1)
  }
</script>

<template>
  <div style="margin-left: 50px">
    <el-form label-width="100px">
      <el-row :gutter="200">
        <el-col :span="12">
          <el-form-item label="探测主机">
            <el-input v-model="props.containerItem[props.probeType].httpGet.host" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="探测端口">
            <el-input v-model="props.containerItem[props.probeType].httpGet.port" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="探测路径">
            <el-input v-model="props.containerItem[props.probeType].httpGet.path" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="探测协议">
            <el-select v-model="props.containerItem[props.probeType].httpGet.scheme">
              <el-option
                v-for="item in ['http', 'https']"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="自定义请求头">
            <el-table
              :data="props.containerItem[props.probeType].httpGet.httpHeaders"
              style="width: 100%; height: 190px"
              border
            >
              <!-- Name -->
              <el-table-column prop="name" label="Name">
                <template #default="scope">
                  <el-input v-model="scope.row.name" placeholder="请输入请求头的 Name"></el-input>
                </template>
              </el-table-column>
              <!-- Value -->
              <el-table-column prop="value" label="Value">
                <template #default="scope">
                  <el-input v-model="scope.row.value" placeholder="请输入请求头的 Value"></el-input>
                </template>
              </el-table-column>
              <!-- Operation -->
              <el-table-column width="200px" align="center">
                <template #header>
                  <el-button
                    type="primary"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="addTableRow"
                    >添加</el-button
                  >
                </template>
                <template #default="scope">
                  <el-button
                    type="danger"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="deleteTableRow(scope.$index)"
                    >删除</el-button
                  >
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>
