package main

import "zinx-demo/znet"

/*
	基于Zinx框架开发的 服务器端应用程序
*/

func main() {
	// 1、创建server据柄，使用Zinx的api
	s := znet.NewServer("[zinx01]")

	// 2、启动 server
	s.Serve()

}
