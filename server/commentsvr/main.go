package main

import (
	"commentsvr/config"
	"commentsvr/init/cache"
	"commentsvr/init/consul"
	"commentsvr/init/db"
	"commentsvr/log"
	"commentsvr/service"
	"commentsvr/utils"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	uuid "github.com/satori/go.uuid"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func init() {
	//初始化配置文件
	err := config.Init()
	if err != nil {
		log.Fatalf("config.Init() error:%v", err)
	}

	//初始化日志
	log.InitLog()
	log.Info("Init log success")

	//初始化mysql
	db.DBinit()
	log.Info("Init mysql success")

	//初始化redis
	cache.RedisInit()
	log.Info("Init redis success")

	//初始化user rpc模块通信
	utils.InitSvrConn()
	log.Info("Init UserModel RPC service success")
}

func Run() error {
	listen, err := net.Listen("tcp", "0.0.0.0:8002")
	if err != nil {
		log.Fatalf("listen: error %v", err)
		return fmt.Errorf("listen: error %v", err)
	}
	// 端口监听启动成功，启动grpc server
	server := grpc.NewServer()
	// 注册grpc server
	pb.RegisterCommentServiceServer(server, &service.CommentService{}) // 注册服务
	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 注册服务到consul中
	consulClient := consul.NewRegistryClient(config.GetGlobalConfig().ConsulConfig.Host, config.GetGlobalConfig().ConsulConfig.Port)
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	port, _ := strconv.Atoi(config.GetGlobalConfig().CommentSvr.Port)
	if err := consulClient.Register(config.GetGlobalConfig().CommentSvr.Host, port,
		config.GetGlobalConfig().Name, config.GetGlobalConfig().ConsulConfig.Tags, serviceID); err != nil {
		log.Fatal("consul.Register error: ", zap.Error(err))
		return errors.New("consul.Register error: " + err.Error())
	}

	//log.Info("Init Consul Register success")

	// 启动
	log.Infof("DouFlick.comment_svr listening on %s:%d", config.GetGlobalConfig().
		CommentSvr.Host, port)
	err = server.Serve(listen)
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务终止，注销 consul 服务
	if err = consulClient.DeRegister(serviceID); err != nil {
		log.Info("注销失败")
		return fmt.Errorf("注销失败")
	} else {
		log.Info("注销成功")
	}
	return nil
}

func main() {
	defer log.Sync()
	defer db.CloseDB()
	defer cache.CloseRedis()

	if err := Run(); err != nil {
		log.Fatalf("run error:%v", err)
	}
}
