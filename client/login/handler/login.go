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
	processor.Register(codec.MessageType(uipb.MessageType_SC_LOGIN_TYPE), dispatch.HandlerFunc(handleLogin))
}

//登陆成功
func handleLogin(s session.Session, msg interface{}) (err error) {
	log.Debug("login:处理登陆消息")
	scLogin := msg.(*uipb.SCLogin)
	uId := scLogin.GetUserId()
	cs := clientsession.SessionInContext(s.Context())

	pl := player.NewPlayer(uId, cs)
	cs.Auth(pl)
	pl.Auth()
	log.Debug("login:处理登陆消息完成")
	return
}
