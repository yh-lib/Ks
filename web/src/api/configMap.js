import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取 configMap 列表
export const getConfigMapListHandler = (clusterId, nameSpace) => {
  return request(API_CONFIG.configMapListApi, { clusterId, nameSpace }, 'get')
}

// 获取 configMap 详情
export const getConfigMapHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.configMapGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 configMap
export const deleteConfigMapHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.configMapDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 添加 configMap
export const createConfigMapHandler = (clusterId, item) => {
  return request(API_CONFIG.configMapCreateApi, { clusterId, item }, 'post')
}

// 更新 configMap
export const updateConfigMapHandler = (ConfigMapInfo) => {
  return request(API_CONFIG.configMapUpdateApi, ConfigMapInfo, 'post')
}
