import apiAxios from '@/utils/apiAxios'
export const getReplyList = (...args) => apiAxios.get('/backend/api/v1/reply/list', ...args)
export const postAddReply = (...args) => apiAxios.post('/backend/api/v1/reply/add', ...args)
export const deleteReply = (...args) => apiAxios.delete('/backend/api/v1/reply/delete', ...args)
export const likeReply = (...args) => apiAxios.post('/backend/api/v1/reply/like', ...args)
