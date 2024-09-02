package main

import (
	"fmt"
)

// ServiceRegistrar interface    
type ServiceRegistrar interface {
	RegisterService(serviceName string, impl any)
}

// GreeterServer struct    
type GreeterServer struct{}

type gRPCserver struct{}

func (m *gRPCserver) RegisterService(serviceName string, impl any) {
	fmt.Println("Registering service", serviceName)
}

// RegisterGreeterServer 函数接受一个 ServiceRegistrar 和一个 GreeterServer
func RegisterGreeterServer(s ServiceRegistrar, srv *GreeterServer) {
	// 使用传入的 ServiceRegistrar 注册 GreeterServer
	s.RegisterService("GreeterService", srv)
}

func main() {
	// 创建自定义的 ServiceRegistrar 实例
	myServer := &gRPCserver{}

	// 创建 GreeterServer 实例
	myGreeterServer := &GreeterServer{}

	// 调用 RegisterGreeterServer 进行服务注册
	RegisterGreeterServer(myServer, myGreeterServer)
}
