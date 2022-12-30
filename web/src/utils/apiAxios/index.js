import axios from 'axios'
import supportCancelToken from './cancelToken'
import { addSignature } from './signature'
import { notification } from 'ant-design-vue'
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
    console.log('-------get????', ...args)
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

// request interceptors
apiAxios.interceptors.request.use(config => {
  console.log('-------confoig===', config)
  config.headers['Content-Type'] = 'application/json;charset=UTF-8'
  if (config.meta?.withProgressBar) { NProgress.start() }

  if (!Object.prototype.hasOwnProperty.call(config.params || {}, 'sign')) {
    addSignature(config)
  }

  // params encoding for get request
  if (config.method === 'get' && config.params) {
    const params = encodeURIComponent(JSON.stringify(config.params))
    // const serviceName = config.url.split('/')[1]
    config.url = config.url + '?params=' + params
    console.log('---------config.url-------', config.url)
    config.params = {}
  }

  return config
}, error => {
  return Promise.reject(error)
})

// 响应拦截
apiAxios.interceptors.response.use(res => {
  if (res.config.meta?.withProgressBar) { NProgress.done() }
  // 请求成功
  if (`${res.data.errorCode}` !== '0') {
    if (!Object.prototype.hasOwnProperty.call(res.data, 'errorCode') && !Object.prototype.hasOwnProperty.call(res.data, 'errorMsg')) {
      return Promise.resolve(res)
    }
    notification.error({ message: res.data.errorMsg || res.data })
    return Promise.reject(res.data)
  }
  return Promise.resolve(res)
}, error => {
  // 请求失败
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