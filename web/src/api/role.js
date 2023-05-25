import apiAxios from '@/utils/apiAxios'

export const getRoleList = () => apiAxios.get('/backend/api/v1/role/list')
export const postAddRole = (...args) => apiAxios.post('/backend/api/v1/role/add', ...args)
export const putEditRole = (...args) => apiAxios.put('/backend/api/v1/role/edit', ...args)
export const deleteRole = (...args) => apiAxios.delete('/backend/api/v1/role/delete', ...args)
