import request from '@/utils/request'
export function policy() {
  return request({
    url:'/token',
    method:'get',
  })
}
