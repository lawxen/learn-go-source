// 导入Vue
import { createApp } from 'vue'
// 导入Vue扩展插件
import axios from 'axios'
import VueAxios from 'vue-axios'
import { createRouter, createWebHistory } from 'vue-router'
// 导入组件
import App from './App.vue'
import Product from './components/Product.vue'
import Signin from './components/Signin.vue'

// 定义路由
const routes = [
  { path: '/', component: Signin },
  { path: '/product', component: Product },
]
// 创建路由对象
const router = createRouter({
  // 设置历史记录模式
  history: createWebHistory(),
  // routes: routes的缩写
  routes,
})
// 创建Vue对象
const app = createApp(App)
// 将路由对象绑定到Vue对象
app.use(router)
// 将vue-axios与axios关联并绑定到Vue对象
app.use(VueAxios,axios)
// 挂载使用Vue对象
app.mount('#app')
