import apiAxios from '@/utils/apiAxios'

export const getEndpointList = (...args) => apiAxios.get('/backend/api/v1/endpoint/list', ...args)
export const postAddEndpoint = (...args) => apiAxios.post('/backend/api/v1/endpoint/add', ...args)
export const putEditEndpoint = (...args) => apiAxios.put('/backend/api/v1/endpoint/edit', ...args)
export const deleteEndpoint = (...args) => apiAxios.delete('/backend/api/v1/endpoint/delete', ...args)
