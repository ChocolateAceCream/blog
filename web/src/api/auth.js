import apiAxios from '@/utils/apiAxios'

export const getUserList = (...args) => apiAxios.get('/blog/api/v1/user/userList', ...args)
// export const putResetPassword = (...args) => apiAxios.put('/cwifi-vac-admin-provider/vac/sys/user/password/change', ...args)
// export const putUpdateUserInfo = (...args) => apiAxios.put('/cwifi-vac-admin-provider/vac/sys/user/update', ...args)
// export const postSaveUser = (...args) => apiAxios.post('/cwifi-vac-admin-provider/vac/sys/user/add', ...args)
