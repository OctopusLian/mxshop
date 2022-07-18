package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"mxshop/global"
	"mxshop/model"

	"github.com/olivere/elastic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GoodsServer struct{}

type GoodsInfoResponse struct {
	Id              int32                      `json:"id,omitempty"`
	CategoryId      int32                      `json:"categoryId,omitempty"`
	Name            string                     `json:"name,omitempty"`
	GoodsSn         string                     `json:"goodsSn,omitempty"`
	ClickNum        int32                      `json:"clickNum,omitempty"`
	SoldNum         int32                      `json:"soldNum,omitempty"`
	FavNum          int32                      `json:"favNum,omitempty"`
	MarketPrice     float32                    `json:"marketPrice,omitempty"`
	ShopPrice       float32                    `json:"shopPrice,omitempty"`
	GoodsBrief      string                     `json:"goodsBrief,omitempty"`
	GoodsDesc       string                     `json:"goodsDesc,omitempty"`
	ShipFree        bool                       `json:"shipFree,omitempty"`
	Images          []string                   `json:"images,omitempty"`
	DescImages      []string                   `json:"descImages,omitempty"`
	GoodsFrontImage string                     `json:"goodsFrontImage,omitempty"`
	IsNew           bool                       `json:"isNew,omitempty"`
	IsHot           bool                       `json:"isHot,omitempty"`
	OnSale          bool                       `json:"onSale,omitempty"`
	AddTime         int64                      `json:"addTime,omitempty"`
	Category        *CategoryBriefInfoResponse `json:"category,omitempty"`
	Brand           *BrandInfoResponse         `json:"brand,omitempty"`
}

type CategoryBriefInfoResponse struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s *GoodsServer) ModelToResponse(goods model.Goods) GoodsInfoResponse {
	return GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

type GoodsFilterRequest struct {
	PriceMin    int32  `json:"priceMin,omitempty"`
	PriceMax    int32  `json:"priceMax,omitempty"`
	IsHot       bool   `json:"isHot,omitempty"`
	IsNew       bool   `json:"isNew,omitempty"`
	IsTab       bool   `json:"isTab,omitempty"`
	TopCategory int32  `json:"topCategory,omitempty"`
	Pages       int32  `json:"pages,omitempty"`
	PagePerNums int32  `json:"pagePerNums,omitempty"`
	KeyWords    string `json:"keyWords,omitempty"`
	Brand       int32  `json:"brand,omitempty"`
}

type GoodsListResponse struct {
	Total int32                `json:"total,omitempty"`
	Data  []*GoodsInfoResponse `json:"data,omitempty"`
}

func (s *GoodsServer) GoodsList(ctx context.Context, req *GoodsFilterRequest) (*GoodsListResponse, error) {
	//使用es的目的是搜索出商品的id来，通过id拿到具体的字段信息是通过mysql来完成
	//我们使用es是用来做搜索的， 是否应该将所有的mysql字段全部在es中保存一份
	//es用来做搜索，这个时候我们一般只把搜索和过滤的字段信息保存到es中
	//es可以用来当做mysql使用， 但是实际上mysql和es之间是互补的关系， 一般mysql用来做存储使用，es用来做搜索使用
	//es想要提高性能， 就要将es的内存设置的够大， 1k 2k

	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	goodsListResponse := &GoodsListResponse{}

	//match bool 复合查询
	q := elastic.NewBoolQuery()
	localDB := global.DB.Model(model.Goods{})
	if req.KeyWords != "" {
		q = q.Must(elastic.NewMultiMatchQuery(req.KeyWords, "name", "goods_brief"))
	}
	if req.IsHot {
		localDB = localDB.Where(model.Goods{IsHot: true})
		q = q.Filter(elastic.NewTermQuery("is_hot", req.IsHot))
	}
	if req.IsNew {
		q = q.Filter(elastic.NewTermQuery("is_new", req.IsNew))
	}

	if req.PriceMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Gte(req.PriceMin))
	}
	if req.PriceMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Lte(req.PriceMax))
	}

	if req.Brand > 0 {
		q = q.Filter(elastic.NewTermQuery("brands_id", req.Brand))
	}

	//通过category去查询商品
	var subQuery string
	categoryIds := make([]interface{}, 0)
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", req.TopCategory)
		}

		type Result struct {
			ID int32
		}
		var results []Result
		global.DB.Model(model.Category{}).Raw(subQuery).Scan(&results)
		for _, re := range results {
			categoryIds = append(categoryIds, re.ID)
		}

		//生成terms查询
		q = q.Filter(elastic.NewTermsQuery("category_id", categoryIds...))
	}

	//分页
	if req.Pages == 0 {
		req.Pages = 1
	}

	switch {
	case req.PagePerNums > 100:
		req.PagePerNums = 100
	case req.PagePerNums <= 0:
		req.PagePerNums = 10
	}
	result, err := global.EsClient.Search().Index(model.EsGoods{}.GetIndexName()).Query(q).From(int(req.Pages)).Size(int(req.PagePerNums)).Do(context.Background())
	if err != nil {
		return nil, err
	}

	goodsIds := make([]int32, 0)
	goodsListResponse.Total = int32(result.Hits.TotalHits)
	for _, value := range result.Hits.Hits {
		goods := model.EsGoods{}
		_ = json.Unmarshal(*value.Source, &goods)
		goodsIds = append(goodsIds, goods.ID)
	}

	//查询id在某个数组中的值
	var goods []model.Goods
	re := localDB.Preload("Category").Preload("Brands").Find(&goods, goodsIds)
	if re.Error != nil {
		return nil, re.Error
	}

	for _, good := range goods {
		goodsInfoResponse := s.ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}

	return goodsListResponse, nil
}

