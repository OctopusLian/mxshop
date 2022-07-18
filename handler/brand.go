package handler

import (
	"context"
	"mxshop/global"
	"mxshop/model"
)

type BrandInfoResponse struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type BrandListResponse struct {
	Total int32                `json:"total,omitempty"`
	Data  []*BrandInfoResponse `json:"data,omitempty"`
}

type CategoryBrandFilterRequest struct {
	Pages       int32 `json:"pages,omitempty"`
	PagePerNums int32 `json:"pagePerNums,omitempty"`
}

type BrandFilterRequest struct {
	Pages       int32 `json:"pages,omitempty"`
	PagePerNums int32 `json:"pagePerNums,omitempty"`
}

//品牌和轮播图
func (s *GoodsServer) BrandList(ctx context.Context, req *BrandFilterRequest) (*BrandListResponse, error) {
	brandListResponse := BrandListResponse{}

	var brands []model.Brands
	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponses []*BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}
