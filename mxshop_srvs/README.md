<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-05-21 23:16:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:59:22
-->
# 生鲜电商系统服务端  

## 生成Proto  

用户：`protoc -I . user.proto --go_out=plugins=grpc:.`

## 微服务拆分  

- [用户](./user_srv/)  
- [商品](./goods_srv/)  