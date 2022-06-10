import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/less/index.less'

Vue.config.productionTip = false

import '../public/js/rem.js'
import ElementUi from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
// 引入echarts
import * as echarts from 'echarts'
Vue.prototype.$echarts = echarts

// 导入富文本编辑器
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
Vue.use(VueQuillEditor)
Vue.use(ElementUi)



new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
