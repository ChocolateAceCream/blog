import apiAxios from '@/utils/apiAxios'
export const getNotificationList = (...args) => apiAxios.get('/backend/api/v1/notification/list', ...args)
export const getUnreadNotificationCount = () => apiAxios.get('/backend/api/v1/notification/unreadCount')
export const deleteNotification = (...args) => apiAxios.delete('/backend/api/v1/notification/delete', ...args)
export const readNotification = (...args) => apiAxios.patch('/backend/api/v1/notification/read', ...args)
