import { createApp } from 'vue'
import store from './store'
import App from './App.vue'
import router from './router'
import installElementPlus from './plugins/element'

const app = createApp(App)
installElementPlus(app)
app
  .use(router)
  .use(store)
  .mount('#app')

