package service

import (
	"context"
	"favoritesvr/config"
	"favoritesvr/log"
	"favoritesvr/utils"
	"strconv"

	"github.com/taosu0216/DouFlick/pkg/pb"
)

type FavoriteService struct {
	pb.UnimplementedFavoriteServiceServer
}

func (FavoriteService) FavoriteAction(ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	if req.ActionType == 1 {
		//点赞
		err := utils.FavoriteAdd(req.UserId, req.VideoId)
		if err != nil {
			log.Error("FavoriteAdd error", err)
			return nil, err
		}
	} else {
		//取消点赞
		err := utils.FavoriteDelete(req.UserId, req.VideoId)
		if err != nil {
			log.Error("FavoriteDelete error", err)
			return nil, err
		}
	}
	return &pb.FavoriteActionResponse{CommonRsp: &pb.CommonResponse{
		Code: config.SuccessCode,
		Msg:  config.SuccessMsg,
	}}, nil
}

func (FavoriteService) GetFavoriteList(ctx context.Context, req *pb.GetFavoriteListRequest) (*pb.GetFavoriteListResponse, error) {
	favoriteList, err := utils.GetFavoriteList(req.UserId)
	if err != nil {
		log.Error("GetFavoriteList error", err)
		return nil, err
	}
	return &pb.GetFavoriteListResponse{VideoIdList: favoriteList}, nil
}

func (FavoriteService) IsFavoriteDict(ctx context.Context, req *pb.IsFavoriteDictRequest) (*pb.IsFavoriteDictResponse, error) {
	dict := make(map[string]bool)
	for _, unit := range req.FavoriteUintList {
		isFav, err := utils.IsFavoriteDict(unit.UserId, unit.VideoId)
		if err != nil {
			log.Error("IsFavoriteDict error", err)
			return nil, err
		}
		key := strconv.FormatInt(unit.UserId, 10) + "_" + strconv.FormatInt(unit.VideoId, 10)
		dict[key] = isFav
	}
	return &pb.IsFavoriteDictResponse{IsFavoriteDict: dict}, nil
}
