package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	// 1、直接连接远程服务器，得到一个 conn 连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err:", err)
		return
	}

	for {
		// 2、连接调用 write 写数据
		_, err := conn.Write([]byte("hello zinx v1.0"))
		if err != nil {
			fmt.Println("write conn err:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}

		fmt.Printf("server call back: %s, cut = %d\n", buf, cnt)

		// cpu 堵塞
		time.Sleep(11 * time.Second)
	}
}
