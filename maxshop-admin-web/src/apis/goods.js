import request from '@/utils/request'

let host = 'http://shop.projectsedu.com';
// let baseUrl = "http://39.107.30.137:8000"
let goodsUrl = "http://124.223.201.78:8022"
let orderUrl = "http://124.223.201.78:8023"
let userUrl = "http://124.223.201.78:8021"
let userOpUrl = "http://124.223.201.78:8027"
export let ossUrl = "http://124.223.201.78:8029"
// 分类
export function getCategorys(params) {
  return request({
    url:goodsUrl+'/g/v1/categorys',
    method:'get',
    params:params
  })
}
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
export function getCategoryDetail(id) {
  return request({
    url:goodsUrl+'/g/v1/categorys/'+id,
    method:'get'
  })
}
export function postCategorys(params) {
  return request({
    url:goodsUrl+'/g/v1/categorys',
    method:'post',
    data:params
  })
}
export function putCategorys(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorys/'+id,
    method:'put',
    data:params
  })
}
export function deleteCategorys(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorys/'+id,
    method:'delete',
    data:params
  })
}
// 品牌
export function getBrands(params) {
  return request({
    url:goodsUrl+'/g/v1/brands',
    method:'get',
    params:params
  })
}
export function getBrandsByCate(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands/'+id,
    method:'get',
    params:params
  })
}
export function createBrand(data) {
  return request({
    url:goodsUrl+'/g/v1/brands',
    method:'post',
    data:data
  })
}
export function putBrands(id,params) {
  return request({
    url:goodsUrl+'/brands/'+id,
    method:'put',
    data:params
  })
}
export function deleteBrands(id,params) {
  return request({
    url:goodsUrl+'/g/v1/brands/'+id,
    method:'delete',
    data:params
  })
}
// 品牌+分类关系
export function getBrandToCate(params) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands',
    method:'get',
    params:params
  })
}

export function getBrandToCateDetail(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands/'+id,
    method:'get',
    params:params
  })
}
export function createBrandToCate(data) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands',
    method:'post',
    data:data
  })
}
export function putBrandToCate(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands/'+id,
    method:'put',
    data:params
  })
}
export function deleteBrandToCate(id,params) {
  return request({
    url:goodsUrl+'/g/v1/categorybrands/'+id,
    method:'delete',
    data:params
  })
}
// 商品管理
export function getGoods(params) {
  return request({
    url:goodsUrl+'/g/v1/goods',
    method:'get',
    params:params
  })
}
export function getGoodsEach(params) {
  return request({
    url:goodsUrl+'/g/v1/goods/'+params,
    method:'get',
  })
}

export function createGoods(data) {
  return request({
    url:goodsUrl+'/g/v1/goods',
    method:'post',
    data:data
  })
}
export function putGoods(id,params) {
  return request({
    url:goodsUrl+'/g/v1/goods/'+id,
    method:'put',
    data:params
  })
}
export function deleteGoods(id,params) {
  return request({
    url:goodsUrl+'/g/v1/goods/'+id,
    method:'delete',
    data:params
  })
}
export function putGoodsStatus(id,params) {
  return request({
    url:goodsUrl+'/g/v1/goods/'+id,
    method:'patch',
    data:params
  })
}



// 订单
export function getOrder(params) {
  return request({
    url:orderUrl+'/o/v1/orders',
    method:'get',
    params:params
  })
}
export function getOrderEach(id,params) {
  return request({
    url:orderUrl+'/o/v1/orders/'+id,
    method:'get',
    params:params
  })
}

export function createOrder(data) {
  return request({
    url:orderUrl+'/orders',
    method:'post',
    data:data
  })
}
export function putOrder(id,params) {
  return request({
    url:orderUrl+'/orders/'+id,
    method:'put',
    data:params
  })
}
export function deleteOrder(id,params) {
  return request({
    url:orderUrl+'/orders/'+id,
    method:'delete',
    data:params
  })
}

// 留言
export function getMessage(params) {
  return request({
    url:userOpUrl+'/up/v1/message',
    method:'get',
    params:params
  })
}

export function createMessage(data) {
  return request({
    url:userOpUrl+'/message',
    method:'post',
    data:data
  })
}
export function putMessage(id,params) {
  return request({
    url:userOpUrl+'/message/'+id,
    method:'put',
    data:params
  })
}
export function deleteMessage(id,params) {
  return request({
    url:userOpUrl+'/message/'+id,
    method:'delete',
    data:params
  })
}
// 收藏
export function getuserfav(params) {
  return request({
    url:userOpUrl+'/userfav',
    method:'get',
    params:params
  })
}

export function createuserfav(data) {
  return request({
    url:userOpUrl+'/userfav',
    method:'post',
    data:data
  })
}
export function putuserfav(id,params) {
  return request({
    url:userOpUrl+'/userfav/'+id,
    method:'put',
    data:params
  })
}
export function deleteuserfav(id,params) {
  return request({
    url:userOpUrl+'/userfav/'+id,
    method:'delete',
    data:params
  })
}
//用户地址、
export function getaddress(params) {
  return request({
    url:userOpUrl+'/up/v1/address',
    method:'get',
    params:params
  })
}
export function getaddressEach(params) {
  return request({
    url:userOpUrl+'/address',
    method:'get',
    params:params
  })
}

export function createaddress(data) {
  return request({
    url:userOpUrl+'/address',
    method:'post',
    data:data
  })
}
export function putaddress(id,params) {
  return request({
    url:userOpUrl+'/address/'+id,
    method:'put',
    data:params
  })
}
export function deleteaddress(id,params) {
  return request({
    url:userOpUrl+'/address/'+id,
    method:'delete',
    data:params
  })
}

// 用户列表

export function getUserList(params) {
  return request({
    url:userUrl+'/u/v1/user',
    method:'get',
    params:params
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
    url:goodsUrl+'/brand/update/factoryStatus',
    method:'post',
    data:data
  })
}

export function deleteBrand(id) {
  return request({
    url:goodsUrl+'/brand/delete/'+id,
    method:'get',
  })
}

export function getBrand(id) {
  return request({
    url:goodsUrl+'/brand/'+id,
    method:'get',
  })
}

export function updateBrand(id,data) {
  return request({
    url:goodsUrl+'/brand/update/'+id,
    method:'post',
    data:data
  })
}


export function getBanners(params) {
  return request({
    url:goodsUrl+'/g/v1/banners',
    method:'get',
    params:params
  })
}

export function createBanner(data) {
  return request({
    url:goodsUrl+'/g/v1/banners',
    method:'post',
    data:data
  })
}
export function putBanner(id,params) {
  return request({
    url:goodsUrl+'/g/v1/banners/'+id,
    method:'put',
    data:params
  })
}
export function deleteBanners(id,params) {
  return request({
    url:goodsUrl+'/g/v1/banners/'+id,
    method:'delete',
    data:params
  })
}
