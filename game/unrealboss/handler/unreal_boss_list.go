package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/unrealboss/pbutil"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	// playertypes "fgame/fgame/game/player/types"
	// playerunrealboss "fgame/fgame/game/unrealboss/player"
	"fgame/fgame/game/unrealboss/unrealboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_UNREAL_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerUnrealBossList))
}

//幻境BOSS列表请求
func handlerUnrealBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("unrealBoss:处理幻境BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = unrealBossList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("unrealBoss:处理幻境BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("unrealBoss:处理幻境BOSS列表请求完成")

	return
}

func unrealBossList(pl player.Player) (err error) {
	bossList := unrealboss.GetUnrealBossService().GetUnrealBossList()
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	unrealManager.RefreshPilao()

	curPilao := unrealManager.GetCurPilaoNum()
	curBuyTimes := unrealManager.GetPilaoBuyTimes()
	scMsg := pbutil.BuildSCUnrealBossList(bossList, curPilao, curBuyTimes)
	pl.SendMsg(scMsg)
	return
}
