import apiAxios from '@/utils/apiAxios'

export const getUserList = (...args) => apiAxios.get('/backend/api/v1/user/list', ...args)
export const postVerificationCode = (...args) => apiAxios.post('/backend/api/public/auth/captcha', ...args)
export const postLogin = (...args) => apiAxios.post('/backend/api/public/auth/login', ...args)
export const putResetPassword = (...args) => apiAxios.put('/backend/api/v1/user/resetPassword', ...args)
export const postSendEmailCode = (...args) => apiAxios.post('/backend/api/public/auth/sendEmailCode', ...args)
export const postRegister = (...args) => apiAxios.post('/backend/api/public/auth/register', ...args)
// export const putResetPassword = (...args) => apiAxios.put('/cwifi-vac-admin-provider/vac/sys/user/password/change', ...args)
export const putUpdateUserInfo = (...args) => apiAxios.put('/cwifi-vac-admin-provider/vac/sys/user/update', ...args)
