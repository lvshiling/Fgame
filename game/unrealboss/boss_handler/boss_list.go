package boss_handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	"fgame/fgame/game/unrealboss/unrealboss"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeUnrealBoss, worldboss.BossListHandlerFunc(unrealBossList))
}

func unrealBossList(pl player.Player, typ worldbosstypes.BossType) {
	bossList := unrealboss.GetUnrealBossService().GetUnrealBossList()
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	unrealManager.RefreshPilao()

	curPilao := unrealManager.GetCurPilaoNum()
	curBuyTimes := unrealManager.GetPilaoBuyTimes()
	scMsg := pbutil.BuildSCWorldBossListUnrealBoss(bossList, curPilao, curBuyTimes, int32(typ))
	pl.SendMsg(scMsg)
}
