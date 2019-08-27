package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_ACTIVE_TYPE), dispatch.HandlerFunc(handleLingTongActivate))

}

//处理灵童激活信息
func handleLingTongActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童激活消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongActive := msg.(*uipb.CSLingTongActive)
	lingTongId := csLingTongActive.GetLingTongId()

	err = lingTongActivate(tpl, lingTongId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Error("lingtong:处理获取灵童激活消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Debug("lingtong:处理获取灵童激活消息完成")
	return nil

}

//获取灵童激活界面信息逻辑
func lingTongActivate(pl player.Player, lingTongId int32) (err error) {
	return lingtonglogic.HandleLingTongActivate(pl, lingTongId)

}
