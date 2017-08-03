package main

// 协议拆包封包
type myProtocolProcessor struct {
}
func (p myProtocolProcessor) NewProtocol(msgid uint32) (interface{}){
	return NewProtocol(msgid)
}
func (p myProtocolProcessor) GetProtocolID(proto interface{}) uint32 {
	return GetProtocolID(proto)
}
