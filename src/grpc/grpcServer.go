package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-micro-demo/src/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
	grpc服务端实现流程
	1. 创建一个结构体，实现service中server接口
	2. 使用结构体注册服务
	3. 启动监听
	4. 使用grpc服务，传入监听对象
**/

// 创建结构体，去实现pb中定义的用户服务接口
type userServer struct {
	service.UnimplementedUserServerServer
}

func (u userServer) InitUser(ctx context.Context, user *service.User) (*service.User, error) {
	user.Name = "tom"
	user.Age = 15
	user.AddressInfo = &service.AddressInfo{
		Address: "cheng-du",
		Email:   "qq.com",
	}
	user.Gender = service.Gender_female
	user.Score = 100.0
	return user, nil
}

// 将服务注册到consul
func registerGrpcToConsul() {
	// 1.创建一个consul配置项，这里采用默认配置
	consulConfig := api.DefaultConfig()
	// 2.创建一个consul客户端
	consulClient, _ := api.NewClient(consulConfig)
	// 3.创建一个服务的配置信息
	serverRegisterConfig := api.AgentServiceRegistration{
		ID:      "demo1",
		Name:    "user-server",
		Tags:    []string{"user", "user-server"},
		Port:    8080,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8080",
			Interval: "5s",
			Timeout:  "2s",
		},
	}
	// 将service的配置注册到consul的client上
	consulClient.Agent().ServiceRegister(&serverRegisterConfig)
}

// 创建grpc服务器流程
func createGrpcServer() {
	// 1. 创建grpc服务器
	grpcServer := grpc.NewServer()
	// 2. 注册一个服务
	service.RegisterUserServerServer(grpcServer, &userServer{})
	// 3. 启动一个tcp监听
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("启动服务...")
	defer listener.Close()
	// 4. 将监听对象传入grpc服务
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	registerGrpcToConsul()
	createGrpcServer()
}
