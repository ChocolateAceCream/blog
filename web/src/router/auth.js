export const auth = {
  path: '/auth', name: 'auth', redirect: '/auth/login', component: () => import('@/views/auth/index'),
  children: [
    { path: 'login', name: 'login', component: () => import('@/views/auth/login/') },
    { path: 'resetPassword', name: 'resetPassword', component: () => import('@/views/auth/ResetPassword') },
    { path: 'register', name: 'register', component: () => import('@/views/auth/Register') },
  ]
}
