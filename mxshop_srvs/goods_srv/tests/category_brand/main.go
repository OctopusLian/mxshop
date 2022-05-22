/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 21:56:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:58:16
 */
package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetCategoryBrandList() {
	rsp, err := brandClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.Data)
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
	//TestGetCategoryList()
	TestGetCategoryBrandList()

	conn.Close()
}
