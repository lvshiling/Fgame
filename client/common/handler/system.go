package handler

import (
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_SYSTEM_MESSAGE_TYPE), dispatch.HandlerFunc(handleSystemMessage))
}

//系统错误
func handleSystemMessage(s session.Session, msg interface{}) (err error) {
	log.Debug("common:处理系统提示消息")
	scSystemMessage := msg.(*uipb.SCSystemMessage)
	content := scSystemMessage.GetContent()
	args := scSystemMessage.GetArgs()
	cs := clientsession.SessionInContext(s.Context())

	sessionSystemMessage(cs, content, args)

	log.WithFields(
		log.Fields{

			"content": content,
			"args":    args,
		}).Debug("common:处理系统提示消息完成")
	return
}

func sessionSystemMessage(s clientsession.Session, content string, args []string) (err error) {
	log.WithFields(
		log.Fields{

			"content": content,
			"args":    args,
		}).Info("player:玩家收到系统消息")
	return
}
