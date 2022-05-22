/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-21 22:22:56
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 16:02:42
 */
package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/user_srv/proto"
	"net"

	"mxshop_srvs/user_srv/handler"

	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("PORT", 50051, "端口号")

	flag.Parse()
	fmt.Println(*IP)
	fmt.Println(*Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed_to_listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed_to_start_grpc:" + err.Error())
	}
}
