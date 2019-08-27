package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHALLENGE_WORLD_BOSS_TYPE), dispatch.HandlerFunc(handlerCommonKillBoss))
}

//前往击杀通用BOSS请求
func handlerCommonKillBoss(s session.Session, msg interface{}) (err error) {
	log.Debug("worldBoss:处理前往击杀通用BOSS请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSChallengeWorldBoss)

	bossBiologyId := csMsg.GetBiologyId()
	typ := csMsg.GetBossType()
	bossType := worldbosstypes.BossType(typ)
	if !bossType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:处理前往击杀通用BOSS请求，boss类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = commonkillBoss(tpl, bossType, bossBiologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("worldBoss:处理前往击杀通用BOSS请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("worldBoss:处理前往击杀通用BOSS请求完成")

	return
}

func commonkillBoss(pl player.Player, bossType worldbosstypes.BossType, bossBiologyId int32) (err error) {
	h := worldboss.GetKillBossHandler(bossType)
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:处理前往击杀通用BOSS请求，处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.WorldBossHandlerNotExist)
		return
	}

	h.KillBoss(pl, bossType, bossBiologyId)
	return
}
