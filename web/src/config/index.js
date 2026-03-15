// 放置项目的配置信息

// 后端接口配置
const BASE_URL = import.meta.env.VITE_BASE_URL
export const API_CONFIG = {
    loginApi: `${BASE_URL}/auth/login`,
    logoutApi: `${BASE_URL}/auth/logout`,
    userListApi: `${BASE_URL}/user/list`,
    userDeleteApi: `${BASE_URL}/user/delete`,
    userAddApi: `${BASE_URL}/user/add`,
    userUpdateApi: `${BASE_URL}/user/update`
}

export const CONFIG = {
    TOKEN_NAME: "Authorization"
}