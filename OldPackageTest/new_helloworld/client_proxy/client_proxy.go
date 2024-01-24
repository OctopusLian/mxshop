package client_proxy

import (
	"OldPackageTest/new_helloworld/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// NewHelloServiceClient 在go语言中没有类、对象 就意味着没有初始化方法
func NewHelloServiceClient(protcol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protcol, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
