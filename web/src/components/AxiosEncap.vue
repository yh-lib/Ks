<script setup>
  import request from '../api/axiosEncap.js'

  const encapGet = () => {
    request('/getUserList', { id: 1001 }, 'get', 3000)
  }

  const encapPost = () => {
    request('/postOrderList', { Authorization: 'yyyyyyyyyyyyyyy' }, 'post', 3000).then(
      (response) => {
        alert(response.data.msg)
      }
    )
  }
</script>

<template>
  <button @click="encapGet">封装测试Get</button>
  <button @click="encapPost">封装测试Post</button>
</template>
