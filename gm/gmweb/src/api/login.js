import request from '@/utils/request'

export function loginByUsername(username, password) {
  return request({
    url: '/user/login',
    method: 'post',
    data: {
      userName: username,
      password: password
    }
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}

export function getUserInfo(token) {
  return request({
    url: '/user/get_info',
    method: 'post',
    data: {}
  })
}

