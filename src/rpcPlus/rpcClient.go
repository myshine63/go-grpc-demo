package main

import (
	"fmt"
	"go-micro-demo/src/rpcPlus/pb"
)

/*
	rpc客户端创建流程
	1. 用rpc.Dial访问服务端
	2. 使用Call("服务名.方法名",参数,接收返回值指针)方法去调用服务
**/

func rpcClient() {
	client, err := pb.InitClient("localhost:8080")
	if err != nil {
		fmt.Println("连接失败！", err)
	}
	var res string
	err = client.HelloWord("tom", &res)
	if err != nil {
		fmt.Println("调用服务方法失败", err)
	}
	fmt.Println(res)
}
func main() {
	rpcClient()
}
