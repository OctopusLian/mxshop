import axios from 'axios';


let host = 'http://shop.projectsedu.com';
let goodsUrl = "http://124.223.201.78:8022/g"
let orderUrl = "http://124.223.201.78:8023/o"
let userUrl = "http://124.223.201.78:8021/u"
let userOpUrl = "http://124.223.201.78:8027/up"
export const ossUrl = "http://124.223.201.78:8029"

// let host = 'http://127.0.0.1:8000';

//上传文件
export const upload = (url,params) => { return axios.post(url,params) }
//获取商品类别信息
export const queryCategorygoods = params => { return axios.get(`${host}/indexgoods/`) }

// //获取首页中的新品
// export const newGoods = params => { return axios.get(`${host}/newgoods/`) }

//获取轮播图
export const bannerGoods = params => { return axios.get(`${goodsUrl}/v1/banners`) }

//获取商品类别信息
export const getCategory = params => {
  if('id' in params){
    return axios.get(`${goodsUrl}/v1/categorys/`+params.id);
  }
  else {
    return axios.get(`${goodsUrl}/v1/categorys`, params);
  }
};


//获取热门搜索关键词
export const getHotSearch = params => { return axios.get(`${host}/hotsearchs`) }


//获取验证码
export function getCaptcha(params) {
  return axios.get(userUrl+'/v1/base/captcha')
}
//获取商品列表
export const getGoods = params => { return axios.get(`${goodsUrl}/v1/goods`, { params: params }) }

//商品详情
export const getGoodsDetail = goodId => { return axios.get(`${goodsUrl}/v1/goods/${goodId}`) }

//获取购物车商品
export const getShopCarts = params => { return axios.get(`${orderUrl}/v1/shopcarts`) }
// 添加商品到购物车
export const addShopCart = params => { return axios.post(`${orderUrl}/v1/shopcarts`, params) }
//更新购物车商品信息
export const updateShopCart = (goodsId, params) => { return axios.patch(`${orderUrl}/v1/shopcarts/`+goodsId, params) }
//删除某个商品的购物记录
export const deleteShopCart = goodsId => { return axios.delete(`${orderUrl}/v1/shopcarts/`+goodsId) }

//收藏
export const addFav = params => { return axios.post(`${userOpUrl}/v1/userfavs`, params) }

//取消收藏
export const delFav = goodsId => { return axios.delete(`${userOpUrl}/v1/userfavs/`+goodsId) }

export const getAllFavs = () => { return axios.get(`${userOpUrl}/v1/userfavs`) }

//判断是否收藏getAllFavs
export const getFav = goodsId => { return axios.get(`${userOpUrl}/v1/userfavs/`+goodsId) }

//登录
export const login = params => {
  return axios.post(`${userUrl}/v1/user/pwd_login`, params)
}

//注册

export const register = parmas => { return axios.post(`${userUrl}/v1/user/register`, parmas) }

//短信
export const getMessage = parmas => { return axios.post(`${userUrl}/v1/base/send_sms`, parmas) }


//获取用户信息
export const getUserDetail = () => { return axios.get(`${userUrl}/v1/user/detail`) }

//修改用户信息
export const updateUserInfo = params => { return axios.patch(`${userUrl}/v1/user/update`, params) }


//获取订单
export const getOrders = () => { return axios.get(`${orderUrl}/v1/orders`) }
//删除订单
export const delOrder = orderId => { return axios.delete(`${orderUrl}/v1/orders/`+orderId) }
//添加订单
export const createOrder = params => {return axios.post(`${orderUrl}/v1/orders`, params)}
//获取订单详情
export const getOrderDetail = orderId => {return axios.get(`${orderUrl}/v1/orders/`+orderId)}


//获取留言
export const getMessages = () => {return axios.get(`${userOpUrl}/v1/message`)}

//添加留言
export const addMessage = params => {return axios.post(`${userOpUrl}/v1/message`, params)}

//删除留言
export const delMessages = messageId => {return axios.delete(`${userOpUrl}/v1/message/`+messageId)}

//添加收货地址
export const addAddress = params => {return axios.post(`${userOpUrl}/v1/address`, params)}

//删除收货地址
export const delAddress = addressId => {return axios.delete(`${userOpUrl}/v1/address/`+addressId)}

//修改收货地址
export const updateAddress = (addressId, params) => {return axios.patch(`${userOpUrl}/v1/address/`+addressId, params)}

//获取收货地址
export const getAddress = () => {return axios.get(`${userOpUrl}/v1/address`)}
