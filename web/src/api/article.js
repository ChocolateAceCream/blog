import apiAxios from '@/utils/apiAxios'

export const getArticleFile = (...args) => apiAxios.get('/backend/api/v1/article/preview', ...args)
export const postAddArticle = () => apiAxios.post('/backend/api/v1/article/add')
export const putEditArticle = (...args) => apiAxios.put('/backend/api/v1/article/edit', ...args)
export const getArticleList = (...args) => apiAxios.get('/backend/api/v1/article/list', ...args)
export const getArticleSearchList = (...args) => apiAxios.get('/backend/api/v1/article/search', ...args)
export const deleteArticle = (...args) => apiAxios.delete('/backend/api/v1/article/delete', ...args)

// export const getArticleFile = () => apiAxios.get('/backend/api/v1/article/preview', { responseType: 'blob' })
