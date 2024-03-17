package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/inventory_srv/proto"
	"sync"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func TestSetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}

func TestInvDetail(goodsId int32) {
	rsp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Num)
}

func TestSell(wg *sync.WaitGroup) {
	/*
		1. 第一件扣减成功： 第二件： 1. 没有库存信息 2. 库存不足
		2. 两件都扣减成功
	*/
	defer wg.Done()
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},
			//{GoodsId: 422, Num: 30},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("库存扣减成功")
}

func TestReback() {
	_, err := invClient.Reback(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 10},
			{GoodsId: 422, Num: 30},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("归还成功")
}

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}

func main() {
	Init()
	//var i int32
	//for i = 421; i<=840; i++ {
	//	TestSetInv(i, 100)
	//}
	//并发情况之下 库存无法正确的扣减
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go TestSell(&wg)
	}

	wg.Wait()

	//TestInvDetail(421)
	//TestSell()
	//TestReback()
	conn.Close()
}
