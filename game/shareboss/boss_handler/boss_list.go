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
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeShareBoss, worldboss.BossListHandlerFunc(shareBossList))
}

func shareBossList(pl player.Player, typ worldbosstypes.BossType) {
	//功能开启判断
	flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypeCrossWorldBoss)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shareBoss:跨服跨服世界BOSS列表错误，功能未开启")

		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	bossList := shareboss.GetShareBossService().GetShareBossList(typ)
	scShareBossList := pbutil.BuildSCWorldBossListShareBoss(bossList, int32(typ))
	pl.SendMsg(scShareBossList)
}
