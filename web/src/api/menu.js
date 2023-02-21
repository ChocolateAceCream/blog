import apiAxios from '@/utils/apiAxios'

export const getCurrentUserMenu = () => apiAxios.get('/blog/api/v1/menu/currentUserMenu')
export const postAddMenu = (...args) => apiAxios.post('/blog/api/v1/menu/create', ...args)
