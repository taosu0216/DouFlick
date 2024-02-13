package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
	"usersvr/config"
	"usersvr/log"
	"usersvr/middleware/lock"
	"usersvr/utils"
)

/*
	      需要满足的方法
		  rpc CheckPassWord(CheckPassWordRequest) returns (CheckPassWordResponse); //检查密码
		  rpc GetUserInfo(UserInfoRequest) returns (UserInfoResponse); //获取用户信息
		  rpc GetUserInfoList(UserInfoListRequest) returns (UserInfoListResponse); //获取用户信息列表
		  rpc Register(RegisterRequest) returns (RegisterResponse); //注册
		  rpc GetUserInfoDict(UserInfoDictRequest) returns (UserInfoDictResponse); //获取用户信息字典
	      rpc CacheChangeUserCount(CacheChangeUserCountRequest) returns (CacheChangeUserCountResponse); //修改用户信息
	      TODO: CacheGetAuthor未实现
	      rpc CacheGetAuthor(CacheGetAuthorRequest) returns (CacheGetAuthorResponse); //获取视频的作者

		  // 数据库更新

		  // 更新我的获赞数
		  rpc UpdateMyFavouritedCount(UpdateMyFavouritedCountRequest) returns (UpdateMyFavouritedCountResponse);
		  // 更新我喜欢的视频数
		  rpc UpdateMyFavouriteCount(UpdateMyFavouriteCountRequest) returns (UpdateMyFavouriteCountResponse);
		  // 更新我的关注数
		  rpc UpdateMyFollowCount(UpdateMyFollowCountRequest) returns (UpdateMyFollowCountResponse);
	      // 更新我的粉丝数
		  rpc UpdateMyFollowerCount(UpdateMyFollowerCountRequest) returns (UpdateMyFollowerCountResponse);
*/
var (
	Secret              = []byte("DouFlick")
	TokenExpireDuration = time.Hour * 24 //过期时间
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

// UserService 满足rpc方式的接口 --pb.UnimplementedUserServiceServer对象才能调用rpc方法
type UserService struct {
	pb.UnimplementedUserServiceServer
}

// CheckPassWord 检查密码
func (UserService) CheckPassWord(ctx context.Context, req *pb.CheckPassWordRequest) (*pb.CheckPassWordResponse, error) {
	info, err := utils.GetUserInfoByAny(req.Username)
	if err != nil {
		log.Error("service/CheckPassWord() GetUser error is : ", err)
		return nil, err
	}
	//验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password wrong")
	}
	token, err := genToken(info.Id, req.Username)
	if err != nil {
		log.Error("service/CheckPassWord() genToken error is : ", err)
		return nil, errors.New("token generate err")
	}
	resp := &pb.CheckPassWordResponse{
		UserId: info.Id,
		Token:  token,
	}
	return resp, nil
}

// GetUserInfo 获取用户信息(优先从redis中获取,redis没有则从mysql中获取)
func (UserService) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	user, err := utils.GetUserInfoByAny(req.Id)
	if err != nil {
		log.Error("service/GetUserInfo() GetUser error is : ", err)
		return nil, err
	}
	resp := &pb.UserInfoResponse{
		UserInfo: UserToPbUserInfo(user),
	}
	return resp, nil
}

// GetUserInfoList 获取用户信息列表
func (UserService) GetUserInfoList(ctx context.Context, req *pb.UserInfoListRequest) (*pb.UserInfoListResponse, error) {
	resp := &pb.UserInfoListResponse{}
	users, err := utils.GetUserList(req.IdList)
	if err != nil {
		log.Error("service/GetUserInfoList() GetUserList error is : ", err)
		return nil, err
	}
	for _, user := range users {
		resp.UserInfoList = append(resp.UserInfoList, UserToPbUserInfo(*user))
	}
	return resp, nil
}

// Register 注册
func (UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	isExist, err := utils.IsUserNameExist(req.Username)
	if err != nil {
		log.Error("service/Register() IsUserNameExist error is : ", err)
		return nil, err
	} else if isExist {
		return nil, errors.New("username exist")
	}
	user, err := utils.CreateUserToMysqlAndCache(req.Username, req.Password)
	fmt.Println("username", req.Username, "  password: ", req.Password, "#######################")
	if err != nil {
		log.Error("service/Register() CreateUserToMysqlAndCache error is : ", err)
		return nil, err
	}
	token, err := genToken(user.Id, user.Name)
	if err != nil {
		return nil, errors.New("token generate err")
	}
	resp := &pb.RegisterResponse{
		UserId: user.Id,
		Token:  token,
	}
	return resp, nil
}

// GetUserInfoDict 获取用户信息字典(感觉这玩意没啥用)
func (UserService) GetUserInfoDict(ctx context.Context, req *pb.UserInfoDictRequest) (*pb.UserInfoDictResponse, error) {
	userslist, err := utils.GetUserList(req.IdList)
	if err != nil {
		log.Error("service/GetUserInfoDict() GetUserList error is : ", err)
		return nil, err
	}
	resp := &pb.UserInfoDictResponse{UserInfoDict: make(map[int64]*pb.UserInfo)}
	for _, user := range userslist {
		resp.UserInfoDict[user.Id] = &pb.UserInfo{
			Id:              user.Id,
			Name:            user.Name,
			Avatar:          user.Avatar,
			FollowCount:     user.Follow,
			FollowerCount:   user.Follower,
			Background:      user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavourited: user.TotalFav,
			FavouriteCount:  user.FavCount,
		}
	}
	return resp, nil
}