type BatchGoodsIdInfo struct {
	Id []int32 `json:"id,omitempty"`
}

//现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *BatchGoodsIdInfo) (*GoodsListResponse, error) {
	goodsListResponse := &GoodsListResponse{}
	var goods []model.Goods

	//调用where并不会真正执行sql 只是用来生成sql的 当调用find， first才会去执行sql，
	result := global.DB.Where(req.Id).Find(&goods)
	for _, good := range goods {
		goodsInfoResponse := s.ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)
	return goodsListResponse, nil
}

type GoodInfoRequest struct {
	Id int32 `json:"id,omitempty"`
}

func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *GoodInfoRequest) (*GoodsInfoResponse, error) {
	var goods model.Goods

	if result := global.DB.Preload("Category").Preload("Brands").First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	goodsInfoResponse := s.ModelToResponse(goods)
	return &goodsInfoResponse, nil
}

type CreateGoodsInfo struct {
	Id              int32    `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	GoodsSn         string   `json:"goodsSn,omitempty"`
	Stocks          int32    `json:"stocks,omitempty"` //库存，
	MarketPrice     float32  `json:"marketPrice,omitempty"`
	ShopPrice       float32  `json:"shopPrice,omitempty"`
	GoodsBrief      string   `json:"goodsBrief,omitempty"`
	GoodsDesc       string   `json:"goodsDesc,omitempty"`
	ShipFree        bool     `json:"shipFree,omitempty"`
	Images          []string `json:"images,omitempty"`
	DescImages      []string `json:"descImages,omitempty"`
	GoodsFrontImage string   `json:"goodsFrontImage,omitempty"`
	IsNew           bool     `json:"isNew,omitempty"`
	IsHot           bool     `json:"isHot,omitempty"`
	OnSale          bool     `json:"onSale,omitempty"`
	CategoryId      int32    `json:"categoryId,omitempty"`
	BrandId         int32    `json:"brandId,omitempty"`
}

func (s *GoodsServer) CreateGoods(ctx context.Context, req *CreateGoodsInfo) (*GoodsInfoResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	//先检查redis中是否有这个token
	//防止同一个token的数据重复插入到数据库中，如果redis中没有这个token则放入redis
	//这里没有看到图片文件是如何上传， 在微服务中 普通的文件上传已经不再使用
	goods := model.Goods{
		Brands:          brand,
		BrandsID:        brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		OnSale:          req.OnSale,
	}

	//srv之间互相调用了
	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}

type DeleteGoodsInfo struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (s *GoodsServer) DeleteGoods(ctx context.Context, req *DeleteGoodsInfo) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Goods{BaseModel: model.BaseModel{ID: req.Id}}, req.Id); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateGoods(ctx context.Context, req *CreateGoodsInfo) (*emptypb.Empty, error) {
	var goods model.Goods

	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	goods.Brands = brand
	goods.BrandsID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.GoodsBrief = req.GoodsBrief
	goods.ShipFree = req.ShipFree
	goods.Images = req.Images
	goods.DescImages = req.DescImages
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale

	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
