package handler

import (
	wxcontext "github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
)

type MsgTransferCustomerHandler struct {
}

func (m *MsgTransferCustomerHandler) HandlerMsg(v message.MixMessage, ctx *wxcontext.Context) *message.Reply {
	trans := message.NewTransferCustomer("")
	reply := &message.Reply{
		MsgType: message.MsgTypeTransfer,
		MsgData: trans,
	}
	return reply
}
