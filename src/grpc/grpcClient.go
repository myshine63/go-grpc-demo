package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-micro-demo/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// 通过服务发现，去consul中查询注册的服务
func getServiceHost() string {
	// 1.创建consul配置项
	consulConfig := api.DefaultConfig()
	// 2.使用配置项创建一个consul client
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("创建consul client 失败", err)
	}
	// 3.通过consul获取指定的注册服务的健康服务，返回指定服务的健康服务切片
	// 需要传入：服务名，服务的tag，是否通过健康检查（一般都传true），查询条件.
	healthServices, _, _ := consulClient.Health().Service("user-server", "user", true, &api.QueryOptions{})
	// 4.通过第一个健康的服务，获取注册服务的ip和端口
	address := healthServices[0].Service.Address
	ipPort := healthServices[0].Service.Port
	host := fmt.Sprint(address, ":", ipPort)
	fmt.Printf("服务的host为：%s\n", host)
	return host
}

// 直接使用grpc可以调用服务
func grpcClient(host string) {
	// 5.通过grpc连接服务
	grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("连接失败", err)
	}
	// 6.通过service创建一个客户端
	client := service.NewUserServerClient(grpcConn)
	user := &service.User{}
	// 7.通过客户端去调用远程方法
	res, _ := client.InitUser(context.TODO(), user)
	fmt.Println(res)
}

func main() {
	host := getServiceHost()
	grpcClient(host)
}
