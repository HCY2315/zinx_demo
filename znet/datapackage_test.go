package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//负责测试 dataPackage拆包、封包的单元测试
func TestDataPack(t *testing.T) {
	/*
		模拟客户端
	*/
	// 创建socketTCP
	Listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}

	// 从客户端读取
	go func() {
		for {
			conn, err := Listen.Accept()
			if err != nil {
				fmt.Println("server accept error:", err)
				return
			}
			go func(conn net.Conn) {
				// 处理客户端的请求
				//-----> 拆包的过程 <-----
				// 定义一个拆包的对象
				dp := NewDataPack()
				for {
					// 1、第一次从 conn 读，把包的 head 读出来；
					headData := make([]byte, dp.GetHandLen())
					if _, err := io.ReadFull(conn, headData); err != nil {
						fmt.Println("read head error", err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error:", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						// msg 是有数据的，需要进行第二次读取
						// 第二次从 conn 读，根据 head 中的 datalen 再读取 data 内容；
						msg.Data = make([]byte, msg.GetMsgLen())
						if _, err := io.ReadFull(conn, msg.Data); err != nil {
							fmt.Println("server unpack error:", err)
							return
						}
						// 完整的消息已经读取完成
						fmt.Println("-----> Recv MsgID:", msg.Id, "datalen:", msg.DataLen, "data: ", msg.Data)
					}
				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}

	// 创建一个封包对象 dp
	dp := NewDataPack()

	// 模拟粘包，封装两个msg一同发送
	// 封装第一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 3,
		Data:    []byte{'1', '2', '3'},
	}
	sendPack1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
	}

	// 封装第一个msg12包
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'n', 'i', 'h', 'a', '0'},
	}
	sendPack2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
	}

	// 将两个包粘在一起
	sendPack1 = append(sendPack1, sendPack2...)

	// 一次性发送给服务端
	conn.Write(sendPack1)

	// 客户端阻塞
	select {}
}
