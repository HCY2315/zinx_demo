/*
抽象层
*/

package ziface

// 定义服务器接口
type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 运行服务器
	Serve()

	// 初始化服务器方法

}
