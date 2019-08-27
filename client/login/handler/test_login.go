package handler

import (
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_TEST_LOGIN_TYPE), dispatch.HandlerFunc(handleTestLogin))
}

//测试登陆成功
func handleTestLogin(s session.Session, msg interface{}) (err error) {
	log.Debug("login:处理测试登陆消息")
	scTestLogin := msg.(*uipb.SCTestLogin)
	uId := scTestLogin.GetUserId()
	cs := clientsession.SessionInContext(s.Context())

	pl := player.NewPlayer(uId, cs)
	cs.Auth(pl)
	pl.Auth()
	log.Debug("login:处理登陆消息完成")
	return
}
