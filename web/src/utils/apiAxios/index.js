import axios from 'axios'
import supportCancelToken from './cancelToken'
import { addSignature } from './signature'
import { ElMessage } from 'element-plus'
import { logout } from '@/shared/hooks/index'
import NProgress from 'nprogress'
import qs from 'qs'

const apiAxios = new Proxy(axios.create({
  // https://cn.vitejs.dev/guide/env-and-mode.html
  baseURL: import.meta.env.VITE_APP_API_BASE_URL || '/',
  timeout: 1000 * 60,
  paramsSerializer: {
    serialize: function(params) {
      return qs.stringify(params, { indices: false })
    }
  }
}), {
  get(target, ...args) {
    return Reflect.get(target, ...args) || Reflect.get(axios, ...args)
  }
})

apiAxios.defaults.meta = {
  retry: 0/* times*/, retryDelay: 100/* ms*/, curRetry: 0/* times*/,
  // 断开相同请求，判断条件 如果!!cancelToken存在 则计算config.url+cancelToken的值作为唯一key值，key值相同，则断开之前请求
  cancelToken: '',
  withProgressBar: false,
}

apiAxios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8'
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'

supportCancelToken(apiAxios)

let activeRequest = 0
let logoutFlag = false
// request interceptors
apiAxios.interceptors.request.use(config => {
  activeRequest++
  if (config['Content-Type']) {
    config.headers['Content-Type'] = config['Content-Type']
  } else {
    config.headers['Content-Type'] = 'application/json;charset=UTF-8'
  }
  if (config.meta?.withProgressBar) { NProgress.start() }
  if (!Object.prototype.hasOwnProperty.call(config.params || {}, 'sign')) {
    addSignature(config)
  }

  // params encoding for get request
  if (config.method === 'get' && config.params) {
    const params = encodeURIComponent(JSON.stringify(config.params))
    // const serviceName = config.url.split('/')[1]
    config.url = config.url + '?params=' + params
    // console.log('---------config.url-------', config.url)
    config.params = {}
  }

  return config
}, error => {
  activeRequest--
  return Promise.reject(error)
})

// 响应拦截
apiAxios.interceptors.response.use(res => {
  if (res.config.meta?.withProgressBar) { NProgress.done() }
  // 请求成功
  activeRequest--
  if (res.data.errorCode !== 0) {
    ElMessage({
      message: res.data.msg,
      type: 'error',
      duration: 5 * 1000
    })
    // unauthorized, logout current user
    if (res.data.errorCode === 401) {
      if (!logoutFlag) {
        logout()
        logoutFlag = true
      }
      if (activeRequest === 0) {
        // last request
        logoutFlag = false
      }
    }
    return Promise.reject(res.data)
  }
  return Promise.resolve(res)
}, error => {
  // 请求失败
  activeRequest--
  if (axios.isCancel(error)) {
    console.error('cancel by client')
  } else {
    const config = error.config
    if (config?.meta && config.meta.curRetry !== config.meta.retry) {
      config.meta.curRetry++
      return new Promise(resolve => {
        setTimeout(() => {
          console.warn(`${config.url},retry: ${config.meta.curRetry} times`)
          resolve(apiAxios(config))
        }, config.meta.retryDelay, 1000)
      })
    }
  }
  return Promise.reject(error)
})
export default apiAxios
