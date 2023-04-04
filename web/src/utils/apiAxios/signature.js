// 对象的属性进行排序
import md5 from 'js-md5'

export function addSignature(config = {}) {
  const SIGNATURE_SECRET = import.meta.env.VITE_APP_API_SIGNATURE_SECRET
  const APP_ID = import.meta.env.VITE_APP_APP_ID
  const timestamp = parseInt((new Date()).getTime() / 1000)
  const nonce = Math.floor(Math.random() * 8999) + 1000
  const strArr = [APP_ID, nonce, timestamp, SIGNATURE_SECRET]
  config.headers.nonce = nonce
  config.headers.appId = APP_ID
  config.headers.timestamp = timestamp
  config.headers.sign = md5(strArr.join('')).toUpperCase()
  // console.log('------origin----', strArr.join(''))
}
