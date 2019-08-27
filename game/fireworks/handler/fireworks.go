package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	fireworkslogic "fgame/fgame/game/fireworks/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FIREWORKS_TYPE), dispatch.HandlerFunc(handleShootFireworks))
}

//处理烟花信息
func handleShootFireworks(s session.Session, msg interface{}) (err error) {
	log.Debug("fireworks:处理烟花消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFireWorks := msg.(*uipb.CSFireWorks)
	itemId := csFireWorks.GetItemId()
	num := csFireWorks.GetNum()

	err = shootFireworks(tpl, itemId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
				"error":    err,
			}).Error("fireworks:处理烟花消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Debug("fireworks:处理烟花消息完成")
	return nil
}

//处理烟花信息逻辑
func shootFireworks(pl player.Player, itemId int32, num int32) (err error) {
	fireworkslogic.ShootFireworks(pl, itemId, num, true)
	return
}
