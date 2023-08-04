import apiAxios from '@/utils/apiAxios'
export const getCommentList = (...args) => apiAxios.get('/backend/api/v1/comment/list', ...args)
export const postAddComment = (...args) => apiAxios.post('/backend/api/v1/comment/add', ...args)
export const deleteComment = (...args) => apiAxios.delete('/backend/api/v1/comment/delete', ...args)
