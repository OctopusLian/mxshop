package main

import (
	"OldPackageTest/new_helloworld/hanlder"
	"net"
	"net/rpc"

	"OldPackageTest/new_helloworld/server_proxy"
)

func main() {
	//1. 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	//2. 注册处理逻辑 handler
	_ = server_proxy.RegisterHelloService(&hanlder.NewHelloService{})
	//3. 启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候，
		go rpc.ServeConn(conn)
	}

	//一连串的代码大部分都是net的包好像和rpc没有关系
	//不行。rpc调用中有几个问题需要解决 1. call id 2. 序列化和反序列化 编码和解码
	//python下的开发而言 这个就显得不好用
	//可以跨语言调用呢 1. go语言的rpc的序列化协议是什么（Gob） 2. 能否替换成常见的序列化
}
