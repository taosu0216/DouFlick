package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
	"usersvr/config"
	"usersvr/log"
	"usersvr/middleware/cache"
	"usersvr/middleware/consul"
	"usersvr/middleware/db"
	"usersvr/middleware/lock"
	"usersvr/service"
)

func init() {
	//初始化配置文件
	err := config.Init()
	if err != nil {
		log.Fatalf("config.Init() error:%v", err)
	}

	//初始化日志
	log.InitLog()

	//初始化mysql
	db.DBinit()

}
func run() error {
	//TODO: consul没看懂
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "", config.GetGlobalConfig().UserSvr.Port))
	if err != nil {
		log.Fatalf("net.Listen error:%v", err)
		return err
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &service.UserService{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	consulClient := consul.NewRegistryClient(config.GetGlobalConfig().ConsulConfig.Host, config.GetGlobalConfig().ConsulConfig.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	if err = consulClient.Register(config.GetGlobalConfig().UserSvr.Host, config.GetGlobalConfig().UserSvr.Port, config.GetGlobalConfig().UserSvr.Name, config.GetGlobalConfig().ConsulConfig.Tags, serviceId); err != nil {
		log.Fatalf("consulClient.Register error:%v", err)
		return err
	}
	log.Infof("Init Consul Register success")

	log.Infof("DouFlick.relation_svr listening on %s:%d", config.GetGlobalConfig().UserSvr.Host, config.GetGlobalConfig().UserSvr.Port)
	go func() {
		err = server.Serve(listen)
		if err != nil {
			log.Fatalf("server.Serve error:%v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
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
	defer lock.CloseLock()
	if err := run(); err != nil {
		log.Fatalf("run error:%v", err)
	}
}
