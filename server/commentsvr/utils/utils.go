package utils

import (
	"commentsvr/config"
	"commentsvr/log"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/taosu0216/DouFlick/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: 没看懂这里是干什么的,回来再看看`

var (
	UserSvrClient pb.UserServiceClient
)

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=5s", consulInfo.Host, consulInfo.Port, svrName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		// grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	return conn, nil
}

func NewUserSvrClient(svrName string) pb.UserServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewUserServiceClient(conn)
}

func InitSvrConn() {
	UserSvrClient = NewUserSvrClient(config.GetGlobalConfig().CommentSvr.UserSvrName)
}

func GetUserSvrClient() pb.UserServiceClient {
	return UserSvrClient
}
