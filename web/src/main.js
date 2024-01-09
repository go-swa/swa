import { createApp } from 'vue'
import router from '@/router/index'
import run from '@/core/swa.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import Nprogress from 'nprogress'
import '@/permission'
import './core/swa'
import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import './styles/element_visiable.scss'
import 'nprogress/nprogress.css'
import 'default-passive-events'
import App from './App.vue'

Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()

const app = createApp(App)
app.config.productionTip = false

app
  .use(run)
  .use(store)
  .use(auth)
  .use(router)
  .mount('#app')
export default app
