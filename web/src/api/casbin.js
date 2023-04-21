import apiAxios from '@/utils/apiAxios'

export const getCasbinByRoleId = (...args) => apiAxios.get('/blog/api/v1/casbin/list', ...args)
export const postUpdateCasbin = (...args) => apiAxios.post('/blog/api/v1/casbin/update', ...args)
