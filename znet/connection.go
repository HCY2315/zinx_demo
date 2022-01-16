package znet

import (
	"fmt"
	"net"
	"zinx-demo/ziface"
)

/*
	连接模块
*/
type Connection struct {
	// 当前连接的socket TCP 套接字
	Conn *net.TCPConn

	// 连接ID
	ConnID uint32

	// 当前连接状态
	IsClosed bool

	// 告知当前连接已经退出的/停止 channel
	ExitChan chan bool

	// 该连接处理的方法
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 连接的读业务的方法
func (c *Connection) StartReader() {
	fmt.Println(" Reader Goroutine is running...")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remote addr is", c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 得到当前 conn 数据的 Request 请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// // 调用当前连接所绑定的HandlerAPI
		// if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("ConnID", c.ConnID, "handle is error", err)
		// 	break
		// }
	}
}

// 启动连接，让当前的连接准备工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	// 启动从当前连接的读数据的连接
	go c.StartReader()
	// TODO 启动当前连接写数据的业务

}

// 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop() ... ConnID = ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true

	// 调用 Socker 连接
	c.Conn.Close()

	close(c.ExitChan)
}

// 获取当前连接的绑定 socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 TCP 状态（ip\port）
func (c *Connection) GetRemoveAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送到远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
