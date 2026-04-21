<script setup>
import { ElMessage } from 'element-plus'
import { ref } from 'vue'

const props = defineProps(['containerType'])

// 定义初始状态
let tabIndex = 1
const activeTabsValue = ref('1')

const editableTabs = ref([
  {
    title: `Container-${tabIndex}`,
    name: tabIndex,
    content: `Container-${tabIndex}`,
  }
])

const handleTabsEdit = (targetName, action) => {
  if (action === 'add') {
    const newTabName = `${++tabIndex}`
    editableTabs.value.push({
      // 新增标签也使用 tabIndex 构建标题，保持风格一致
      title: `Container-${newTabName}`,
      name: newTabName,
      // content 设置为同样的动态字符串
      content: `Container-${newTabName}`,
    })
    activeTabsValue.value = newTabName
  } else if (action === 'remove') {
    // 如果当前只有一个容器，提示并中止后续删除逻辑
    if (editableTabs.value.length === 1 && props.containerType == 'container') {
      ElMessage.error('至少需要保留一个容器')
      return
    }
    const tabs = editableTabs.value
    let activeName = activeTabsValue.value
    if (activeName === targetName) {
      tabs.forEach((tab, index) => {
        if (tab.name === targetName) {
          const nextTab = tabs[index + 1] || tabs[index - 1]
          if (nextTab) {
            activeName = nextTab.name
          }
        }
      })
    }
    
    activeTabsValue.value = activeName
    editableTabs.value = tabs.filter((tab) => tab.name !== targetName)
  }
}

</script>

<template>
    <div>
        <el-button 
            type="primary"
            link
            size="large"
            style="position: absolute;right: 81px;top:8px;font-size: 16px;z-index: 1;"
            @click="handleTabsEdit(nil,'add')"
        >
            添加
        </el-button>
    </div>
    <el-tabs
    v-model="activeTabsValue"
    type="card"
    closable
    @edit="handleTabsEdit"
    style="margin-top: 10px;height: 557px;"
    >
    <el-tab-pane
        v-for="item in editableTabs"
        :key="item.name"
        :label="item.title"
        :name="item.name"
    >
        {{ item.content }}
    </el-tab-pane>
    </el-tabs>
</template>    