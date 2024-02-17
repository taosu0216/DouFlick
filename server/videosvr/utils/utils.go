package utils

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os/exec"
	"path/filepath"
	"strings"
	"videosvr/config"
	"videosvr/log"
)

var (
	FavoriteSvrClient pb.FavoriteServiceClient
	RelationSvrClient pb.RelationServiceClient
)

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//TODO:链路追踪未完成
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	return conn, nil
}

func NewRelationSvrClient(name string) pb.RelationServiceClient {
	conn, err := NewSvrConn(name)
	if err != nil {
		log.Errorf("NewRelationSvrClient err:%v", err)
	}
	return pb.NewRelationServiceClient(conn)
}

func NewFavoriteSvrClient(name string) pb.FavoriteServiceClient {
	conn, err := NewSvrConn(name)
	if err != nil {
		log.Errorf("NewFavoriteSvrClient err:%v", err)
	}
	return pb.NewFavoriteServiceClient(conn)
}
func InitSvrConn() {
	RelationSvrClient = NewRelationSvrClient(config.GetGlobalConfig().SvrConfig.RelationSvrName)
	FavoriteSvrClient = NewFavoriteSvrClient(config.GetGlobalConfig().SvrConfig.FavoriteSvrName)
}
func GetRelationSvrClient() pb.RelationServiceClient {
	return RelationSvrClient
}

func GetFavoriteSvrClient() pb.FavoriteServiceClient {
	return FavoriteSvrClient
}

func GetImage(savefile string) (string, error) {
	tmp := strings.Split(savefile, "/")
	videoName := tmp[len(tmp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-3]) + "jpg"
	picPath := config.GetGlobalConfig().MinioConfig.PicPath
	picName := filepath.Join(picPath, videoName)
	cmd := exec.Command("ffmpeg", "-i", savefile, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picName)
	err := cmd.Run()
	if err != nil {
		log.Errorf("cmd.Run() failed with %s\n", err)
		return "", err
	}
	return picName, nil
}
