<script setup>
import { ref,reactive, watch } from 'vue';
// 导入图标
import { User, Lock } from '@element-plus/icons-vue';
import { login } from '../api/login';
const loginInfo = reactive({
    username: '',
    password: '',
})
// 表单校验
const rules = reactive({
    username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, message: '请输入6位以上的密码', trigger: 'blur' },
    ],
})
// 声明 loginRef，用于存放表单数据
const loginRef = ref()
// 表单校验失败时，禁用登录按钮
let loginButtonDisabled = ref(true)
watch(
    [
        () => loginInfo.username,
        () => loginInfo.password
    ],
    () => {
        loginRef.value.validate(
            (valid) => {
                valid?loginButtonDisabled.value = false : loginButtonDisabled.value = true
            }
        )
    }
)
// 登录按钮调用后端接口
// const submitForm = () => login(loginInfo.username,loginInfo.password)
</script>

<template>  
<div class="login">
  <el-card class="login__el-card">
    <!-- 标题 -->
    <h2>kubernetes 管理平台</h2>
    <!-- 表单 -->
    <el-form
        ref="loginRef"
        :model="loginInfo"
        :rules="rules"
        label-width="auto"
        class="demo-ruleForm"
    >
        <!-- 用户名: 输入表单 -->
        <el-form-item prop="username">
            <el-input 
                v-model="loginInfo.username"
                placeholder="输入用户名" 
                :prefix-icon="User" 
                clearable
            />
        </el-form-item>
        <!-- 密码：输入表单 -->
        <el-form-item prop="password">
            <el-input 
                v-model="loginInfo.password"
                type="password"
                placeholder="输入密码"
                :prefix-icon="Lock"
                show-password
                
            />
        </el-form-item>
        <!-- 登录：按钮 -->
            <el-button type="primary" @click="login(loginInfo.username,loginInfo.password)" :disabled="loginButtonDisabled"> 登录 </el-button>
    </el-form>    
  </el-card>
</div>
</template>

<style scoped>
.login{
    width: 100vw;
    height: 100vh; 
    display: flex; 
    justify-content: center; 
    /* align-items: center; */
    .login__el-card{
        width: 550px;
        height: 310px;
        margin-top: 10%;
    }
    .demo-ruleForm{
        width: 80%;
        margin: 7% auto auto auto;
    }
}
</style>