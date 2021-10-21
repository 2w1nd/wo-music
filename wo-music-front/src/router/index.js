import {createRouter, createWebHashHistory} from "vue-router"; 
import Home from '../view/Home.vue';


const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    children: [
      
    ]
  }, {
    path: "/login",
    name: "Login",
    component: () => import ( /* webpackChunkName: "login" */ "../view/Login.vue")
  }

];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});


export default router;
