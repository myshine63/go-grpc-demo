package pb

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
	rpc服务端封装流程
	1. 定义一个rpc服务接口，列出需要实现的方法
	2. 定义rpc服务注册方法，以rpc服务接口为参数
	3. 定义一个客户端，客户端实现所有方法
	4. 定义一个初始化客户端的方法，返回一个客户端，通过这个客户端就可以调用所有的方法
**/

type GreetingServer interface {
	HelloWord(string, *string) error
}

// RegisterService 注册rpc服务
func RegisterService(i GreetingServer) {
	err := rpc.RegisterName("greeting", i)
	if err != nil {
		fmt.Println("注册服务失败")
	}
}

type GreetingClient struct {
	client *rpc.Client
}

func InitClient(addr string) (*GreetingClient, error) {
	conn, err := jsonrpc.Dial("tcp", addr)
	return &GreetingClient{client: conn}, err
}

func (g *GreetingClient) HelloWord(s string, s2 *string) error {
	return g.client.Call("greeting.HelloWord", s, s2)
}
