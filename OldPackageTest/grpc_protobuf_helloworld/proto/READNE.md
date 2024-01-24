
```shell
# 在proto下执行
protoc -I . helloworld.proto --go_out=plugins=grpc:.


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```
