package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1. 建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}
	var reply string //string有默认值
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
