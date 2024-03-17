package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
}

func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 135487,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	//TestCreateUser()
	TestGetSubCategoryList()
	TestGetCategoryList()

	conn.Close()
}
