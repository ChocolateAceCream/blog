import apiAxios from '@/utils/apiAxios'

export const getCurrentUserMenu = () => apiAxios.get('/backend/api/v1/menu/currentUserMenu')
export const getMenuList = () => apiAxios.get('/backend/api/v1/menu/list')
export const postAddMenu = (...args) => apiAxios.post('/backend/api/v1/menu/add', ...args)
export const putEditMenu = (...args) => apiAxios.put('/backend/api/v1/menu/edit', ...args)
export const deleteMenu = (...args) => apiAxios.delete('/backend/api/v1/menu/delete', ...args)
export const getRoleMenuTree = (...args) => apiAxios.post('/backend/api/v1/menu/getRoleMenuTree', ...args)
export const assignRoleMenus = (...args) => apiAxios.post('/backend/api/v1/menu/assignRoleMenus', ...args)
