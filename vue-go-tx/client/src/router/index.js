import VueRouter from "vue-router";
import store from '@/store'; // 确保路径正确
const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('@/view/login/index.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('@/view/home/index.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/view/register/index.vue')
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    // 返回滚动位置
    return savedPosition || { x: 0, y: 0 };
  }
});

// 处理重复导航错误
const originalPush = router.push;
router.push = function(location, onComplete, onAbort) {
  return originalPush.call(this, location, onComplete, onAbort).catch(err => {
    if (err.name !== 'NavigationDuplicated') {
      throw err; // 其他错误
    }
  });
};

// 路由导航守卫
router.beforeEach((to, from, next) => {
  // 获取当前用户信息
  const currentUser = store.state.sharedData.username

  // 如果用户信息为空,且正在访问非登录页面
  if (!currentUser && to.path !== '/' && to.path !== '/register') {
    // 重定向到登录页面
    next('/')
  } else {
    // 用户存在,或正在访问登录页面,允许导航
    next()
  }
})

export default router;