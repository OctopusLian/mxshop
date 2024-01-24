package server_proxy

import (
	"OldPackageTest/new_helloworld/hanlder"
	"net/rpc"
)

type HelloService interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService 如果做到解耦 - 我们关系的是函数 鸭子类型
func RegisterHelloService(srv HelloService) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
