import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist'
import App from './App.vue'
import router from './router'
import i18n from './locales'

import moment from 'moment'

import './assets/styles/element-variables.scss'
import './assets/styles/application.scss'


const app = createApp(App)

// piania persistent
const store = createPinia()
store.use(piniaPluginPersist)
app.use(store)
app.use(i18n)


app.config.globalProperties.$moment = moment
app.use(router)

app.mount('#app')

