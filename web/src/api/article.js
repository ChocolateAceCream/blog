import apiAxios from '@/utils/apiAxios'

export const getArticleFile = (...args) => apiAxios.get('/backend/api/v1/article/preview', ...args)

// export const getArticleFile = () => apiAxios.get('/backend/api/v1/article/preview', { responseType: 'blob' })
