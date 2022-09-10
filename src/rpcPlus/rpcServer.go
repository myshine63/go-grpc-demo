package main

import (
	"fmt"
	"go-micro-demo/src/rpcPlus/pb"
	"net"
	"net/rpc"
)

/*
	rpc服务端封装流程
	1. 定义一个rpc服务接口，列出需要实现的方法
	2. 定义rpc服务注册方法，以rpc服务接口为参数
	3. 定义一个结构体，实现rpc服务接口
	4. 使用rpc注册方法， 注册rpc服务
**/

type Greeting struct{} // 定义结构体，实现rpc服务接口

func (*Greeting) HelloWord(name string, res *string) error {
	*res = "hello " + name
	return nil
}

func rpcServer() {
	// 注册rpc服务
	pb.RegisterService(&Greeting{})
	// 启动监听
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("启动服务器失败！", err)
	}
	defer listener.Close()
	fmt.Println("启动服务成功，等待用户连接...")
	// 等待连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("连接失败", err)
	}
	defer conn.Close()
	// 使用rcp服务连接对象
	rpc.ServeConn(conn)
}

func main() {
	rpcServer()
}