// CacheChangeUserCount 修改用户信息
// 添加了resp的信息
func (UserService) CacheChangeUserCount(ctx context.Context, req *pb.CacheChangeUserCountRequest) (*pb.CacheChangeUserCountResponse, error) {
	resp := &pb.CacheChangeUserCountResponse{CommonResponse: &pb.CommonResponse{}}
	uid := strconv.FormatInt(req.UserId, 10)
	mutex := lock.GetLock("user_" + uid)
	defer lock.UnLock(mutex)
	user, err := utils.CacheGetUserInfoById(req.UserId)
	if err != nil {
		log.Error("service/CacheChangeUserCount() CacheGetUserInfoById error is : ", err)
		resp.CommonResponse.Code = config.ErrorCode
		resp.CommonResponse.Msg = config.ErrorMsg
		return resp, err
	}
	switch req.CountType {
	case "follow":
		user.Follow += req.Op
	case "follower":
		user.Follower += req.Op
	case "like":
		user.FavCount += req.Op
	case "liked":
		user.TotalFav += req.Op
	}
	utils.CacheSetUser(user)
	resp.CommonResponse.Code = config.SuccessCode
	resp.CommonResponse.Msg = config.SuccessMsg
	return resp, nil
}

// UpdateMyTotalFavouritedCount 更新我的获赞数
// ActionType 1：表示赞数+1
//
//	2：赞数-1
func (UserService) UpdateMyTotalFavouritedCount(ctx context.Context, req *pb.UpdateMyFavouritedCountRequest) (*pb.UpdateMyFavouritedCountResponse, error) {
	resp := &pb.UpdateMyFavouritedCountResponse{CommonResponse: &pb.CommonResponse{}}
	err := utils.UpdateMyTotalFavouritedCount(req.UserId, req.ActionType)
	if err != nil {
		log.Error("service/UpdateMyTotalFavouritedCount() UpdateMyTotalFavouritedCount error is : ", err)
		resp.CommonResponse.Code = config.ErrorCode
		resp.CommonResponse.Msg = config.ErrorMsg
		return resp, err
	}
	resp.CommonResponse.Code = config.SuccessCode
	resp.CommonResponse.Msg = config.SuccessMsg
	return resp, nil
}

// UpdateMyFavouriteCount 更新我喜欢的视频数
// ActionType 1：表示喜欢数+1
// 2：喜欢数-1
func (UserService) UpdateMyFavouriteCount(ctx context.Context, req *pb.UpdateMyFavouriteCountRequest) (*pb.UpdateMyFavouriteCountResponse, error) {
	resp := &pb.UpdateMyFavouriteCountResponse{CommonResponse: &pb.CommonResponse{}}
	err := utils.UpdateMyFavouriteVideoCount(req.UserId, req.ActionType)
	if err != nil {
		log.Error("service/UpdateMyFavouriteCount() UpdateMyFavouriteVideoCount error is : ", err)
		resp.CommonResponse.Code = config.ErrorCode
		resp.CommonResponse.Msg = config.ErrorMsg
		return resp, err
	}
	resp.CommonResponse.Code = config.SuccessCode
	resp.CommonResponse.Msg = config.SuccessMsg
	return resp, nil
}

// UpdateMyFollowCount 更新我关注的数量
func (UserService) UpdateMyFollowCount(ctx context.Context, req *pb.UpdateMyFollowCountRequest) (*pb.UpdateMyFollowCountResponse, error) {
	resp := &pb.UpdateMyFollowCountResponse{CommonResponse: &pb.CommonResponse{}}
	err := utils.UpdateMyFollowCount(req.UserId, req.ActionType)
	if err != nil {
		log.Error("service/UpdateMyFollowCount() UpdateMyFollowCount error is : ", err)
		resp.CommonResponse.Code = config.ErrorCode
		resp.CommonResponse.Msg = config.ErrorMsg
		return resp, err
	}
	resp.CommonResponse.Code = config.SuccessCode
	resp.CommonResponse.Msg = config.SuccessMsg
	return resp, nil
}

func (UserService) UpdateMyFollowerCount(ctx context.Context, req *pb.UpdateMyFollowerCountRequest) (*pb.UpdateMyFollowerCountResponse, error) {
	resp := &pb.UpdateMyFollowerCountResponse{CommonResponse: &pb.CommonResponse{}}
	err := utils.UpdateMyFollowerCount(req.UserId, req.ActionType)
	if err != nil {
		log.Error("service/UpdateMyFollowerCount() UpdateMyFollowerCount error is : ", err)
		resp.CommonResponse.Code = config.ErrorCode
		resp.CommonResponse.Msg = config.ErrorMsg
		return resp, err
	}
	resp.CommonResponse.Code = config.SuccessCode
	resp.CommonResponse.Msg = config.SuccessMsg
	return resp, nil
}

// 两个功能函数(感觉放在这里其实不大合适)

// UserToPbUserInfo  把自定义的repository.User{}结构体转换成pb.UserInfo{}的形式
func UserToPbUserInfo(info utils.User) *pb.UserInfo {
	return &pb.UserInfo{
		Id:              info.Id,
		Name:            info.Name,
		FollowCount:     info.Follow,
		FollowerCount:   info.Follower,
		IsFollow:        false,
		Avatar:          info.Avatar,
		Background:      info.BackgroundImage,
		Signature:       info.Signature,
		TotalFavourited: info.TotalFav,
		FavouriteCount:  info.FavCount,
	}
}

// 生成token
func genToken(id int64, name string) (string, error) {
	expireTime := &jwt.NumericDate{
		Time: time.Now().Add(TokenExpireDuration), //可用于设定token过期时间
	}
	claims := JWTClaims{
		UserId:   id,
		Username: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "taosu",
			ExpiresAt: expireTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
