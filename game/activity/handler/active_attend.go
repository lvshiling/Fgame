package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ACTIVITY_ATTEND_TYPE), dispatch.HandlerFunc(handlerActiveAttend))
}

//进入活动
func handlerActiveAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("active:处理进入活动请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSActivityAttend)
	activeId := csMsg.GetActiveId()
	args := csMsg.GetArgs()

	activeType := activitytypes.ActivityType(activeId)
	if !activeType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"activeId": activeId,
			}).Warn("activity:进入活动请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = activeAttend(tpl, activeType, args)

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"activeId": activeId,
				"err":      err,
			}).Error("active:处理进入活动请求，错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"activeId": activeId,
		}).Info("active:处理进入活动请求，完成")
	return
}

//活动进入请求逻辑
func activeAttend(pl player.Player, activeType activitytypes.ActivityType, args string) (err error) {
	return activitylogic.HandleActiveAttend(pl, activeType, args)
}
