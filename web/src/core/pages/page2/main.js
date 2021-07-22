import Vue from 'vue'

import App from './App';
import '../../plugins/element.js'
import '../../plugins/axios-plug.js'

Vue.config.productionTip = false


new Vue({
    render: h => h(App),
}).$mount('#app')
