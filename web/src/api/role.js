import apiAxios from '@/utils/apiAxios'

export const getRoleList = () => apiAxios.get('/blog/api/v1/role/list')
export const postAddRole = (...args) => apiAxios.post('/blog/api/v1/role/add', ...args)
export const putEditRole = (...args) => apiAxios.put('/blog/api/v1/role/edit', ...args)
export const deleteRole = (...args) => apiAxios.delete('/blog/api/v1/role/delete', ...args)
