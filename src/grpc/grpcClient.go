package main

import (
	"context"
	"fmt"
	"go-micro-demo/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	grpcConn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("连接失败", err)
	}
	client := service.NewUserServerClient(grpcConn)
	user := &service.User{}
	res, _ := client.InitUser(context.TODO(), user)
	fmt.Println(res)
}
