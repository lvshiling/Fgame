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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GET_ATTACHMENT_BATCH_TYPE), dispatch.HandlerFunc(handlerGetEmailAttachementBatch))
}

//处理一键领取附件
func handlerGetEmailAttachementBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("email：处理一键领取附件请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	err = getEmailAttachementBatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("email:处理一键领取附件请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("email:处理一键领取附件请求完成")
	return
}

//处理一键领取附件逻辑
func getEmailAttachementBatch(pl player.Player) (err error) {
	return emaillogic.HandleGetEmailAttachementBatch(pl)
}
