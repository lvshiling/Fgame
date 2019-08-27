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
	processor.Register(codec.MessageType(uipb.MessageType_CS_WORLD_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerCommonBossList))
}

//通用BOSS列表请求
func handlerCommonBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("worldBoss:处理通用BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWorldBossList)
	typ := csMsg.GetBossType()

	bossType := worldbosstypes.BossType(typ)
	if !bossType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossType": typ,
			}).Warn("transport:通用BOSS列表错误，boss类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = commonBossList(tpl, bossType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("worldBoss:处理通用BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("worldBoss:处理通用BOSS列表请求完成")

	return
}

func commonBossList(pl player.Player, bossType worldbosstypes.BossType) (err error) {
	h := worldboss.GetBossListHandler(bossType)
	if h == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:通用BOSS列表错误，处理器不存在")
		playerlogic.SendSystemMessage(pl, lang.WorldBossHandlerNotExist)
		return
	}

	h.BossList(pl, bossType)
	return
}
