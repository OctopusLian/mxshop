package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"OldPackageTest/grpc_test/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (
	*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 实例化grpc
	g := grpc.NewServer()
	// 注册server
	proto.RegisterGreeterServer(g, &Server{})
	// 启动
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
