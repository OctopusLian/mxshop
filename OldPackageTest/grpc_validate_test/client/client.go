package main

import (
	"OldPackageTest/grpc_validate_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type customCredential struct{}

func main() {
	var opts []grpc.DialOption

	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//rsp, _ := c.Search(context.Background(), &empty.Empty{})
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "bobby@imooc.com",
		Mobile: "18888888888",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}
