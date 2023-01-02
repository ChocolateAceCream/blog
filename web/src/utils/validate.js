// 密码
const validatePassword = (rule, value, callback) => {
  if (!new RegExp('^(?=.*[a-zA-Z0-9].*)(?=.*[a-zA-Z\\W].*)(?=.*[0-9\\W].*).{8,20}$').test(value)) {
    //  if (!/(?!^(\d+|[A-Za-z]+|[~!@$%^&*()+]+)$)^([A-Za-z0-9]+|[~!@$%^&*()+]){6,20}$/.test(value)) {
    callback(new Error('8~20位, 至少包含数字、字母、特殊字符中的两种'))
  }
  callback()
}
//  ip地址
function validateIp(ip) {
  const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
  return reg.test(ip)
}

function validateFormIp(rule, value, callback) {
  const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
  // return reg.test(ip)
  if (!value || reg.test(value)) {
    callback()
  } else {
    callback(new Error('请输入正确的IP地址'))
  }
}
// 网关id （数字 20字符内）
function validateDeviceId(rule, value, callback) {
  if (value !== '' && !/^[0-9]{1,20}$/.test(value)) {
    callback(new Error('级联网关ID为数字（20字符内）'))
  } else {
    callback()
  }
}
// 端口 （数字 5字符内）
function validatePort(rule, value, callback) {
  if (!!value && !/^[0-9]{1,5}$/.test(value)) {
    callback(new Error('端口为数字（5字符内）'))
  } else {
    callback()
  }
}
// 外域编码（数字 20字符内）
function validateCascadeDeviceId(rule, value, callback) {
  if (value !== '' && !/^[0-9]{1,20}$/.test(value)) {
    callback(new Error('外域编码为数字（20字符内）'))
  } else {
    callback()
  }
}
// 邮箱校验
function validateEmail(rule, value, callback, source) {
  if (value && !/^[a-z0-9_\.]{1,}@[a-z0-9]{1,}([-\.]{1}[a-z0-9]{1,}){0,}[\.]{1}[a-z]{1,4}$/.test(value)) {
    callback(new Error('请输入正确的邮箱地址'))
  } else {
    callback()
  }
}
// 手机号校验
function validateCellphone(rule, value, callback, source) {
  if (value !== '' && !/^1[3|4|5|7|8][0-9]{9}$/.test(value)) {
    callback(new Error('请输入正确的11位手机号'))
  } else {
    callback()
  }
}
// 校验数字（限制只能输入数字，且不能输入小数）
function validateNum(rule, value, callback, source) {
  if (value && !/^[0-9]+$/.test(value)) {
    callback(new Error('请输入数字'))
  } else {
    callback()
  }
}

function validateName(rule, value, callback, source) {
  if (value !== '' && !/^[0-9a-zA-Z\u4e00-\u9fa5]+$/.test(value)) {
    callback(new Error('请输入中英文数字'))
  } else {
    callback()
  }
}
// 校验触发器条件
function validateTermItem(rule, value, callback) {
  if (value.value === '') {
    callback(new Error('请输入域名'))
  } else {
    callback()
  }
}

function generateRegexValidator(options) {
  function regexValidator(rule, value, callback) {
    if (options.regex && value && !value.match(options.regex)) {
      return callback(new Error('匹配失败'))
    }
    callback()
  }
  return Object.assign({ validator: regexValidator, message: '格式错误', trigger: 'blur' }, options)
}

export {
  validatePassword,
  validateIp,
  validateFormIp,
  validateDeviceId,
  validatePort,
  validateCascadeDeviceId,
  validateEmail,
  validateCellphone,
  validateNum,
  validateName,
  validateTermItem,
  generateRegexValidator
}
