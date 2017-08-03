package socket

import "github.com/badforlabor/cellnet"

func MessageRegistedCount(evd cellnet.EventDispatcher, msgName string) int {

	msgMeta := cellnet.MessageMetaByName(msgName)
	if msgMeta == nil {
		return 0
	}

	return evd.CountByID(msgMeta.ID)
}

type RegisterMessageContext struct {
	*cellnet.MessageMeta
	*cellnet.CallbackContext
}

// [废弃] 注册连接消息
func RegisterMessage(evd cellnet.EventDispatcher, msgName string, userHandler func(interface{}, cellnet.Session)) *RegisterMessageContext {

	msgMeta := cellnet.MessageMetaByName(msgName)

	if msgMeta == nil {
		log.Errorf("message register failed, %s", msgName)
		return nil
	}

	ctx := evd.AddCallback(msgMeta.ID, func(data interface{}) {

		if ev, ok := data.(*SessionEvent); ok {

			rawMsg, err := cellnet.ParsePacket1(ev.Packet, msgMeta.Type)

			if err != nil {
				log.Errorf("unmarshaling error: %v, raw: %v", err, ev.Packet)
				return
			}

			userHandler(rawMsg, ev.Ses)

		}

	})

	return &RegisterMessageContext{MessageMeta: msgMeta, CallbackContext: ctx}
}
// 注册自定义协议
func NewRegisterMessage(evd cellnet.EventDispatcher, id uint32, userHandler func(interface{}, cellnet.Session))  {

	evd.AddCallback(id, func(data interface{}) {

		if ev, ok := data.(*SessionEvent); ok {

			rawMsg, err := cellnet.ParsePacket(ev.Packet)

			if err != nil {
				log.Errorf("unmarshaling error: %v, raw: %v", err, ev.Packet)
				return
			}

			userHandler(rawMsg, ev.Ses)

		}

	})
}
// 注册系统消息：比如Event_SessionConnected，Event_SessionClosed
func RegisterSysMessage(evd cellnet.EventDispatcher, id uint32, userHandler func(interface{}, cellnet.Session))  {

	evd.AddCallback(id, func(data interface{}) {

		if ev, ok := data.(*SessionEvent); ok {

			userHandler(nil, ev.Ses)

		}

	})
}

