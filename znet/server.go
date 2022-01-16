/*
实例层
*/

package znet

import (
	"fmt"
	"net"
	"zinx-demo/ziface"
)

// IServer 的接口实现，定义一个server的服务模块
type Server struct {
	// 服务器名称
	Name string

	// 服务器绑定IP地址
	IPVersion string

	// 服务器监听端口号
	IP string

	// 服务器监听的端口
	Port int

	// 当前的 server 添加一个 router,server 注册的连接对应的处理业务
	Router ziface.IRouter
}

// // 定义当前客户端连接的所绑定handle Api(目前这个handler是写死的，以后优化应用，有用户来定义)
// func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
// 	// 回显的功能
// 	fmt.Println("[Conn Handler] CallBackToClient ...")
// 	if _, err := conn.Write(data[:cnt]); err != nil {
// 		fmt.Println("Read Buf Error: ", err)
// 		return errors.New("CallBack")
// 	}
// 	return nil
// }

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP : %s, Port: %d, is Starting\n", s.IP, s.Port)

	go func() {
		// 1、获取一个 TCP 的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}

		// 2、监听服务器的地址
		listerner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s, s.IPVersion, "err:", err)
			return
		}

		fmt.Println("start Zinx server succ", s.Name, "sunc, Listenning...")
		var cid uint32
		cid = 0

		// 3、阻塞等待的客户端连接，处理客户端连接业务
		for {
			// 如果有客户端连接，阻塞会返回
			conn, err := listerner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept, err:", err)
				continue
			}

			// 处理新连接的业务方法和 conn 进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前的连接业务处理
			go dealConn.Start()

			// // 已经与客户端建立连接， 做一些业务，做最基本的512字节长度的回显业务
			// go func() {
			// 	for {
			// 		buf := make([]byte, 512)
			// 		cnt, err := conn.Read(buf)
			// 		if err != nil {
			// 			fmt.Println("recv buf err", err)
			// 			continue
			// 		}

			// 		fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)

			// 		// 回显功能
			// 		if _, err := conn.Write(buf[:cnt]); err != nil {
			// 			fmt.Println("write back buf err", err)
			// 			continue
			// 		}
			// 	}
			// }()
		}
	}()

}

func (s *Server) Stop() {
	// TODO 将服务器的资源、状态或者一些已经开辟的连接信息，进行停止或者回收

}

func (s *Server) Serve() {
	s.Start()

	// TODO 做一些启动服务之后的额外业务

	// 阻塞状态
	select {}

}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!!")
}

/*
	初始化 server 的方法
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
