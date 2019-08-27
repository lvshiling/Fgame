package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GET_ATTACHMENT_TYPE), dispatch.HandlerFunc(handlerGetEmailAttachement))
}

//领取附件请求
func handlerGetEmailAttachement(s session.Session, msg interface{}) (err error) {
	log.Debug("email：处理领取附件请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	emailId := msg.(*uipb.CSGetAttachment).GetEmailId()
	err = getEmailAttachement(tpl, emailId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("email:处理领取附件请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("email:处理领取附件请求完成")
	return
}

//领取附件逻辑
func getEmailAttachement(pl player.Player, emailId int64) (err error) {
	return emaillogic.HandleGetEmailAttachement(pl, emailId)
}
