import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
// 引入组件库的少量全局样式变量
import 'tdesign-vue-next/es/style/index.css';
import TDesign from 'tdesign-vue-next';


const pinia = createPinia()
createApp(App).use(TDesign).use(pinia).mount('#app')
