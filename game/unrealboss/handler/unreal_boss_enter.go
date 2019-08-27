package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"

	gamesession "fgame/fgame/game/session"
	unrealbosslogic "fgame/fgame/game/unrealboss/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerUnrealBossChallenge))
}

//幻境boss挑战
func handlerUnrealBossChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("unrealboss:处理幻境boss挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSUnrealBossChallenge)
	biologyId := csMsg.GetBiologyId()

	err = unrealbossChallenge(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"biologyId": biologyId,
				"err":       err,
			}).Error("unrealboss:处理幻境boss挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  tpl.GetId(),
			"biologyId": biologyId,
		}).Debug("unrealboss：处理幻境boss挑战请求完成")

	return
}

//幻境boss挑战逻辑
func unrealbossChallenge(pl player.Player, biologyId int32) (err error) {
	return unrealbosslogic.HandleUnrealbossChallenge(pl, biologyId)

}
