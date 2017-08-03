package socket

import (
	"fmt"

	"github.com/badforlabor/cellnet"
	_ "github.com/badforlabor/cellnet/proto/gamedef"
	"math"
)

var (
	Event_SessionConnected     = uint32(math.MaxUint32 - 1)
	Event_SessionClosed        = uint32(math.MaxUint32 - 2)
	Event_SessionAccepted      = uint32(math.MaxUint32 - 3)
	Event_SessionAcceptFailed  = uint32(math.MaxUint32 - 4)
	Event_SessionConnectFailed = uint32(math.MaxUint32 - 5)
)

// 会话事件
type SessionEvent struct {
	*cellnet.Packet
	Ses cellnet.Session
}

func (self SessionEvent) String() string {
	return fmt.Sprintf("SessionEvent msgid: %d data: %v", self.MsgID, self.Data)
}

func NewSessionEvent(msgid uint32, s cellnet.Session, data []byte) *SessionEvent {
	return &SessionEvent{
		Packet: &cellnet.Packet{MsgID: msgid, Data: data},
		Ses:    s,
	}
}

func newSessionEvent(msgid uint32, s cellnet.Session, msg interface{}) *SessionEvent {

	pkt, _ := cellnet.BuildPacket1(msg)

	return &SessionEvent{
		Packet: pkt,
		Ses:    s,
	}

}
