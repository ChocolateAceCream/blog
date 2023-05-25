import apiAxios from '@/utils/apiAxios'

export const getCasbinByRoleId = (...args) => apiAxios.get('/backend/api/v1/casbin/list', ...args)
export const postUpdateCasbin = (...args) => apiAxios.post('/backend/api/v1/casbin/update', ...args)
