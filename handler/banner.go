/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-26 22:38:30
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-26 22:52:14
 */
package handler

import (
	"context"
	"mxshop/global"
	"mxshop/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BannerListResponse struct {
	Total int32             `json:"total,omitempty"`
	Data  []*BannerResponse `json:"data,omitempty"`
}

type BannerResponse struct {
	Id    int32  `json:"id,omitempty"`
	Index int32  `json:"index,omitempty"`
	Image string `json:"image,omitempty"`
	Url   string `json:"url,omitempty"`
}

type BannerRequest struct {
	Id    int32  `json:"id,omitempty"`
	Index int32  `json:"index,omitempty"`
	Image string `json:"image,omitempty"`
	Url   string `json:"url,omitempty"`
}

//轮播图
func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*BannerListResponse, error) {
	bannerListResponse := BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil
}

func (s *GoodsServer) CreateBanner(ctx context.Context, req *BannerRequest) (*BannerResponse, error) {
	banner := model.Banner{}

	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url

	global.DB.Save(&banner)

	return &BannerResponse{Id: banner.ID}, nil
}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *BannerRequest) (*emptypb.Empty, error) {
	var banner model.Banner

	if result := global.DB.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}

	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}

	global.DB.Save(&banner)

	return &emptypb.Empty{}, nil
}
