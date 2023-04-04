import apiAxios from '@/utils/apiAxios'

export const getCurrentUserMenu = () => apiAxios.get('/blog/api/v1/menu/currentUserMenu')
export const getMenuList = () => apiAxios.get('/blog/api/v1/menu/list')
export const postAddMenu = (...args) => apiAxios.post('/blog/api/v1/menu/add', ...args)
export const putEditMenu = (...args) => apiAxios.put('/blog/api/v1/menu/edit', ...args)
export const deleteMenu = (...args) => apiAxios.delete('/blog/api/v1/menu/delete', ...args)
