package main

import (
	"fmt"
	"zinx-demo/ziface"
	"zinx-demo/znet"
)

/*
	基于Zinx框架开发的 服务器端应用程序
*/

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// test ProHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back brfore ping error")
	}
}

// test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	if err != nil {
		fmt.Println("call back ping... error")
	}
}

// test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router postHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping... error")
	}
}

func main() {
	// 1、创建server据柄，使用Zinx的api
	s := znet.NewServer()

	// 2、给当前 zinx 框架添加一个自定义的 router
	s.AddRouter(&PingRouter{})

	// 3、启动 server
	s.Serve()
}
