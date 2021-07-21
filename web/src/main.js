import Vue from 'vue'

import App from './App';
import {MyApi} from "./api/MyApi";
import './plugins/element.js'

Vue.config.productionTip = false


//全局注册api接口
Vue.prototype.$api = MyApi
new Vue({
    render: h => h(App),
}).$mount('#app')
