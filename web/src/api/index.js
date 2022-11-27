import axios from 'axios'
import apiAxios from '@/utils/apiAxios'

export default {
  install(app) {
    Object.defineProperty(app.config.globalProperties, '$axios', {
      get: () => axios,
    })
    Object.defineProperty(app.config.globalProperties, '$apiAxios', {
      get: () => apiAxios,
    })
  },
}
