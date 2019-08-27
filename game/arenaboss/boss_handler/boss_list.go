package boss_handler

import (
	"fgame/fgame/common/lang"

	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/shareboss/shareboss"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeArena, worldboss.BossListHandlerFunc(arenaBossList))
}

func arenaBossList(pl player.Player, t worldbosstypes.BossType) {
	//功能开启判断
	flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypeShengShouBoss)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("arenaboss:圣境BOSS列表错误，功能未开启")

		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	bossList := shareboss.GetShareBossService().GetShareBossList(t)
	scShareBossList := pbutil.BuildSCWorldBossListShareBoss(bossList, int32(t))
	pl.SendMsg(scShareBossList)
}
