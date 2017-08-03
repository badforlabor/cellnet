package cellnet

import (
	"reflect"

	"fmt"
	"github.com/badforlabor/cellnet/proto/newprotocol"
	"github.com/golang/protobuf/proto"
)

// 协议拆包封包
type DefaultProtocolProcessor struct {
}
func (p DefaultProtocolProcessor) NewProtocol(msgid uint32) (interface{}){
	return newprotocol.NewProtocol(msgid)
}
func (p DefaultProtocolProcessor) GetProtocolID(proto interface{}) uint32 {
	return newprotocol.GetProtocolID(proto)
}
var protocolProcessor newprotocol.BinaryProtocolProcessor = DefaultProtocolProcessor{}
func SetProtocoloProcessor(p newprotocol.BinaryProtocolProcessor){
	protocolProcessor = p
}


// 普通封包
type Packet struct {
	MsgID uint32 // 消息ID
	Data  []byte
	rawMsg newprotocol.BinaryProtocol
}

func (self Packet) ContextID() uint32 {
	return self.MsgID
}

// 消息到封包
func BuildPacket1(data interface{}) (*Packet, *MessageMeta) {

	msg := data.(proto.Message)

	rawdata, err := proto.Marshal(msg)

	if err != nil {
		log.Errorln(err)
	}

	meta := MessageMetaByName(MessageFullName(reflect.TypeOf(msg)))

	return &Packet{
		MsgID: uint32(meta.ID),
		Data:  rawdata,
	}, meta
}

// 封包到消息
func ParsePacket1(pkt *Packet, msgType reflect.Type) (interface{}, error) {
	// msgType 为ptr类型, new时需要非ptr型

	rawMsg := reflect.New(msgType.Elem()).Interface()

	err := proto.Unmarshal(pkt.Data, rawMsg.(proto.Message))

	if err != nil {
		return nil, err
	}

	return rawMsg, nil
}

func BuildPacket(data interface{}) (*Packet, *MessageMeta) {

	pkg := &Packet{}

	defer func() {
		r := recover()
		if r != nil {
			pkg.MsgID = 0
			pkg.Data = nil
		}
	}()

	if _, ok := data.(newprotocol.BinaryProtocol); ok {
		pkg.MsgID = protocolProcessor.GetProtocolID(data)
		buffer := newprotocol.NewBinaryBuffer(nil)
		(data.(newprotocol.BinaryProtocol)).WriteMsg(buffer)
		pkg.Data = buffer.GetBytes()
	} else {
		fmt.Println("unsupport msg:", data)
	}

	return pkg, nil
}
func ParsePacket(pkt *Packet) (interface{}, error) {

	return pkt.rawMsg, nil
}

func (self *Packet)PreParsePacket() bool {

	var err error = nil

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			self.rawMsg = nil
		}
		self.Data = nil
	}()

	rawMsg := protocolProcessor.NewProtocol(self.MsgID)
	if rawMsg != nil {
		buffer := newprotocol.NewBinaryBuffer(self.Data)
		rawMsg.(newprotocol.BinaryProtocol).ReadMsg(buffer)
	}
	self.rawMsg = rawMsg.(newprotocol.BinaryProtocol)

	return self.rawMsg != nil
}
