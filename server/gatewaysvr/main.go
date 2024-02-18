package main

import (
	"fmt"
	"gatewaysvr/config"
	"gatewaysvr/log"
	"gatewaysvr/routes"
	"gatewaysvr/utils"
)

func init() {
	err := config.Init()
	if err != nil {
		log.Fatalf("config.Init() error:%v", err)
	}
	log.InitLog()
	log.Info("log init success...")
	utils.InitSvrConn()
	log.Info("Init Grpc Server conn success")
}

func main() {
	defer log.Sync()
	r := routes.RouteInit()
	dsn := fmt.Sprintf(":%d", config.GetGlobalConfig().SvrConfig.Port)
	if err := r.Run(dsn); err != nil {
		panic(err)
	}
}
