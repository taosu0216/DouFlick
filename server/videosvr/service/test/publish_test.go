package test

import (
	"context"
	"fmt"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"testing"
	"videosvr/config"
	"videosvr/init/db"
	"videosvr/log"
	"videosvr/service"
)

func TestPublishVideo(t *testing.T) {
	err := config.Init()
	if err != nil {
		t.Error(err)
	}
	log.InitLog()
	db.DBinit()
	req := &pb.PublishVideoRequest{UserId: 1, SaveFile: `D:\视频\QQ2024217-155637-HD.mp4`, Title: "test"}
	ctx := context.Background()
	resp, err := service.VideoService{}.PublishVideo(ctx, req)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
