import apiAxios from '@/utils/apiAxios'

export const getArticleFile = (...args) => apiAxios.get('/backend/api/v1/article/preview', ...args)
export const postAddArticle = () => apiAxios.post('/backend/api/v1/article/add')
export const putEditArticle = (...args) => apiAxios.put('/backend/api/v1/article/edit', ...args)

// export const getArticleFile = () => apiAxios.get('/backend/api/v1/article/preview', { responseType: 'blob' })
