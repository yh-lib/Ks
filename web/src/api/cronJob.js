import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取 CronJob 列表
export const getCronJobListHandler = (clusterId, nameSpace) => {
  return request(API_CONFIG.cronJobListApi, { clusterId, nameSpace }, 'get')
}

// 获取 CronJob 详情
export const getCronJobHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.cronJobGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 CronJob
export const deleteCronJobHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.cronJobDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 创建 CronJob
export const createCronJobHandler = (itemForm) => {
  return request(API_CONFIG.cronJobCreateApi, itemForm, 'post')
}

// 更新 CronJob
export const updateCronJobHandler = (itemFrom) => {
  return request(API_CONFIG.cronJobUpdateApi, itemFrom, 'post')
}
