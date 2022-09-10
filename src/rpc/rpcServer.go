package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
	rpc服务端创建流程
	1. 注册rpc服务 rpc.RegisterName("服务名",结构体指针)
	2. 创建一个监听
	3. 等待用户连接，返回一个conn
	4. 使用rpc服务连接对象
**/

/*
	注册rpc服务的结构体中暴露的方法必须满足下面条件
	1. 函数必须包外可见，即函数大写首字母
	2. 函数必须有两个参数，第一个单数为传入参数，第二个参数为传出参数，并且第二个参数必须为指针类型
	3. 函数必须返回一个错误，没有错误返回nil
**/

type Greeting struct{} // 结构体必须暴露出去
// HelloWord 暴露的方法必须满足条件
func (*Greeting) HelloWord(name string, res *string) error {
	*res = "hello " + name
	return nil
}

func rpcServer() {
	// 注册rpc服务
	err := rpc.RegisterName("greeter", &Greeting{})
	if err != nil {
		fmt.Println("注册服务失败")
	}
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
