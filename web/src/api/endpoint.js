import apiAxios from '@/utils/apiAxios'

export const getEndpointList = (...args) => apiAxios.get('/blog/api/v1/endpoint/list', ...args)
