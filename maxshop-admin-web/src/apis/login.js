import request from '@/utils/request'
let userUrl = "http://192.168.0.103:8021"
export function login(params) {
  return request({
    url:userUrl+'/u/v1/user/pwd_login',
    method:'post',
    data:params
  })
}
export function getCaptcha(params) {
  return request({
    url:userUrl+'/u/v1/base/captcha',
    method:'get'
  })
}

export function createBrand(data) {
  return request({
    url:'/brand/create',
    method:'post',
    data:data
  })
}
export function updateShowStatus(data) {
  return request({
    url:'/brand/update/showStatus',
    method:'post',
    data:data
  })
}

export function updateFactoryStatus(data) {
  return request({
    url:'/brand/update/factoryStatus',
    method:'post',
    data:data
  })
}

export function deleteBrand(id) {
  return request({
    url:'/brand/delete/'+id,
    method:'get',
  })
}

export function getBrand(id) {
  return request({
    url:'/brand/'+id,
    method:'get',
  })
}

export function updateBrand(id,data) {
  return request({
    url:'/brand/update/'+id,
    method:'post',
    data:data
  })
}

