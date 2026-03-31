import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取集群列表
export const getnodeListHandler = (clusterId) => {
    return request(API_CONFIG.nodeListApi, { clusterId: clusterId }, 'get', 10000)
}

// 获取集群详情
export const getnodeHandler = (clusterId, name) => {
    return request(API_CONFIG.nodeGetApi, { clusterId, name }, 'get', 10000)
}

// 删除集群
export const deletenodeHandler = (nodeId) => {
    return request(API_CONFIG.nodeDeleteApi, { nodeId }, 'get', 10000)
}

// 添加集群
export const addnodeHandler = (itemFrom) => {
    return request(API_CONFIG.nodeAddApi, itemFrom, 'post', 10000)
}


// 更新集群
export const updatenodeHandler = (nodeInfo) => {
    return request(API_CONFIG.nodeUpdateApi, nodeInfo, 'post', 10000)
}