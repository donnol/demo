// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueResource from 'vue-resource'

Vue.use(VueResource)
Vue.http.options.root = 'http://127.0.0.1:8080/#' // 后面没有 '/'
Vue.http.headers.common['Authorization'] = 'Basic YXBpOnBhc3N3b3Jk'
// Vue.http.options.emulateJSON = true // 服务器不支持 application/json 时使用
Vue.http.options.emulateHTTP = true // 服务器不支持 PUT PATCH DELETE 等方法

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<app/>',
  components: {
    App: App
  }
})
