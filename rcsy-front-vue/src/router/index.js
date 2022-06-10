import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index.vue'
import {getStore} from '@/controller/utils'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/Login.vue'),
    meta: {
      title: '基于智能合约的关键密文共享系统'
    }
  },
  {
    path:"/",
            name:'Main',
            component: () => import('../views/Main.vue'),
            redirect:"/home",
            //所有页面都为main组件的子页面
            children:[
                {
                    path:"/home",
                    name:'home',
                    component: () => import('../views/home/index.vue'),
					meta: {
					  title: '基于智能合约的关键密文共享系统'
					}
                },
                {
                    path:'/page1',
                    name:'page1',
                    component: () => import('../views/other/pageOne.vue'),
					meta: {
					  title: '基于智能合约的关键密文共享系统'
					}
                },
                {
                    path:'/page2',
                    name:'page2',
                    component: () => import('../views/other/pageTwo.vue'),
					meta: {
					  title: '基于智能合约的关键密文共享系统'
					}
                }
              ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

router.beforeEach((to, from, next) => {
  if(to.meta.title){
    document.title = to.meta.title
  }
  if (to.meta.requireAuth) { // 判断该路由是否需要登录权限
    if (getStore('username')) { // 判断本地是否存在token
      next()
    } else {
      // 未登录,跳转到登陆页面
      next({
        path: '/login'
      })
    }
  }else{
    next()
  }
});