import App from './App.vue'
import router from './router/router'
import Vue from 'vue'
new Vue({
    render: h => h(App),
    router,
}).$mount('#app')