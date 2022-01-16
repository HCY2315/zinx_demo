package znet

import "zinx-demo/ziface"

/*
	实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写就好
	这里之所以BaseRoouter 的方法
*/
type BaseRouter struct{}

/*
	这里之所以 BaseRoouter 的方法为空
	是因为有的 Router 不希望有 PreHandle PostHandle 这两个业务
	所以 Router 全部继承 BaseRouter 的好处就是，不需要实现 PreHandle PostHandle
*/

// 在处理 conn 业务之前的钩子方法HOOK
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

// 在处理 conn 业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {}

// 在处理 conn 业务之后的钩子方法HOOK
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
