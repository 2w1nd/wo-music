import axios from 'axios'
import { BASE_URL } from './config'
import router from '../router'

axios.defaults.timeout = 5000 // 超时时间设置
axios.defaults.withCredentials = true // true允许跨域
// Content-Type 响应头
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8'

axios.defaults.baseURL = BASE_URL

/**
   * 封装get方法
   * @param url
   * @param data
   * @returns {Promise}
   */
 export function get (url, params = {}, responseType = 'json') {
  return new Promise((resolve, reject) => {
    axios.get(url, {
      params: params,
      responseType
    })
      .then(response => {
        resolve(response.data)
      })
      .catch(err => {
        reject(err)
      })
  })
}

/**
   * 封装post请求
   * @param url
   * @param data
   * @returns {Promise}
   */
export function post (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.post(url, data)
      .then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
  })
}

/**
   * 封装delete请求
   * @param url
   * @param data
   * @returns {Promise}
   */
export function deletes (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.delete(url, data)
      .then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
  })
}

/**
   * 封装put请求
   * @param url
   * @param data
   * @returns {Promise}
   */
export function put (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.put(url, data)
      .then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
  })
}
