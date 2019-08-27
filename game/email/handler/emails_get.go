package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/email/pbutil"
	playeremail "fgame/fgame/game/email/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMAILS_GET_TYPE), dispatch.HandlerFunc(handlerEmailsGet))
}

//获取邮件列表
func handlerEmailsGet(s session.Session, msg interface{}) error {
	log.Debug("email：处理获取邮件列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err := emailsGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("email:处理邮件列表请求，错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("email:处理邮件列表请求完成")
	return nil
}

//获取邮件列表逻辑
func emailsGet(pl player.Player) (err error) {
	emailDataManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	emailList := emailDataManager.GetEmails()
	scEmailsGet := pbutil.BuildSCEmailsGet(emailList)
	pl.SendMsg(scEmailsGet)

	return nil
}
