import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist'
import App from './App.vue'
import router from './router'
import sharedComponents from '@/shared'
import moment from 'moment'
import { setI18n } from './locales'
// import dayjs from 'dayjs'
// import relativeTime from 'dayjs/plugin/relativeTime'

import './assets/styles/element-variables.scss'
import './assets/styles/application.scss'
import './assets/iconfont/iconfont.css'



const app = createApp(App)
// dayjs.extend(relativeTime)
app.provide('dayjs', dayjs)

// piania persistent
const store = createPinia()
store.use(piniaPluginPersist)
app.use(store)
const i18nConfig = setI18n()
app.use(i18nConfig)

app.use(sharedComponents)


app.config.globalProperties.$moment = moment
app.use(router)

app.mount('#app')

