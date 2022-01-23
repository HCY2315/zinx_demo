package ziface

/*
	将请求的数据封装到一个 Message 中， 定义抽象接口
*/

type IMessage interface {
	// 获取消息的ID
	GetMsgId() uint32

	// 获取消息的长度
	GetMsgLen() uint32

	// 获取消息的内容
	GetMsgData() []byte

	// 设置消息的ID
	SetMsgId(uint32)

	// 设置消息的内容
	SetMsgData([]byte)

	// 设置消息的长度
	SetMsgDataLen(uint32)
}
