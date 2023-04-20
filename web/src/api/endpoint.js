import apiAxios from '@/utils/apiAxios'

export const getEndpointList = (...args) => apiAxios.get('/blog/api/v1/endpoint/list', ...args)
export const postAddEndpoint = (...args) => apiAxios.post('/blog/api/v1/endpoint/add', ...args)
export const putEditEndpoint = (...args) => apiAxios.put('/blog/api/v1/endpoint/edit', ...args)
export const deleteEndpoint = (...args) => apiAxios.delete('/blog/api/v1/endpoint/delete', ...args)
