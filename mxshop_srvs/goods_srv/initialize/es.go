/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 21:56:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:57:46
 */
package initialize

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"os"

	"github.com/olivere/elastic/v7"
)

func InitEs() {
	//初始化连接
	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	//新建mapping和index
	exists, err := global.EsClient.IndexExists(model.EsGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err = global.EsClient.CreateIndex(model.EsGoods{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
