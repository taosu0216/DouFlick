package main

import (
	"favoritesvr/config"
	"favoritesvr/init/cache"
	"favoritesvr/init/consul"
	"favoritesvr/init/db"
	"favoritesvr/log"
	"favoritesvr/service"
	"favoritesvr/utils"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"

	"github.com/taosu0216/DouFlick/pkg/pb"
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
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GetGlobalConfig().SvrConfig.Host, config.GetGlobalConfig().SvrConfig.Port))
	if err != nil {
		log.Fatalf("listen: error %v", err)
		return fmt.Errorf("listen: error %v", err)
	}
	// 端口监听启动成功，启动grpc server
	server := grpc.NewServer()
	// 注册grpc server
	pb.RegisterFavoriteServiceServer(server, &service.FavoriteService{}) // 注册服务
	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//注册到consul
	consulClient := consul.NewRegistryClient(config.GetGlobalConfig().ConsulConfig.Host, config.GetGlobalConfig().ConsulConfig.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	if err := consulClient.Register(config.GetGlobalConfig().SvrConfig.Host, config.GetGlobalConfig().SvrConfig.Port, config.GetGlobalConfig().SvrConfig.Name, config.GetGlobalConfig().ConsulConfig.Tags, serviceId); err != nil {
		log.Fatalf("consul register error:%v", err)
		return fmt.Errorf("consul register error:%v", err)
	}
	log.Infof("consul register success, serviceId:%s", serviceId)

	// 启动服务
	log.Infof("DouFlick.favorite_svr listening on %s:%d", config.GetGlobalConfig().SvrConfig.Host, config.GetGlobalConfig().SvrConfig.Port)
	go func() {
		if err := server.Serve(listen); err != nil {
			panic(err)
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务终止，注销 consul 服务
	if err = consulClient.DeRegister(serviceId); err != nil {
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
