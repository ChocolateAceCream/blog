import apiAxios from '@/utils/apiAxios'

export const postUploadFile = (...args) => apiAxios.post('/backend/api/v1/oss/upload', ...args)
