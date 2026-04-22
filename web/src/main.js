// 基础导入
import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

const app = createApp(App)

// 导入 router
import router from './router/index.js'
app.use(router)

// 导入 pinia
import { createPinia } from 'pinia'
const pinia = createPinia()
app.use(pinia)

// 导入 element-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
app.use(ElementPlus)

// 挂在到id为app的html元素
app.mount('#app')
