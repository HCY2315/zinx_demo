package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx-demo/ziface"
)

/*
	存储一切有关 Zinx 框架的全局参数，供其他模块使用
	一些参数是可以通过 Zinx.json 有用户进行配置
*/

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer // 当前 Zinx 全局的Server对象
	Host      string         // 当前主机监听的IP
	TcpPort   int            // 当前服务器监听的端口
	Name      string         // 当前服务器的名称

	/*
		Zinx
	*/
	Version    string // 当前Zinx 的版本号
	MaxConn    int    //当前服务器主机允许最大的连接数
	MaxPackage uint32 // 当前Zinx框架数据包的最大值
}

/*
	定义一个全局的对外GlobalObj
*/

var GlobObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("/Users/houchaoyue/Documents/demo/golang/zinx-demo-03/mydemo/zinx01/conf/zinx.json")
	if err != nil {
		fmt.Println("读取文件失败", err)
	}

	// 将 json 文件的数据解析到 struck 中
	err = json.Unmarshal(data, &GlobObject)
	if err != nil {
		panic(err)
	}
}

/*
	初始化当前的GlobalObject
*/
func init() {
	// 如果配置文件没有加载，使用一下默认值
	GlobObject = &GlobalObj{
		Name:       "ZinxServerApp",
		Version:    "latest",
		TcpPort:    8999,
		Host:       "0.0.0.0",
		MaxConn:    1000,
		MaxPackage: 4096,
	}

	// 应该尝试从conf/zinx.json 中加载一些用户自动移的一些参数
	GlobObject.Reload()
}
