package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

/*
	封包、拆包 模块
	直接面向TCP连接中的数据流，用于处理TCP粘包问题
*/

// 拆包，拆包的主体模块
type DataPack struct{}

// 拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的头的长度
func (dp *DataPack) GetHandLen() uint32 {
	// dataLen uint32(4字节) + id uint32(4字节) = 8字节
	return 8
}

// 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将 dataLen 写进dataBuff中
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	// 将 MsgId 写进dataBuff中
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 将 data数据 写进dataBuff中
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法(将包的Head信息读出来，之后再根据head信息里的data的长度，在进行一次读)
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	// 只解压head信息，得到dataLen和MsgId
	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读MsgId
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断dataLen是否已经超出了我们允许的最大包长度
	if utils.GlobObject.MaxPackage > 0 && msg.DataLen > utils.GlobObject.MaxPackage {
		return nil, errors.New("too Large msg data recv !")
	}

	return msg, nil
}
