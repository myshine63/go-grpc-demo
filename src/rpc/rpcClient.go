package main

import (
	"fmt"
	"net/rpc"
)

/*
	rpc客户端创建流程
	1. 用rpc.Dial访问服务端
	2. 使用Call("服务名.方法名",参数,接收返回值指针)方法去调用服务
**/

func rpcClient() {
	conn, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接失败！", err)
	}
	var res string
	err = conn.Call("greeter.HelloWord", "tom", &res)
	if err != nil {
		fmt.Println("调用服务方法失败", err)
	}
	fmt.Println(res)
}
func main() {
	rpcClient()
}
