package ziface

/*
	路由抽象接口
	路由里的数据都是IRequest
*/

type IRouter interface {
	// 在处理 conn 业务之前的钩子方法HOOK
	PreHandle(request IRouter)

	// 在处理 conn 业务的主方法
	Handle(request IRouter)

	// 在处理 conn 业务之后的钩子方法HOOK
	PostHandle(request IRouter)
}