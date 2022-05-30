
import VueRouter from 'vue-router'
import Vue from 'vue'
Vue.use(VueRouter)
import HelloWorld from "@/components/HelloWorld";
import vmProvider from "@/components/vmProvider";

export default new VueRouter({
    mode: 'history',
    routes: [
        {
            path: "/home",
            name: "home",
            component: vmProvider,
        },
        {
            path: "/",
            component: HelloWorld,
        },
    ]
})

