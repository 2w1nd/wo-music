import { get, post } from './request'


const HttpManager = {
  // ==================> 管理员API
  // 创建
  create: (params) => post(`admin/admin`, params),
}

export { HttpManager }
