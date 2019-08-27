package handler

import (
	log "github.com/Sirupsen/logrus"
	wxcontext "github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/message"
)

type EventSubscribeHandler struct {
}

func (m *EventSubscribeHandler) HandlerMsg(v message.MixMessage, ctx *wxcontext.Context) *message.Reply {
	//tousername := v.ToUserName
	log.Debug("微信请求关注")
	return nil
}
