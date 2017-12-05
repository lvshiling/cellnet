package rpc

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

type RecvMsgEvent struct {
	ses    cellnet.Session
	Msg    interface{}
	callid int64
}

func (self *RecvMsgEvent) Reply(msg interface{}) {

	data, msgid, _ := cellnet.EncodeMessage(msg)

	evFunc := self.ses.Peer().EventFunc()
	if evFunc != nil {

		evFunc(socket.SendEvent{self.ses, &RemoteCallACK{
			MsgID:  msgid,
			Data:   data,
			CallID: self.callid,
		}})
	}
}