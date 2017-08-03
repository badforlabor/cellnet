/*
    此文件为自动生成的。
*/

/*
    语法规则：
        注释：使用C格式的
        关键字package
        关键字message定义协议
        变量类型：int,float,string,嵌套message
        支持数组，使用repeated关键字
*/
package main//gamedef

import 	"github.com/badforlabor/cellnet/proto/newprotocol"

/*
type BinaryProtocol interface {
    ReadMsg(buffer * BinaryBuffer)
    WriteMsg(buffer * BinaryBuffer)
}
*/

// 协议映射关系
const (

	PID_SessionAccepted = 1 + 0

	PID_SessionConnected = 1 + 1

	PID_TestEchoACK = 1 + 2

	PID_NestedMessage = 1 + 3

	PID_ArrayMessage = 1 + 4

)
func NewProtocol(msgid uint32) (interface{}) {
	switch(msgid) {

	case PID_SessionAccepted:
		return &SessionAccepted{}

	case PID_SessionConnected:
		return &SessionConnected{}

	case PID_TestEchoACK:
		return &TestEchoACK{}

	case PID_NestedMessage:
		return &NestedMessage{}

	case PID_ArrayMessage:
		return &ArrayMessage{}
	}
	return nil
}
func GetProtocolID(proto interface{}) uint32 {

	switch proto.(type) {

	case *SessionAccepted:
		return PID_SessionAccepted

	case *SessionConnected:
		return PID_SessionConnected

	case *TestEchoACK:
		return PID_TestEchoACK

	case *NestedMessage:
		return PID_NestedMessage

	case *ArrayMessage:
		return PID_ArrayMessage
	}

	return 0
}



// ==========================================================

// 系统消息

// ==========================================================

// 一个连接接入

type SessionAccepted struct {

}
func (msg *SessionAccepted) ReadMsg(buffer *newprotocol.BinaryBuffer) {

}
func (msg *SessionAccepted) WriteMsg(buffer *newprotocol.BinaryBuffer) {

}

// 已连接

type SessionConnected struct {

}
func (msg *SessionConnected) ReadMsg(buffer *newprotocol.BinaryBuffer) {

}
func (msg *SessionConnected) WriteMsg(buffer *newprotocol.BinaryBuffer) {

}

// ==========================================================

// 测试用消息

// ==========================================================

type TestEchoACK struct {
	Content string
	Content2 string

}
func (msg *TestEchoACK) ReadMsg(buffer *newprotocol.BinaryBuffer) {
	msg.Content = buffer.ReadString()
	msg.Content2 = buffer.ReadString()

}
func (msg *TestEchoACK) WriteMsg(buffer *newprotocol.BinaryBuffer) {
	buffer.WriteString(msg.Content)
	buffer.WriteString(msg.Content2)

}

type NestedMessage struct {
	Acks TestEchoACK // 嵌套语句
	id int32
	f float32
	str string

}
func (msg *NestedMessage) ReadMsg(buffer *newprotocol.BinaryBuffer) {
	msg.Acks.ReadMsg(buffer)

	msg.id = buffer.ReadInt()
	msg.f = buffer.ReadFloat()
	msg.str = buffer.ReadString()

}
func (msg *NestedMessage) WriteMsg(buffer *newprotocol.BinaryBuffer) {
	msg.Acks.WriteMsg(buffer)

	buffer.WriteInt(msg.id)
	buffer.WriteFloat(msg.f)
	buffer.WriteString(msg.str)

}

type ArrayMessage struct {
	datas []int32
	msgs []string
	p []NestedMessage
	id int32
	f float32
	str string

}
func (msg *ArrayMessage) ReadMsg(buffer *newprotocol.BinaryBuffer) {
	msg.datas = buffer.ReadIntArray()
	msg.msgs = buffer.ReadStringArray()

	{
		var size int32 = buffer.ReadInt()
		msg.p = make([]NestedMessage, size)
		for i:=int32(0); i < size; i++ {
			msg.p[i].ReadMsg(buffer)
		}
	}


	msg.id = buffer.ReadInt()
	msg.f = buffer.ReadFloat()
	msg.str = buffer.ReadString()

}
func (msg *ArrayMessage) WriteMsg(buffer *newprotocol.BinaryBuffer) {
	buffer.WriteIntArray(msg.datas)
	buffer.WriteStringArray(msg.msgs)

	{
		var size int32 = int32(len(msg.p))
		buffer.WriteInt(size)
		for i:=int32(0); i < size; i++ {
			msg.p[i].WriteMsg(buffer)
		}
	}


	buffer.WriteInt(msg.id)
	buffer.WriteFloat(msg.f)
	buffer.WriteString(msg.str)

}
