package consul

import (
	"fmt"
	"relationsvr/log"

	"github.com/hashicorp/consul/api"
)

// Registry 是一个结构体，包含注册中心的主机和端口信息
type Registry struct {
	Host string // 注册中心的主机地址
	Port int    // 注册中心的端口号
}

// RegistryClient 是一个接口，定义了注册和注销服务的方法
type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error // 注册服务的方法
	DeRegister(serviceId string) error                                              // 注销服务的方法
}

// NewRegistryClient 是一个函数，用于创建新的注册中心客户端
func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{ // 返回一个Registry结构体的实例
		Host: host, // 设置主机地址
		Port: port, // 设置端口号
	}
}

// Register 是Registry结构体的一个方法，用于向注册中心注册服务
func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()                         // 获取默认的配置
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port) // 设置配置的地址
	client, err := api.NewClient(cfg)                  // 创建新的客户端
	if err != nil {
		log.Fatalf("api.NewClient error:%v", err) // 如果创建客户端出错，记录错误并退出
	}

	// 创建新的服务检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port), // 设置GRPC的地址
		Timeout:                        "30s",                               // 设置超时时间
		Interval:                       "6s",                                // 设置检查间隔
		DeregisterCriticalServiceAfter: "15s",                               // 如果服务处于严重状态超过此时间，自动取消注册
	}

	// 创建新的服务注册对象
	reg := &api.AgentServiceRegistration{
		ID:      id,      // 设置服务ID
		Name:    name,    // 设置服务名称
		Port:    port,    // 设置服务端口
		Tags:    tags,    // 设置服务标签
		Address: address, // 设置服务地址
		Check:   check,   // 设置服务检查对象
	}
	err = client.Agent().ServiceRegister(reg) // 注册服务
	if err != nil {
		log.Fatalf("client.Agent().ServiceRegister error:%v", err) // 如果注册服务出错，记录错误并退出
	}
	return nil // 如果没有错误，返回nil
}

// DeRegister 是Registry结构体的一个方法，用于从注册中心注销服务
func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()                         // 获取默认的配置
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port) // 设置配置的地址

	client, err := api.NewClient(cfg) // 创建新的客户端
	if err != nil {
		return err // 如果创建客户端出错，返回错误
	}
	err = client.Agent().ServiceDeregister(serviceId) // 注销服务
	return err                                        // 返回错误，如果没有错误，返回nil
}
