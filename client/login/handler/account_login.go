package handler

import (
	"fgame/fgame/client/client"
	"fgame/fgame/client/processor"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_ACCOUNT_LOGIN_TYPE), dispatch.HandlerFunc(handleAccountLogin))
}

//处理账户登陆
func handleAccountLogin(s session.Session, msg interface{}) (err error) {
	log.Debug("login:处理账户登陆")
	scAccountLogin := msg.(*uipb.SCAccountLogin)
	expiredTime := scAccountLogin.GetExpiredTime()
	token := scAccountLogin.GetToken()
	// serverInfoList := scAccountLogin.GetServerList()
	// playerInfoList := scAccountLogin.GetServerPlayerInfoList()

	c := client.ClientInContext(s.Context())
	err = accountLogin(c, token, expiredTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"token":       token,
				"expiredTime": expiredTime,
			}).Error("login:处理账户登陆,错误")
		return
	}
	log.Debug("login:处理账户登陆")
	return
}

func accountLogin(c *client.Client, token string, expiredTime int64) (err error) {
	c.AccountLogin(token, expiredTime)
	return nil
}
