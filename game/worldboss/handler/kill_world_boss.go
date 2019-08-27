package handler

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	gamesession "fgame/fgame/game/session"
	worldbosslogic "fgame/fgame/game/worldboss/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	// processor.Register(codec.MessageType(uipb.MessageType_CS_CHALLENGE_WORLD_BOSS_TYPE), dispatch.HandlerFunc(handlerKillWorldBoss))
}

//前往击杀世界BOSS请求
func handlerKillWorldBoss(s session.Session, msg interface{}) (err error) {
	log.Debug("worldBoss:处理前往击杀世界BOSS请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	cs := msg.(*uipb.CSChallengeWorldBoss)
	bossBiologyId := cs.GetBiologyId()

	err = killWorldBoss(tpl, bossBiologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("worldBoss:处理前往击杀世界BOSS请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("worldBoss:处理前往击杀世界BOSS请求完成")

	return
}

func killWorldBoss(pl player.Player, bossBiologyId int32) (err error) {
	return worldbosslogic.HandleKillWorldBoss(pl, bossBiologyId)
}
