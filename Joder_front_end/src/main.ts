import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { loadFonts } from './plugins/webfontloader'
// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import '@/assets/global.scss'
loadFonts()
const vuetify = createVuetify({
  components,
  directives,
})
createApp(App)
  .use(router)
  .use(vuetify)
  .use(createPinia())
  .mount('#app')
