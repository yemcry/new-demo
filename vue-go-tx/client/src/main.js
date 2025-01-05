import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css'; // 确保导入 CSS
import router from './router';
import VueRouter from 'vue-router'
import store from './store';
import './assets/global.css'; // 引入全局样式

Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(VueRouter)
new Vue({
  render: h => h(App),
  store,
  router,
}).$mount('#app')
