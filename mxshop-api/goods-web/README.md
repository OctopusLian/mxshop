# 商品Web服务  

## 生成proto

```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
goods.proto
```
