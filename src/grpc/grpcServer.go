package main

import (
	"context"
	"fmt"
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

func main() {
	grpcServer := grpc.NewServer()
	service.RegisterUserServerServer(grpcServer, &userServer{})
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	fmt.Println("启动服务，等待用户连接。。。")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
