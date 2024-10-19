import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "@/views/LoginPage.vue";
import RegisterPage from "@/views/RegisterPage.vue";
import DashboardPage from "@/views/DashboardPage.vue";

const routes = [
    {
        path: '/',
        redirect: '/login' 
    },

    {
        path: '/login',
        name: 'Login',
        component: LoginPage
    },

    {
        path: '/register',
        name: 'Register',
        component: RegisterPage
    },

    {
        path: '/dashboard',
        name: 'Dashboard',
        component: DashboardPage
    },
]


const router = createRouter({
    history: createWebHistory(),
    routes
})


export default router;